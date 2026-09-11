// Скачать свой трек Suno по ссылке embed/song или по UUID клипа.
//
// Берёт HTML эмбеда, достаёт зашифрованный progressive MP4 (m4a-opus),
// просит ключ у POST /api/mango/rights и расшифровывает AES-CTR.
//
// Креды (необязательно) — JSON, см. suno_creds.example.json рядом:
//
//	{
//	  "cookies": "__session=...; other=1",
//	  "headers": {
//	    "Authorization": "Bearer eyJ..."
//	  }
//	}
//
// cookies — строка заголовка Cookie или объект {имя: значение}.
// Authorization — JWT Clerk; если есть, ключ разворачивается из SHA-256(JWT),
// как getAuthToken + getUserKey в плеере. Без файла кредов работает
// анонимно через guest-токен glt.
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// Совпадает с getApiBaseUrl() в чанке плеера (1jw5qfhlrr3ch.js).
const apiBase = "https://studio-api-prod.suno.com"
const rightsPath = "/api/mango/rights"

// content_type в POST mango/rights — литерал "clip" из loadEncryptedAudio.
const rightsContentType = "clip"
const embedOrigin = "https://suno.com"
const chromeUA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"

var uuidRE = regexp.MustCompile(`(?i)[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)

// {title} и <title> — одно и то же. Угловые скобки в шелле лучше не использовать.
var nameTokenRE = regexp.MustCompile(`\{([A-Za-z_]+)\}|<([A-Za-z_]+)>`)

// Типы, которые плеер считает зашифрованными (ENCRYPTED_CONTENT_TYPES).
var encryptedContentTypes = map[string]bool{
	"m4a-opus":  true,
	"webm-opus": true,
}

var httpClient = &http.Client{Timeout: 60 * time.Second}

type creds struct {
	Cookies json.RawMessage   `json:"cookies"`
	Cookie  json.RawMessage   `json:"cookie"`
	Headers map[string]string `json:"headers"`
	Token   string            `json:"token"`
	JWT     string            `json:"jwt"`
}

type clipJSON struct {
	ID                string       `json:"id"`
	Title             string       `json:"title"`
	DisplayName       string       `json:"display_name"`
	Handle            string       `json:"handle"`
	DisplayTags       string       `json:"display_tags"`
	CreatedAt         string       `json:"created_at"`
	ImageURL          string       `json:"image_url"`
	MajorModelVersion string       `json:"major_model_version"`
	ModelName         string       `json:"model_name"`
	Metadata          clipMetadata `json:"metadata"`
	MediaURLs         []mediaItem  `json:"media_urls"`
}

type clipMetadata struct {
	Tags         string `json:"tags"`
	Prompt       string `json:"prompt"`
	NegativeTags string `json:"negative_tags"`
}

type mediaItem struct {
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
	Delivery    string `json:"delivery"`
	Encoding    string `json:"encoding"`
}

type rightsJSON struct {
	Key string `json:"key"`
	IV  string `json:"iv"`
	GLT string `json:"glt"`
}

// Имя файла без --out. Каталог — cwd, откуда вызвали CLI.
const defaultNameTemplate = "{title} - {id}"

type cliArgs struct {
	url      string
	out      string
	creds    string
	nameTmpl string
}

func usage() {
	fmt.Fprintf(os.Stderr, "Скачать свой трек Suno (расшифрованный MP4/Opus .m4a).\n\n")
	fmt.Fprintf(os.Stderr, "использование: %s [--out путь] [--name шаблон] [--creds suno_creds.json] <url>\n", filepath.Base(os.Args[0]))
	fmt.Fprintf(os.Stderr, "  url: https://suno.com/embed/<uuid>, /song/<uuid> или сам uuid\n")
	fmt.Fprintf(os.Stderr, "  --out    файл .m4a или каталог. Без флага — каталог, откуда вызвали команду\n")
	fmt.Fprintf(os.Stderr, "  --name   шаблон имени, если --out не указывает конкретный файл\n")
	fmt.Fprintf(os.Stderr, "           плейсхолдеры: {title} {artist} {id} {handle} {genre} {date} {model}\n")
	fmt.Fprintf(os.Stderr, "           по умолчанию: %s\n", defaultNameTemplate)
	fmt.Fprintf(os.Stderr, "  --creds  JSON с cookies/headers. Если не указан — ищем suno_creds.json рядом\n")
}

// Флаги можно ставить до и после URL, как в argparse.
func parseArgs(args []string) (cliArgs, error) {
	var a cliArgs
	var rest []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		take := func() (string, error) {
			if i+1 >= len(args) {
				return "", fmt.Errorf("флаг %s требует значение", arg)
			}
			i++
			return args[i], nil
		}
		switch {
		case arg == "-h" || arg == "--help":
			usage()
			os.Exit(0)
		case arg == "--out" || arg == "-out":
			v, err := take()
			if err != nil {
				return cliArgs{}, err
			}
			a.out = v
		case strings.HasPrefix(arg, "--out="):
			a.out = strings.TrimPrefix(arg, "--out=")
		case arg == "--name" || arg == "-name" || arg == "--template":
			v, err := take()
			if err != nil {
				return cliArgs{}, err
			}
			a.nameTmpl = v
		case strings.HasPrefix(arg, "--name="):
			a.nameTmpl = strings.TrimPrefix(arg, "--name=")
		case strings.HasPrefix(arg, "--template="):
			a.nameTmpl = strings.TrimPrefix(arg, "--template=")
		case arg == "--creds" || arg == "-creds":
			v, err := take()
			if err != nil {
				return cliArgs{}, err
			}
			a.creds = v
		case strings.HasPrefix(arg, "--creds="):
			a.creds = strings.TrimPrefix(arg, "--creds=")
		case strings.HasPrefix(arg, "-"):
			return cliArgs{}, fmt.Errorf("неизвестный флаг %s", arg)
		default:
			rest = append(rest, arg)
		}
	}
	if len(rest) != 1 {
		return cliArgs{}, errors.New("нужен ровно один URL или uuid")
	}
	a.url = rest[0]
	if a.nameTmpl == "" {
		a.nameTmpl = defaultNameTemplate
	}
	return a, nil
}

func main() {
	a, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %v\n\n", err)
		usage()
		os.Exit(2)
	}

	credsFile := a.creds
	if credsFile == "" {
		credsFile = defaultCredsPath()
	}

	if err := download(a.url, a.out, a.nameTmpl, credsFile); err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %v\n", err)
		os.Exit(1)
	}
}

func parseClipID(text string) (string, error) {
	m := uuidRE.FindString(strings.TrimSpace(text))
	if m == "" {
		return "", fmt.Errorf("в аргументе нет UUID клипа: %q", text)
	}
	return strings.ToLower(m), nil
}

func loadCreds(path string) (creds, error) {
	var c creds
	if path == "" {
		return c, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return c, nil
		}
		return c, err
	}
	if err := json.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("креды %s — не JSON: %w", path, err)
	}
	return c, nil
}

func cookieHeader(c creds) string {
	raw := c.Cookies
	if len(raw) == 0 {
		raw = c.Cookie
	}
	if len(raw) == 0 {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		return strings.TrimSpace(s)
	}
	var obj map[string]any
	if err := json.Unmarshal(raw, &obj); err != nil {
		return ""
	}
	parts := make([]string, 0, len(obj))
	for k, v := range obj {
		if v == nil {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(parts, "; ")
}

func extraHeaders(c creds) map[string]string {
	if c.Headers == nil {
		return map[string]string{}
	}
	out := make(map[string]string, len(c.Headers))
	for k, v := range c.Headers {
		if v == "" {
			continue
		}
		out[k] = v
	}
	return out
}

var bearerRE = regexp.MustCompile(`(?i)^bearer\s+`)

func stripBearer(token string) string {
	return strings.TrimSpace(bearerRE.ReplaceAllString(token, ""))
}

func authJWT(c creds) string {
	for k, v := range extraHeaders(c) {
		if strings.EqualFold(k, "authorization") {
			return stripBearer(v)
		}
	}
	if s := strings.TrimSpace(c.Token); s != "" {
		return stripBearer(s)
	}
	if s := strings.TrimSpace(c.JWT); s != "" {
		return stripBearer(s)
	}
	return ""
}

func httpRequest(method, url string, body []byte, headers map[string]string) (int, []byte, error) {
	var rdr io.Reader
	if body != nil {
		rdr = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, url, rdr)
	if err != nil {
		return 0, nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("сеть %s: %w", url, err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	if resp.StatusCode >= 400 {
		snip := data
		if len(snip) > 300 {
			snip = snip[:300]
		}
		return resp.StatusCode, data, fmt.Errorf("HTTP %d %s: %q", resp.StatusCode, url, snip)
	}
	return resp.StatusCode, data, nil
}

func baseHeaders(clipID string, c creds) map[string]string {
	h := map[string]string{
		"User-Agent": chromeUA,
		"Accept":     "*/*",
		"Origin":     embedOrigin,
		"Referer":    embedOrigin + "/embed/" + clipID,
	}
	if cookie := cookieHeader(c); cookie != "" {
		h["Cookie"] = cookie
	}
	for k, v := range extraHeaders(c) {
		h[k] = v
	}
	return h
}

func unescapeRSC(text string) string {
	// Снимает экранирование кавычек RSC, не трогая JSON-эскейпы вроде \n.
	return strings.ReplaceAll(strings.ReplaceAll(text, `\/`, "/"), `\"`, `"`)
}

func extractJSONObject(text string, startBrace int) ([]byte, error) {
	depth := 0
	inStr := false
	escape := false
	for i := startBrace; i < len(text); i++ {
		ch := text[i]
		if inStr {
			if escape {
				escape = false
			} else if ch == '\\' {
				escape = true
			} else if ch == '"' {
				inStr = false
			}
			continue
		}
		switch ch {
		case '"':
			inStr = true
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return []byte(text[startBrace : i+1]), nil
			}
		}
	}
	return nil, errors.New("не закрылся JSON-объект клипа в HTML")
}

func parseClipFromHTML(html, clipID string) (clipJSON, error) {
	var last error
	for _, text := range []string{html, unescapeRSC(html)} {
		marker := `"id":"` + clipID + `"`
		idx := 0
		for {
			pos := strings.Index(text[idx:], marker)
			if pos < 0 {
				break
			}
			pos += idx
			clipKey := strings.LastIndex(text[:pos], `"clip":`)
			brace := -1
			if clipKey >= 0 {
				brace = strings.Index(text[clipKey:], "{")
				if brace >= 0 {
					brace += clipKey
				}
			}
			if brace >= 0 && pos-brace < 4000 {
				raw, err := extractJSONObject(text, brace)
				if err != nil {
					last = err
				} else {
					var clip clipJSON
					if err := json.Unmarshal(raw, &clip); err != nil {
						last = err
					} else if strings.ToLower(clip.ID) == clipID {
						return clip, nil
					}
				}
			}
			idx = pos + len(marker)
		}
	}
	if last != nil {
		return clipJSON{}, fmt.Errorf("в HTML эмбеда нет JSON клипа %s (%v)", clipID, last)
	}
	return clipJSON{}, fmt.Errorf("в HTML эмбеда нет JSON клипа %s", clipID)
}

func pickEncryptedMedia(clip clipJSON) (mediaItem, error) {
	// Как findAudioMediaItem(..., {encrypted:true, delivery:'progressive'}).
	if clip.MediaURLs == nil {
		return mediaItem{}, errors.New("у клипа нет media_urls")
	}
	var progressive, anyEnc []mediaItem
	for _, item := range clip.MediaURLs {
		if item.Encoding == "" || item.URL == "" || !encryptedContentTypes[item.ContentType] {
			continue
		}
		anyEnc = append(anyEnc, item)
		if item.Delivery == "progressive" {
			progressive = append(progressive, item)
		}
	}
	choose := func(items []mediaItem) mediaItem {
		// Сначала m4a-opus — это progressive файл, который качаем.
		for _, it := range items {
			if it.ContentType == "m4a-opus" {
				return it
			}
		}
		return items[0]
	}
	if len(progressive) > 0 {
		return choose(progressive), nil
	}
	if len(anyEnc) > 0 {
		return choose(anyEnc), nil
	}
	return mediaItem{}, errors.New("нет зашифрованного media_urls (m4a-opus/webm-opus с encoding). Незашифрованный audio_url скрипт не подставляет")
}

func fetchEmbedClip(clipID string, c creds) (clipJSON, error) {
	url := embedOrigin + "/embed/" + clipID
	fmt.Printf("эмбед: %s\n", url)
	h := baseHeaders(clipID, c)
	h["Accept"] = "text/html,application/xhtml+xml"
	status, body, err := httpRequest(http.MethodGet, url, nil, h)
	if err != nil {
		return clipJSON{}, err
	}
	if status != 200 {
		return clipJSON{}, fmt.Errorf("эмбед HTTP %d", status)
	}
	return parseClipFromHTML(string(body), clipID)
}

func fetchRights(clipID string, c creds) (rightsJSON, error) {
	url := apiBase + rightsPath
	payload, err := json.Marshal(map[string]any{
		"content_params": map[string]string{
			"content_id":   clipID,
			"content_type": rightsContentType,
		},
	})
	if err != nil {
		return rightsJSON{}, err
	}
	h := baseHeaders(clipID, c)
	h["Content-Type"] = "application/json"
	h["Accept"] = "application/json"
	fmt.Printf("лицензия: POST %s\n", url)
	status, body, err := httpRequest(http.MethodPost, url, payload, h)
	if err != nil {
		return rightsJSON{}, err
	}
	if status != 200 {
		snip := body
		if len(snip) > 300 {
			snip = snip[:300]
		}
		return rightsJSON{}, fmt.Errorf("mango/rights HTTP %d: %q", status, snip)
	}
	var r rightsJSON
	if err := json.Unmarshal(body, &r); err != nil {
		return rightsJSON{}, errors.New("mango/rights вернул не JSON")
	}
	if r.Key == "" || r.IV == "" || r.GLT == "" {
		return rightsJSON{}, fmt.Errorf("mango/rights без key/iv/glt")
	}
	return r, nil
}

func userKeyBytes(secret string) []byte {
	// getUserKey: SHA-256(utf-8) → AES-GCM-256.
	sum := sha256.Sum256([]byte(secret))
	return sum[:]
}

func unwrapField(b64Value, clipID string, gcmKey []byte) ([]byte, error) {
	// decodeContentKey / decodeContentIv: AES-GCM, IV=12, AAD=clip id.
	wrapped, err := base64.StdEncoding.DecodeString(b64Value)
	if err != nil {
		return nil, fmt.Errorf("обёртка не base64: %w", err)
	}
	if len(wrapped) < 13+16 {
		return nil, fmt.Errorf("обёртка слишком короткая: %d байт", len(wrapped))
	}
	block, err := aes.NewCipher(gcmKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	pt, err := gcm.Open(nil, wrapped[:12], wrapped[12:], []byte(clipID))
	if err != nil {
		return nil, errors.New("не развернулся key/iv (AES-GCM). Если трек из своего аккаунта — нужен JWT в creds.headers.Authorization")
	}
	return pt, nil
}

func resolveUnwrapKey(r rightsJSON, c creds) ([]byte, error) {
	// resolveLicenseUserKey: JWT, иначе guest glt.
	if jwt := authJWT(c); jwt != "" {
		return userKeyBytes(jwt), nil
	}
	if r.GLT == "" {
		return nil, errors.New("в ответе rights нет glt, а JWT в кредах не задан")
	}
	return userKeyBytes(r.GLT), nil
}

func ctrIV(rawIV []byte) []byte {
	// Как в JS: new Uint8Array(16); r.set(iv).
	buf := make([]byte, 16)
	copy(buf, rawIV)
	return buf
}

func decryptMedia(ciphertext, contentKey, contentIV []byte) ([]byte, error) {
	// AES-CTR, counter length 128 бит, весь файл целиком.
	if n := len(contentKey); n != 16 && n != 24 && n != 32 {
		return nil, fmt.Errorf("неожиданная длина контент-ключа: %d", n)
	}
	block, err := aes.NewCipher(contentKey)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, ctrIV(contentIV))
	out := make([]byte, len(ciphertext))
	stream.XORKeyStream(out, ciphertext)
	return out, nil
}

func assertMP4(data []byte) error {
	// sniffAudioFormat: байты 4..8 == 'ftyp'.
	if len(data) < 12 || string(data[4:8]) != "ftyp" {
		head := ""
		if len(data) > 0 {
			n := 16
			if len(data) < n {
				n = len(data)
			}
			head = hex.EncodeToString(data[:n])
		}
		return fmt.Errorf("после расшифровки нет MP4 ftyp (первые байты %s). Ключ, IV или файл не сошлись", head)
	}
	return nil
}

func sanitizeFilenamePart(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' || unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	spaceRE := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(spaceRE.ReplaceAllString(b.String(), " "))
}

func clipDateStamp(createdAt string) string {
	s := strings.TrimSpace(createdAt)
	if s == "" {
		return ""
	}
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t.Format("2006-01-02")
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t.Format("2006-01-02")
	}
	if len(s) >= 10 {
		return sanitizeFilenamePart(s[:10])
	}
	return sanitizeFilenamePart(s)
}

func filenameFromTemplate(tmpl string, clip clipJSON) (string, error) {
	if strings.TrimSpace(tmpl) == "" {
		tmpl = defaultNameTemplate
	}
	title := strings.TrimSpace(clip.Title)
	if title == "" {
		title = "suno"
	}
	vals := map[string]string{
		"title":  title,
		"artist": strings.TrimSpace(clip.DisplayName),
		"id":     strings.TrimSpace(clip.ID),
		"uuid":   strings.TrimSpace(clip.ID),
		"handle": strings.TrimSpace(clip.Handle),
		"genre":  strings.TrimSpace(clip.DisplayTags),
		"date":   clipDateStamp(clip.CreatedAt),
		"model":  clip.modelLabel(),
	}
	unknown := ""
	name := nameTokenRE.ReplaceAllStringFunc(tmpl, func(m string) string {
		inner := strings.Trim(m, "{}<>")
		v, ok := vals[strings.ToLower(inner)]
		if !ok {
			unknown = inner
			return m
		}
		return sanitizeFilenamePart(v)
	})
	if unknown != "" {
		return "", fmt.Errorf("неизвестный плейсхолдер {%s}; доступны title artist id handle genre date model", unknown)
	}
	name = strings.TrimSpace(name)
	name = strings.Trim(name, "-_ ")
	if name == "" {
		name = "suno"
	}
	if strings.Contains(name, string(os.PathSeparator)) || strings.Contains(name, "/") || strings.Contains(name, `\`) {
		return "", fmt.Errorf("шаблон имени не должен содержать разделители пути")
	}
	if !strings.HasSuffix(strings.ToLower(name), ".m4a") {
		name += ".m4a"
	}
	return name, nil
}

// Каталог по умолчанию — cwd, откуда вызвали CLI. --out с конкретным файлом
// перекрывает --name; существующий каталог (или путь со слэшем на конце) —
// туда кладётся имя из шаблона.
func resolveOutPath(outFlag, nameTmpl string, clip clipJSON) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	name, err := filenameFromTemplate(nameTmpl, clip)
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(outFlag) == "" {
		return filepath.Abs(filepath.Join(cwd, name))
	}
	p := outFlag
	if !filepath.IsAbs(p) {
		p = filepath.Join(cwd, p)
	}
	sep := strings.HasSuffix(outFlag, "/") || strings.HasSuffix(outFlag, string(os.PathSeparator))
	if st, err := os.Stat(p); err == nil && st.IsDir() {
		return filepath.Abs(filepath.Join(p, name))
	}
	if sep {
		return filepath.Abs(filepath.Join(p, name))
	}
	return filepath.Abs(p)
}

func defaultCredsPath() string {
	cwd, _ := os.Getwd()
	exe, _ := os.Executable()
	exeDir := filepath.Dir(exe)
	candidates := []string{
		filepath.Join(cwd, "suno_creds.json"),
		filepath.Join(exeDir, "suno_creds.json"),
		filepath.Join(cwd, "scripts", "suno_creds.json"),
		filepath.Join(cwd, "scripts", "suno-download", "suno_creds.json"),
	}
	for _, p := range candidates {
		if st, err := os.Stat(p); err == nil && !st.IsDir() {
			return p
		}
	}
	return ""
}

func download(urlOrID, outFlag, nameTmpl, credsPath string) error {
	clipID, err := parseClipID(urlOrID)
	if err != nil {
		return err
	}
	fmt.Printf("id: %s\n", clipID)

	c, err := loadCreds(credsPath)
	if err != nil {
		return err
	}
	switch {
	case credsPath != "" && (len(c.Cookies) > 0 || len(c.Cookie) > 0 || len(c.Headers) > 0 || c.Token != "" || c.JWT != ""):
		fmt.Printf("креды: %s\n", credsPath)
	case credsPath != "":
		if _, err := os.Stat(credsPath); os.IsNotExist(err) {
			fmt.Printf("креды: файла нет (%s), иду без сессии\n", credsPath)
		} else {
			fmt.Printf("креды: %s\n", credsPath)
		}
	default:
		fmt.Println("креды: не заданы, иду без сессии")
	}

	clip, err := fetchEmbedClip(clipID, c)
	if err != nil {
		return err
	}
	if clip.ID == "" {
		clip.ID = clipID
	}
	title := clip.Title
	if title == "" {
		title = "suno"
	}
	media, err := pickEncryptedMedia(clip)
	if err != nil {
		return err
	}
	fmt.Printf("медиа: %s %s %s\n", media.ContentType, media.Delivery, media.URL)

	rights, err := fetchRights(clipID, c)
	if err != nil {
		return err
	}
	gcmKey, err := resolveUnwrapKey(rights, c)
	if err != nil {
		return err
	}
	contentKey, err := unwrapField(rights.Key, clipID, gcmKey)
	if err != nil {
		return err
	}
	contentIV, err := unwrapField(rights.IV, clipID, gcmKey)
	if err != nil {
		return err
	}

	fmt.Println("скачиваю шифротекст…")
	status, ciphertext, err := httpRequest(http.MethodGet, media.URL, nil, baseHeaders(clipID, c))
	if err != nil {
		return err
	}
	if status != 200 && status != 206 {
		return fmt.Errorf("CDN HTTP %d", status)
	}
	fmt.Printf("шифротекст: %d байт\n", len(ciphertext))

	fmt.Println("расшифровываю AES-CTR…")
	plaintext, err := decryptMedia(ciphertext, contentKey, contentIV)
	if err != nil {
		return err
	}
	if err := assertMP4(plaintext); err != nil {
		return err
	}

	existingComment, err := readIlstComment(bytes.NewReader(plaintext))
	if err != nil {
		fmt.Printf("comment в файле не прочитался: %v\n", err)
		existingComment = ""
	}

	var cover []byte
	if strings.TrimSpace(clip.ImageURL) != "" {
		fmt.Println("обложка: скачиваю…")
		cover, err = fetchCover(clip.ImageURL, clipID, c)
		if err != nil {
			return err
		}
		fmt.Printf("обложка: %d байт\n", len(cover))
	}

	tags := itunesTags{
		Title:   title,
		Artist:  strings.TrimSpace(clip.DisplayName),
		Genre:   strings.TrimSpace(clip.DisplayTags),
		Date:    strings.TrimSpace(clip.CreatedAt),
		Comment: mergeComment(existingComment, clip.modelLabel(), clip.CreatedAt, clip.ID),
		Cover:   cover,
	}

	outPath, err := resolveOutPath(outFlag, nameTmpl, clip)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}
	fmt.Println("теги: пишу ilst…")
	if err := writeTaggedM4A(plaintext, outPath, tags); err != nil {
		return err
	}
	sidePath, err := writeSidecar(outPath, clip)
	if err != nil {
		return err
	}
	st, err := os.Stat(outPath)
	if err != nil {
		return err
	}
	fmt.Printf("сохранил: %s (%d байт)\n", outPath, st.Size())
	fmt.Printf("sidecar: %s\n", sidePath)
	return nil
}

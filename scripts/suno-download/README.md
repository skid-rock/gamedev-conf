# Скачать свой трек Suno

Утилита по ссылке на **свой** клип (embed, страница песни или UUID) сохраняет расшифрованный MP4/Opus `.m4a`. Чужие каталоги не обходит.

## Сборка

Нужен Go. Из этого каталога:

```bash
make build
```

Бинарник появится рядом: `./suno-download`. Удалить: `make clean`.

В систему (в `GOBIN` или `$(go env GOPATH)/bin`, обычно `~/go/bin` — каталог должен быть в `PATH`):

```bash
make install
```

## Запуск

Из любого каталога — файлы пишутся **туда, откуда вызвали команду**, если не задан `--out`:

```bash
cd ~/Music
/path/to/suno-download 'https://suno.com/embed/dfdc0fc2-679b-4d68-b3a7-e1f76e18bcf5'
# → ~/Music/Dancefloor rollers - dfdc0fc2-….m4a и .json рядом
```

То же по странице песни, share-ссылке (`?sh=` или короткая `/s/…`) или по голому id:

```bash
./suno-download 'https://suno.com/song/dfdc0fc2-679b-4d68-b3a7-e1f76e18bcf5'
./suno-download 'https://suno.com/song/7fb8ea93-43c5-4149-a2a9-68573b357a6f?sh=SdRlW37AUjSqm19I'
./suno-download 'https://suno.com/s/SdRlW37AUjSqm19I'
./suno-download dfdc0fc2-679b-4d68-b3a7-e1f76e18bcf5
```

`/s/<код>` на стороне Suno отвечает 307 на `/song/<uuid>?sh=<код>`. Параметр `sh` после этого уходит в эмбед и в Referer. Кавычки вокруг URL обязательны — иначе шелл отрежет query.

Шаблон имени (кавычки обязательны, иначе шелл съест скобки):

```bash
./suno-download 'https://suno.com/embed/dfdc0fc2-679b-4d68-b3a7-e1f76e18bcf5' --name '{artist} - {title}'
# → SkidRock - Dancefloor rollers.m4a
```

Флаги можно ставить до или после URL:

| Флаг | Смысл |
|---|---|
| `--out путь` | конкретный `.m4a` **или** каталог. Без флага — текущий каталог запуска. Если файл уже есть — `имя 2.m4a`, как Finder «Keep Both» |
| `--name шаблон` | имя файла, если `--out` не указывает на конкретный файл. По умолчанию `{title} - {id}` |
| `--creds suno_creds.json` | cookies и заголовки. Если не указан, ищет `suno_creds.json` рядом с бинарником / в cwd |

Плейсхолдеры в `--name`: `{title}` `{artist}` `{id}` `{handle}` `{genre}` `{date}` `{model}`. То же в угловых скобках (`<title>`) тоже поймёт, но в шелле это редирект — лучше фигурные. Расширение `.m4a` допишется само. Sidecar `.json` берёт то же базовое имя.

Конкретный файл перекрывает шаблон:

```bash
./suno-download 'https://suno.com/embed/…' --out track.m4a
./suno-download 'https://suno.com/embed/…' --out ~/Music/suno/ --name '{artist} - {title}'
```

Креды не обязательны: публичный эмбед скачивается без файла. Формат — `suno_creds.example.json`:

```json
{
  "cookies": "__session=...; other=1",
  "headers": {
    "Authorization": "Bearer eyJ..."
  }
}
```

`cookies` — строка заголовка Cookie или объект `{имя: значение}`. JWT в `Authorization` нужен, если ключ обёрнут аккаунтом, а не гостевым токеном. Реальный `suno_creds.json` в git не класть.

## Теги и sidecar

После расшифровки в `.m4a` дописываются iTunes-атомы `ilst` (title, artist, genre, date, обложка `covr`) **без перекодирования** Opus. Comment, который кладёт сам Suno (`made with suno; created=…; id=…`), сохраняется; к нему может добавиться модель. Рядом с тем же базовым именем пишется `.json`: prompt, style, negative tags, автор, handle — без URL CDN и ключей.

ffmpeg для скачивания не нужен.

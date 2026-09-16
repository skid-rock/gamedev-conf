package main

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

// Лад: семь ступеней в полутонах от тоники.
type scaleType struct {
	Name  string
	Steps [7]int
}

var scaleTypes = []scaleType{
	{"major", [7]int{0, 2, 4, 5, 7, 9, 11}},
	{"minor", [7]int{0, 2, 3, 5, 7, 8, 10}},
	{"mixolydian", [7]int{0, 2, 4, 5, 7, 9, 10}},
	{"dorian", [7]int{0, 2, 3, 5, 7, 9, 10}},
	{"lydian", [7]int{0, 2, 4, 6, 7, 9, 11}},
	{"phrygian", [7]int{0, 1, 3, 5, 7, 8, 10}},
	{"locrian", [7]int{0, 1, 3, 5, 6, 8, 10}},
	{"harmonic minor", [7]int{0, 2, 3, 5, 7, 8, 11}},
}

// Вес ступени в шаблоне лада, по порядку I–VII.
// Ориентир — профиль Крумхансла — Кесслера для мажора, но II, VI и VII подняты:
// у него они почти вровень с нотами вне лада, и соседние лады не различаются.
var degreeWeights = [7]float64{6.0, 3.5, 4.5, 4.0, 5.0, 3.5, 3.5}

// Вес ноты вне звукоряда.
const outsideWeight = 2.0

// Тональность-кандидат и её счёт — корреляция хромы с шаблоном лада.
type keyCandidate struct {
	Tonic int
	Scale int
	Score float64
}

func (k keyCandidate) Name() string {
	return noteNames[k.Tonic] + " " + scaleTypes[k.Scale].Name
}

func (k keyCandidate) template() chroma {
	var t chroma
	for i := range t {
		t[i] = outsideWeight
	}
	for d, step := range scaleTypes[k.Scale].Steps {
		t[(k.Tonic+step)%12] = degreeWeights[d]
	}
	return t
}

// Ноты звукоряда как множество классов высоты.
func (k keyCandidate) pitchSet() [12]bool {
	var set [12]bool
	for _, step := range scaleTypes[k.Scale].Steps {
		set[(k.Tonic+step)%12] = true
	}
	return set
}

func (k keyCandidate) sameKey(o keyCandidate) bool {
	return k.Tonic == o.Tonic && k.Scale == o.Scale
}

// Те же семь нот при другой тонике: A mixolydian и D major.
func (k keyCandidate) samePitchSet(o keyCandidate) bool {
	return k.pitchSet() == o.pitchSet()
}

// Все тональности по убыванию счёта.
func rankKeys(c chroma) []keyCandidate {
	var out []keyCandidate
	for tonic := 0; tonic < 12; tonic++ {
		for scale := range scaleTypes {
			k := keyCandidate{Tonic: tonic, Scale: scale}
			k.Score = correlation(c, k.template())
			out = append(out, k)
		}
	}
	sort.SliceStable(out, func(a, b int) bool { return out[a].Score > out[b].Score })
	return out
}

// Место тональности в списке, с единицы; ноль — не нашлась.
func rankOf(ranked []keyCandidate, k keyCandidate) (int, keyCandidate) {
	for i, r := range ranked {
		if r.sameKey(k) {
			return i + 1, r
		}
	}
	return 0, k
}

// Доля энергии хромы внутри семи нот лада.
func scaleShare(c chroma, k keyCandidate) float64 {
	set := k.pitchSet()
	var in float64
	for i, v := range c.normalized() {
		if set[i] {
			in += v
		}
	}
	return in
}

func correlation(a, b chroma) float64 {
	var ma, mb float64
	for i := range a {
		ma += a[i]
		mb += b[i]
	}
	ma /= 12
	mb /= 12
	var num, da, db float64
	for i := range a {
		num += (a[i] - ma) * (b[i] - mb)
		da += (a[i] - ma) * (a[i] - ma)
		db += (b[i] - mb) * (b[i] - mb)
	}
	if da == 0 || db == 0 {
		return 0
	}
	return num / math.Sqrt(da*db)
}

// Пары ступеней, которыми лады отличаются друг от друга при одной тонике.
// Мажор: 2 3 4 6 7. Миксолидийский: 2 3 4 6 b7. Дорийский: 2 b3 4 6 b7.
// Минор: 2 b3 4 b6 b7. Лидийский: 2 3 #4 6 7. Фригийский: b2 b3 4 b6 b7.
var degreePairs = [][2]int{{1, 2}, {3, 4}, {5, 6}, {8, 9}, {10, 11}}

// Опознавательная пара: две соседние ступени и их доли в хроме.
type degreePair struct {
	Low, High           int
	LowShare, HighShare float64
}

func degreeEvidence(c chroma, tonic int) []degreePair {
	n := c.normalized()
	out := make([]degreePair, 0, len(degreePairs))
	for _, p := range degreePairs {
		out = append(out, degreePair{
			Low:       p[0],
			High:      p[1],
			LowShare:  n[(tonic+p[0])%12],
			HighShare: n[(tonic+p[1])%12],
		})
	}
	return out
}

// Тональность в тексте промпта: «in E major», «A mixolydian», «F sharp minor».
// Буква ноты — только заглавная, иначе «a major hit» сойдёт за ля мажор.
var keyInText = regexp.MustCompile(`\b([A-G])(#|♯|b|♭|-sharp|-flat|\s+sharp|\s+flat)?\s+(?i:(harmonic\s+minor|major|minor|ionian|dorian|phrygian|lydian|mixolydian|aeolian|locrian))\b`)

var letterPitch = map[string]int{"C": 0, "D": 2, "E": 4, "F": 5, "G": 7, "A": 9, "B": 11}

var scaleAliases = map[string]string{"ionian": "major", "aeolian": "minor"}

func parseKey(text string) (keyCandidate, bool) {
	m := keyInText.FindStringSubmatch(text)
	if m == nil {
		return keyCandidate{}, false
	}

	tonic := letterPitch[m[1]]
	switch strings.TrimSpace(strings.TrimPrefix(m[2], "-")) {
	case "#", "♯", "sharp":
		tonic++
	case "b", "♭", "flat":
		tonic--
	}
	tonic = (tonic + 12) % 12

	name := strings.Join(strings.Fields(strings.ToLower(m[3])), " ")
	if alias, ok := scaleAliases[name]; ok {
		name = alias
	}
	for i, s := range scaleTypes {
		if s.Name == name {
			return keyCandidate{Tonic: tonic, Scale: i}, true
		}
	}
	return keyCandidate{}, false
}

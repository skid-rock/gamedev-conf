package main

// Концовка трека: последние такты перед тишиной.
// Финальная нота баса — самый прямой довод о тонике. Общий профиль его не видит:
// в A02-4a минорный VI (F#m) звучит дольше тоники, а заканчивается трек на A.
type ending struct {
	Found bool
	// Окно концовки, секунды.
	From, To float64
	// Доли нот концовки: бас, а если баса нет — середина.
	Notes  chroma
	Source string
	Chord  chordGuess
}

// Самая громкая нота концовки и её доля.
func (e ending) note() (int, float64) {
	pc := e.Notes.top(1)[0]
	return pc, e.Notes[pc]
}

// Длина окна концовки: два такта, без темпа — пять секунд.
func endingWindow(bpm float64) float64 {
	if bpm <= 0 {
		return 5
	}
	return 2 * barSeconds(bpm)
}

// Найти концовку: окно заданной длины, которое кончается на последнем звучащем кадре.
// Хвост реверберации тише порога тишины и в окно не попадает.
func findEnding(frames []frameChroma, window float64) ending {
	last := -1
	for i := len(frames) - 1; i >= 0; i-- {
		if frames[i].Weight > 0 {
			last = i
			break
		}
	}
	if last < 0 {
		return ending{}
	}

	e := ending{Found: true, To: frames[last].Time, From: frames[last].Time - window}
	// Кадры входят долями, а не энергией: финальная нота часто тише предыдущего
	// такта, и доли секунды громкого соседа иначе перевешивают её целиком.
	var bass, mid chroma
	for i := last; i >= 0 && frames[i].Time >= e.From; i-- {
		if frames[i].Weight == 0 {
			continue
		}
		bass.add(frames[i].Bass.normalized(), 1)
		mid.add(frames[i].Mid.normalized(), 1)
	}

	e.Chord = guessChord(bass, mid)
	if bass.sum() > 0 {
		e.Notes, e.Source = bass.normalized(), "бас"
	} else {
		e.Notes, e.Source = mid.normalized(), "середина"
	}
	return e
}

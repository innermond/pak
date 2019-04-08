package pak

type ScoreBoardEntry struct {
	bin   *Bin
	box   *Box
	score *Score
}

func NewScoreBoardEntry(bin *Bin, box *Box) *ScoreBoardEntry {
	return &ScoreBoardEntry{bin: bin, box: box}
}

func (e *ScoreBoardEntry) calculate() *Score {
	e.score = e.bin.scoreFor(e.box)
	return e.score
}

func (e *ScoreBoardEntry) fit() bool {
	return !e.score.IsBlank()
}

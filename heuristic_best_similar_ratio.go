package pak

import "math"

type BestSimilarRatio struct{}

func (bsr *BestSimilarRatio) calculateScore(freeRect *FreeSpaceBox, rectWidth, rectHeight float64) *Score {
	leftOverHoriz := math.Abs(freeRect.W - rectWidth)
	leftOverVert := math.Abs(freeRect.H - rectHeight)
	min := int(math.Min(leftOverHoriz, leftOverVert))
	max := int(math.Max(leftOverHoriz, leftOverVert))

	rr := (freeRect.W / freeRect.H) / (rectWidth / rectHeight)
	if rr >= 1.1 || rr <= 0.9 {
		min *= 2
		max *= 2
	}

	return NewScore(min, max)
}

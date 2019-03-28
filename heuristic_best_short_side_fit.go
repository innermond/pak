package main

import "math"

type BestShortSide struct{}

func (bss *BestShortSide) calculateScore(freeRect *FreeSpaceBox, rectWidth, rectHeight float64) *Score {
	leftOverHoriz := math.Abs(freeRect.W - rectWidth)
	leftOverVert := math.Abs(freeRect.H - rectHeight)
	min := int(math.Min(leftOverHoriz, leftOverVert))
	max := int(math.Max(leftOverHoriz, leftOverVert))

	return NewScore(min, max)
}

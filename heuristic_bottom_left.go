package main

type BottomLeft struct{}

func (bl *BottomLeft) calculateScore(freeRect *FreeSpaceBox, rectWidth, rectHeight float64) *Score {
	y := int(freeRect.Y + rectHeight)
	x := int(freeRect.X)
	return NewScore(y, x)
}

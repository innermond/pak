package pak

import "math"

type BestAreaFit struct{}

func (baf *BestAreaFit) calculateScore(freeRect *FreeSpaceBox, rectWidth, rectHeight float64) *Score {
	areaFit := int(freeRect.W*freeRect.H - rectWidth*rectHeight)
	leftOverHoriz := math.Abs(freeRect.W - rectWidth)
	leftOverVert := math.Abs(freeRect.H - rectHeight)
	shortSideFit := int(math.Min(leftOverHoriz, leftOverVert))

	return NewScore(areaFit, shortSideFit)
}

package main

type Scorer interface {
	calculateScore(freeRect *FreeSpaceBox, rectWidth, rectHeight float64) *Score
}

type Base struct {
	Scorer
}

func (b *Base) FindPositionForNewNode(box *Box, freeRects []*FreeSpaceBox) *Score {
	bestScore := NewScoreBlank()
	width := box.W
	height := box.H

	for _, freeRect := range freeRects {
		b.tryPlaceRectIn(freeRect, box, width, height, bestScore)
		if box.CanRotate {
			b.tryPlaceRectIn(freeRect, box, height, width, bestScore)
		}
	}

	return bestScore
}

func (b *Base) tryPlaceRectIn(freeRect *FreeSpaceBox, box *Box, rectWidth float64, rectHeight float64, bestScore *Score) {
	if freeRect.W >= rectWidth && freeRect.H >= rectHeight {
		score := b.calculateScore(freeRect, rectWidth, rectHeight)
		if score.Bigger(bestScore) {
			box.X = freeRect.X
			box.Y = freeRect.Y
			if box.CanRotate && (box.W == rectHeight && box.H == rectWidth) {
				// box might been rotated on a previous attempt to fit into a free rect
				box.Rotated = !box.Rotated
			}
			box.W = rectWidth
			box.H = rectHeight
			box.Packed = true
			bestScore.Assign(score)
		}
	}
}

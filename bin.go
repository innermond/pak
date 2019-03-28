package main

import "fmt"

type Bin struct {
	W, H           float64
	Boxes          []*Box
	FreeRectangles []*FreeSpaceBox
	Heuristic      *Base
}

func NewBin(w float64, h float64, s *Base) *Bin {
	if s == nil {
		s = &Base{&BestAreaFit{}}
	}

	return &Bin{w, h, nil, []*FreeSpaceBox{{W: w, H: h}}, s}
}

func (b *Bin) Area() float64 {
	return b.W * b.H
}

func (b *Bin) Eficiency() float64 {
	boxesArea := 0.0
	for _, box := range b.Boxes {
		boxesArea += box.Area()
	}
	return boxesArea * 100 / b.Area()
}

func (b *Bin) Label() string {
	return fmt.Sprintf("%.2fx%.2f %.2f", b.W, b.H, b.Eficiency())
}

func (b *Bin) Insert(box *Box) bool {
	if box.Packed {
		return false
	}

	b.Heuristic.FindPositionForNewNode(box, b.FreeRectangles)
	if !box.Packed {
		return false
	}

	numRectanglesToProcess := len(b.FreeRectangles)
	i := 0
	for i < numRectanglesToProcess {
		if b.splitFreeNode(b.FreeRectangles[i], box) {
			b.FreeRectangles = append(b.FreeRectangles[:i], b.FreeRectangles[i+1:]...)
			numRectanglesToProcess--
		} else {
			i++
		}
	}
	b.pruneFreeList()
	b.Boxes = append(b.Boxes, box)
	return true
}

func (b *Bin) scoreFor(box *Box) *Score {
	copyBox := NewBox(box.W, box.H)
	score := b.Heuristic.FindPositionForNewNode(copyBox, b.FreeRectangles)
	return score
}

func (b *Bin) isLargerThan(box *Box) bool {
	return (b.W >= box.W && b.H >= box.H) || (b.H >= box.W && b.W >= box.H)
}

func (b *Bin) splitFreeNode(freeNode *FreeSpaceBox, usedNode *Box) bool {
	if usedNode.X >= freeNode.X+freeNode.W ||
		usedNode.X+usedNode.W <= freeNode.X ||
		usedNode.Y >= freeNode.Y+freeNode.H ||
		usedNode.Y+usedNode.H <= freeNode.Y {
		return false
	}

	b.trySplitFreeNodeVertically(freeNode, usedNode)
	b.trySplitFreeNodeHorizontally(freeNode, usedNode)

	return true
}

func (b *Bin) trySplitFreeNodeVertically(freeNode *FreeSpaceBox, usedNode *Box) {
	if usedNode.X < freeNode.X+freeNode.W && usedNode.X+usedNode.W > freeNode.X {
		b.tryLeaveFreeSpaceAtTop(freeNode, usedNode)
		b.tryLeaveFreeSpaceAtBottom(freeNode, usedNode)
	}
}

func (b *Bin) tryLeaveFreeSpaceAtTop(freeNode *FreeSpaceBox, usedNode *Box) {
	if usedNode.Y > freeNode.Y && usedNode.Y < freeNode.Y+freeNode.H {
		newNode := &FreeSpaceBox{freeNode.W, freeNode.H, freeNode.X, freeNode.Y}
		newNode.H = usedNode.Y - newNode.Y
		b.FreeRectangles = append(b.FreeRectangles, newNode)
	}
}

func (b *Bin) tryLeaveFreeSpaceAtBottom(freeNode *FreeSpaceBox, usedNode *Box) {
	if usedNode.Y+usedNode.H < freeNode.Y+freeNode.H {
		newNode := &FreeSpaceBox{freeNode.W, freeNode.H, freeNode.X, freeNode.Y}
		newNode.Y = usedNode.Y + usedNode.H
		newNode.H = freeNode.Y + freeNode.H - (usedNode.Y + usedNode.H)
		b.FreeRectangles = append(b.FreeRectangles, newNode)
	}
}

func (b *Bin) trySplitFreeNodeHorizontally(freeNode *FreeSpaceBox, usedNode *Box) {
	if usedNode.Y < freeNode.Y+freeNode.H && usedNode.Y+usedNode.H > freeNode.Y {
		b.tryLeaveFreeSpaceOnLeft(freeNode, usedNode)
		b.tryLeaveFreeSpaceOnRight(freeNode, usedNode)
	}
}

func (b *Bin) tryLeaveFreeSpaceOnLeft(freeNode *FreeSpaceBox, usedNode *Box) {
	if usedNode.X > freeNode.X && usedNode.X < freeNode.X+freeNode.W {
		newNode := &FreeSpaceBox{freeNode.W, freeNode.H, freeNode.X, freeNode.Y}
		newNode.W = usedNode.X - newNode.X
		b.FreeRectangles = append(b.FreeRectangles, newNode)
	}
}

func (b *Bin) tryLeaveFreeSpaceOnRight(freeNode *FreeSpaceBox, usedNode *Box) {
	if usedNode.X+usedNode.W < freeNode.X+freeNode.W {
		newNode := &FreeSpaceBox{freeNode.W, freeNode.H, freeNode.X, freeNode.Y}
		newNode.X = usedNode.X + usedNode.W
		newNode.W = freeNode.X + freeNode.W - (usedNode.X + usedNode.W)
		b.FreeRectangles = append(b.FreeRectangles, newNode)
	}
}

/**
*    * Goes through the free rectangle list and removes any redundant entries.
*       */
func (b *Bin) pruneFreeList() {
	i := 0
	for i < len(b.FreeRectangles) {
		j := i + 1
		for j < len(b.FreeRectangles) {
			if b.isContainedIn(b.FreeRectangles[i], b.FreeRectangles[j]) {
				b.FreeRectangles = append(b.FreeRectangles[:i], b.FreeRectangles[i+1:]...)
				i--
				break
			}
			if b.isContainedIn(b.FreeRectangles[j], b.FreeRectangles[i]) {
				b.FreeRectangles = append(b.FreeRectangles[:j], b.FreeRectangles[j+1:]...)
			} else {
				j++
			}
		}
		i++
	}
}

func (b *Bin) isContainedIn(rectA, rectB *FreeSpaceBox) bool {
	return rectA != nil && rectB != nil &&
		rectA.X >= rectB.X && rectA.Y >= rectB.Y &&
		rectA.X+rectA.W <= rectB.X+rectB.W &&
		rectA.Y+rectA.H <= rectB.Y+rectB.H
}

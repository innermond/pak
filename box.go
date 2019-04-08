package pak

import (
	"fmt"
	"math"
	"strconv"
)

type Box struct {
	W, H, X, Y                 float64
	Packed, CanRotate, Rotated bool
}

func NewBox(w, h float64) *Box {
	return &Box{W: w, H: h}
}

func (b *Box) Area() float64 {
	return b.W * b.H
}

func (b *Box) Rotate() {
	b.W, b.H = b.H, b.W
}

func (b *Box) Label() string {
	return fmt.Sprintf("%.2fx%.2f at [%.2f, %.2f]", b.W, b.H, b.X, b.Y)
}

func BoxCode(bb []*Box) (code string) {
	cc := map[string]int{}
	min, max := 0.0, 0.0
	lbl := ""
	for _, b := range bb {
		min = math.Min(b.W, b.H)
		max = math.Max(b.W, b.H)
		lbl = fmt.Sprintf("%.2fx%.2f", min, max)
		k, ok := cc[lbl]
		if ok {
			cc[lbl] = k + 1
		} else {
			cc[lbl] = 1
		}
	}

	for c, k := range cc {
		if k > 1 {
			code += " " + c + "x" + strconv.Itoa(k)
		} else {
			code += " " + c
		}
	}

	return
}

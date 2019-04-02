package main

import "fmt"

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

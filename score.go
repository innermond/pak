package main

import "math"

type Score struct {
	S1, S2 int
}

func NewScore(s1, s2 int) *Score {
	return &Score{s1, s2}
}

func NewScoreBlank() *Score {
	return NewScore(math.MaxInt64, math.MaxInt64)
}

func (s *Score) Bigger(o *Score) bool {
	return s.S1 < o.S1 || (s.S1 == o.S1 && s.S2 < o.S2)
}

func (s *Score) Smaller(o *Score) bool {
	return s.S1 > o.S1 || (s.S1 == o.S1 && s.S2 > o.S2)
}

func (s *Score) Equal(o *Score) bool {
	return s.S1 == o.S1 && s.S2 == o.S2
}

func (s *Score) Assign(o *Score) {
	s.S1, s.S2 = o.S1, o.S2
}

func (s *Score) IsBlank() bool {
	return s.S1 == math.MaxInt64
}

func (s *Score) DecreaseBy(delta int) {
	s.S1 += delta
	s.S2 += delta
}

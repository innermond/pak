package main

import "math"

type Packer struct {
	bins          []*Bin
	unpackedBoxes []*Box
}

type PackerOptions struct {
	limit int
}

func NewPacker(bins []*Bin) *Packer {
	return &Packer{bins: bins}
}

func (pk *Packer) pack(unfiltered []*Box, options *PackerOptions) []*Box {
	var (
		packed []*Box
		boxes  []*Box
		i      int
	)
	for _, box := range unfiltered {
		if box.Packed {
			i++
			boxes = append(boxes, box)
		}
	}
	if i == 0 {
		return packed
	}

	limit := math.MaxInt64
	if options != nil && options.limit != 0 {
		limit = options.limit
	}

	board := NewScoreBoard(pk.bins, boxes)
	var (
		entry *ScoreBoardEntry
	)
	for {
		entry = board.bestFit()
		if entry == nil {
			break
		}
		entry.bin.Insert(entry.box)
		board.removeBox(entry.box)
		board.recalculateBin(entry.bin)
		packed = append(packed, entry.box)
		if len(packed) >= limit {
			break
		}

	}
	return packed
}

package pak

type ScoreBoard struct {
	entries []*ScoreBoardEntry
}

func NewScoreBoard(bins []*Bin, boxes []*Box) *ScoreBoard {
	sb := &ScoreBoard{}
	for _, bin := range bins {
		sb.addBinEntries(bin, boxes)
	}
	return sb
}

func (sb *ScoreBoard) addBinEntries(bin *Bin, boxes []*Box) {
	for _, box := range boxes {
		entry := NewScoreBoardEntry(bin, box)
		entry.calculate()
		sb.entries = append(sb.entries, entry)
	}
}

func (sb *ScoreBoard) any() []*ScoreBoardEntry {
	return sb.entries
}

func (sb *ScoreBoard) largestNotFitingBox() *Box {
	var (
		unfit *ScoreBoardEntry
	)

	for _, entry := range sb.entries {
		if entry.fit() {
			continue
		}
		if unfit == nil || unfit.box.Area() < entry.box.Area() {
			unfit = entry
		}
	}

	return unfit.box
}

func (sb *ScoreBoard) bestFit() *ScoreBoardEntry {
	var best *ScoreBoardEntry
	for _, entry := range sb.entries {
		if !entry.fit() {
			continue
		}
		if best == nil || best.score.Smaller(entry.score) {
			best = entry
		}
	}
	return best
}

func (sb *ScoreBoard) removeBox(box *Box) {
	for i, entry := range sb.entries {
		if entry.box == box {
			sb.entries = append(sb.entries[:i], sb.entries[i+1:]...)
			break
		}
	}
}

func (sb *ScoreBoard) addBin(bin *Bin) {
	sb.addBinEntries(bin, sb.currentBoxes())
}

func (sb *ScoreBoard) recalculateBin(bin *Bin) {
	for _, entry := range sb.entries {
		if entry.bin == bin {
			entry.calculate()
		}
	}
}

func (sb *ScoreBoard) currentBoxes() []*Box {
	var boxes []*Box
	seen := make(map[*Box]struct{}, len(sb.entries))
	for _, entry := range sb.entries {
		if _, ok := seen[entry.box]; ok {
			continue
		}
		seen[entry.box] = struct{}{}
		boxes = append(boxes, entry.box)
	}
	return boxes
}

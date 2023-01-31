package rail_fence_cipher

type TextConsumer struct {
	s             string
	runeSlice     []rune
	index         int
	fillCharacter rune
}

func NewTextConsumer(s string, fillCharacter rune) *TextConsumer {
	return &TextConsumer{
		s:             s,
		runeSlice:     []rune(s),
		index:         0,
		fillCharacter: fillCharacter,
	}
}

func (x *TextConsumer) Take() rune {
	if x.index < len(x.runeSlice) {
		r := x.runeSlice[x.index]
		x.index++
		return r
	} else {
		return x.fillCharacter
	}
}

package hexgen

type Hexagram struct {
	top    int64
	bottom int64
}

func (h Hexagram) Generate() []int64 {
	hex := make([]int64, 0, 2)
	hex = append(hex, h.top, h.bottom)
	return hex
}

func NewHexagram(top, bottom int64) *Hexagram {
	return &Hexagram{
		top:    top,
		bottom: bottom,
	}
}

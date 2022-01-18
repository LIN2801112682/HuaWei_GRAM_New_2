package build_VGram_index

type Inverted_index struct {
	Sid      int
	Position int
}

func NewInverted_index(sid int, p int) *Inverted_index {
	return &Inverted_index{
		Sid:      sid,
		Position: p,
	}
}

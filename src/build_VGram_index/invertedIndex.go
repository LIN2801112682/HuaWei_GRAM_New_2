package build_VGram_index

type Inverted_index struct {
	Sid      SeriesId
	PosArray []int
}

func NewInverted_index(sid SeriesId, posArray []int) *Inverted_index {
	return &Inverted_index{
		Sid:      sid,
		PosArray: posArray,
	}
}

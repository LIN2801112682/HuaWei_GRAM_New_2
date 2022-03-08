package new_precise_query

import "build_VGram_index"

type PosList struct {
	Sid      build_VGram_index.SeriesId
	PosArray []int
}

func NewPosList(sid build_VGram_index.SeriesId, posArray []int) PosList {
	return PosList{
		Sid:      sid,
		PosArray: posArray,
	}
}

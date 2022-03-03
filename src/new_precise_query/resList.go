package new_precise_query

import "build_VGram_index"

type ResList struct {
	Sid      build_VGram_index.SeriesId
	PosArray []int
}

func NewResList(sid build_VGram_index.SeriesId, posArray []int) ResList {
	return ResList{
		Sid:      sid,
		PosArray: posArray,
	}
}

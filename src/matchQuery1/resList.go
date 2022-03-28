package matchQuery1

import "index07"

type ResList struct {
	Sid      index07.SeriesId
	PosArray []int
}

func NewResList(sid index07.SeriesId, posArray []int) ResList {
	return ResList{
		Sid:      sid,
		PosArray: posArray,
	}
}

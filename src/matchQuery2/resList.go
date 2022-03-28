package matchQuery2

import "index07"

type ResList struct {
	sid      index07.SeriesId
	posArray []int
}

func NewResList(sid index07.SeriesId, posArray []int) ResList {
	return ResList{
		sid:      sid,
		posArray: posArray,
	}
}

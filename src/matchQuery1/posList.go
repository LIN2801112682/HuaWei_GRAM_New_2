package matchQuery1

import "index07"

type PosList struct {
	sid      index07.SeriesId
	posArray []int
}

func (p *PosList) Sid() index07.SeriesId {
	return p.sid
}

func (p *PosList) SetSid(sid index07.SeriesId) {
	p.sid = sid
}

func (p *PosList) PosArray() []int {
	return p.posArray
}

func (p *PosList) SetPosArray(posArray []int) {
	p.posArray = posArray
}

func NewPosList(sid index07.SeriesId, posArray []int) PosList {
	return PosList{
		sid:      sid,
		posArray: posArray,
	}
}

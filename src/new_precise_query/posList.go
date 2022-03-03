package new_precise_query

type PosList struct {
	Kid      int
	PosArray []int
}

func NewPosList(kid int, posArray []int) PosList {
	return PosList{
		Kid:      kid,
		PosArray: posArray,
	}
}

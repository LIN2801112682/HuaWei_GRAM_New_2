package index07

type SeriesId struct {
	id   int32
	time int64
}

func (s *SeriesId) Id() int32 {
	return s.id
}

func (s *SeriesId) SetId(id int32) {
	s.id = id
}

func (s *SeriesId) Time() int64 {
	return s.time
}

func (s *SeriesId) SetTime(time int64) {
	s.time = time
}

func NewSeriesId(id int32, t int64) *SeriesId {
	return &SeriesId{
		id:   id,
		time: t,
	}
}

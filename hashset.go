package hashset

type Int64Set map[int64]struct{}

func NewInt64() *Int64Set {
	return &Int64Set{}
}

func (s *Int64Set) Add(value int64) bool {
	if s.Contains(value) {
		return false
	}
	(*s)[value] = struct{}{}
	return true
}

func (s *Int64Set) Contains(value int64) bool {
	if _, ok := (*s)[value]; ok {
		return true
	}
	return false
}

func (s *Int64Set) Remove(value int64) bool {
	if s.Contains(value) {
		delete(*s, value)
		return true
	}
	return false
}

func (s *Int64Set) Range(f func(value int64) bool) {
	for k := range *s {
		f(k)
	}
}

func (s *Int64Set) Len() int {
	return int(len(*s))
}

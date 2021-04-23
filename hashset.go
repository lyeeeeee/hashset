package hashset

type Int64Set map[int64]struct{}

func NewInt64() *Int64Set {
	return &Int64Set{}
}

func (s *Int64Set) Add(value int64) {
	(*s)[value] = struct{}{}
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
		if !f(k) {
			break
		}
	}
}

func (s *Int64Set) Len() int {
	return len(*s)
}

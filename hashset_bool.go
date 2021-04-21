package hashset

type Int64SetBool map[int64]bool

func NewInt64Bool() *Int64SetBool {
	return &Int64SetBool{}
}

func (s *Int64SetBool) Add(value int64) bool {
	if s.Contains(value) {
		return false
	}
	(*s)[value] = true
	return true
}

func (s *Int64SetBool) Contains(value int64) bool {
	if _, ok := (*s)[value]; ok {
		return true
	}
	return false
}

func (s *Int64SetBool) Remove(value int64) bool {
	if s.Contains(value) {
		delete(*s, value)
		return true
	}
	return false
}

func (s *Int64SetBool) Range(f func(value int64) bool) {
	for k := range *s {
		f(k)
	}
}

func (s *Int64SetBool) Len() int {
	return int(len(*s))
}

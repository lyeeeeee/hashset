package internal

type Int64SetBool map[int64]bool

func NewInt64Bool() *Int64SetBool {
	return &Int64SetBool{}
}

func (s *Int64SetBool) Add(value int64) bool {
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
	delete(*s, value)
	return true
}

func (s *Int64SetBool) Range(f func(value int64) bool) {
	for k := range *s {
		if !f(k) {
			break
		}
	}
}

func (s *Int64SetBool) Len() int {
	return len(*s)
}

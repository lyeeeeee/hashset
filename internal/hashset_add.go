package internal

type Int64SetAdd map[int64]struct{}

func NewInt64Add() *Int64SetAdd {
	return &Int64SetAdd{}
}

func (s *Int64SetAdd) Add(value int64) bool {
	if s.Contains(value) {
		return true
	}
	(*s)[value] = struct{}{}
	return true
}

func (s *Int64SetAdd) Contains(value int64) bool {
	if _, ok := (*s)[value]; ok {
		return true
	}
	return false
}

func (s *Int64SetAdd) Remove(value int64) bool {
	if s.Contains(value) {
		delete(*s, value)
		return true
	}
	return false
}

func (s *Int64SetAdd) Range(f func(value int64) bool) {
	for k := range *s {
		if !f(k) {
			break
		}
	}
}

func (s *Int64SetAdd) Len() int {
	return len(*s)
}

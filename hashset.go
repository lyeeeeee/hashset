package main

type Int64Set struct {
	set map[int64]struct{}
	length int64
}

func NewInt64() *Int64Set {
	result := Int64Set{make(map[int64]struct{}), 0}
	return &result
}

func (s *Int64Set) Add(value int64) bool {
	if s.Contains(value) {
		return false
	}
	(*s).length++
	(*s).set[value] = struct{}{}
	return true
}

func (s *Int64Set) Contains(value int64) bool {
	if _, ok := (*s).set[value]; ok {
		return true
	}
	return false
}

func (s *Int64Set) Remove(value int64) bool {
	if s.Contains(value) {
		delete((*s).set, value)
		(*s).length--
		return true
	}
	return false
}

func (s *Int64Set) Range(f func(value int64) bool) {
	for k,_:=range (*s).set {
		f(k)
	}
}

func (s *Int64Set) Len() int {
	return int(len((*s).set))
}
package main

import (
	"github.com/zhangyunhao116/wyhash"
)

func hash(s string) uint64 {
	return wyhash.Sum64String(s)
}

type StringSet struct {
	set map[uint64]string
	length int64
}

func NewString() *StringSet {
	result := StringSet{make(map[uint64]string), 0}
	return &result
}

func (s *StringSet) Add(value string) bool {
	if s.Contains(value) {
		return false
	}
	(*s).length++
	(*s).set[hash(value)] = value
	return true
}

func (s *StringSet) Contains(value string) bool {
	if _, ok := (*s).set[hash(value)]; ok {
		return true
	}
	return false
}

func (s *StringSet) Remove(value string) bool {

	if s.Contains(value) {
		delete((*s).set, hash(value))
		(*s).length--
		return true
	}
	return false
}

func (s *StringSet) Range(f func(value string) bool) {
	for _,v:=range (*s).set {
		f(v)
	}
}

func (s *StringSet) Len() int {
	return int(len((*s).set))
}
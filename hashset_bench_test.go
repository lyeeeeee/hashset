package hashset

import (
	"math/rand"
	"testing"
)

const benchtime = 50000000
var lst [benchtime] int64

func init(){
	for i:=0; i<benchtime; i++{
		lst[i] = rand.Int63()
	}
}

func BenchmarkBool(b *testing.B){
	b.ResetTimer()
	l := NewInt64Bool()
	for i:=0;i<b.N;i++{
		l.Add(lst[i])
	}
}

func BenchmarkEmptyStruct(b *testing.B){
	b.ResetTimer()
	l := NewInt64()
	for i:=0;i<b.N;i++{
		l.Add(lst[i])
	}
}
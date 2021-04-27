package hashset

import (
	"./internal"
	"math/rand"
	"testing"
)

const capacity = 10000000

var randomList [capacity]int64

func init() {
	for i := 0; i < capacity; i++ {
		randomList[i] = int64(rand.Int63())
	}
}
func BenchmarkValueAsBool(b *testing.B) {
	b.ResetTimer()

	l := internal.NewInt64Bool()
	for n := 0; n < b.N; n++ {
		l.Add(randomList[n%capacity])
	}
}

func BenchmarkValueAsEmptyStruct(b *testing.B) {
	b.ResetTimer()
	l := NewInt64()
	for n := 0; n < b.N; n++ {
		l.Add(randomList[n%capacity])
	}
}
func BenchmarkAddAfterContains(b *testing.B) {
	b.ResetTimer()
	l := internal.NewInt64Add()
	for n := 0; n < b.N; n++ {
		l.Add(randomList[n%capacity])
	}
}

func BenchmarkAddWithoutContains(b *testing.B) {
	b.ResetTimer()
	l := NewInt64()
	for n := 0; n < b.N; n++ {
		l.Add(randomList[n%capacity])
	}
}

func BenchmarkRemoveAfterContains_Missing(b *testing.B) {
	l := internal.NewInt64Add()
	for n := 0; n < b.N; n++ {
		l.Add(randomList[n%capacity])
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		l.Remove(int64(rand.Int63()))
	}
}

func BenchmarkRemoveWithoutContains_Missing(b *testing.B) {

	l := NewInt64()
	for n := 0; n < b.N; n++ {
		l.Add(randomList[n%capacity])
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		l.Remove(int64(rand.Int63()))
	}
}

func BenchmarkRemoveAfterContains_Hitting(b *testing.B) {
	l := internal.NewInt64Add()
	for n := 0; n < b.N; n++ {
		l.Add(randomList[n%capacity])
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Remove(randomList[n%capacity])
	}
}

func BenchmarkRemoveWithoutContains_Hitting(b *testing.B) {
	l := NewInt64()
	for n := 0; n < b.N; n++ {
		l.Add(randomList[n%capacity])
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Remove(randomList[n%capacity])
	}
}

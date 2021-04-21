package hashset

import (
	"fmt"
	"testing"
)

func Example() {
	l := NewInt64()

	for _, v := range []int{10, 12, 15} {
		if l.Add(int64(v)) {
			fmt.Println("skipset add", v)
		}
	}

	if l.Contains(10) {
		fmt.Println("skipset contains 10")
	}

	l.Range(func(value int64) bool {
		fmt.Println("skipset range found ", value)
		return true
	})

	l.Remove(15)
	fmt.Printf("skipset contains %d items\r\n", l.Len())
}

func TestIntSet(t *testing.T) {
	// Correctness.
	l := NewInt64()
	if l.length != 0 {
		t.Fatal("invalid length")
	}
	if l.Contains(0) {
		t.Fatal("invalid contains")
	}

	if !l.Add(0) || l.length != 1 {
		t.Fatal("invalid add")
	}
	if !l.Contains(0) {
		t.Fatal("invalid contains")
	}
	if !l.Remove(0) || l.length != 0 {
		t.Fatal("invalid remove")
	}

	if !l.Add(20) || l.length != 1 {
		t.Fatal("invalid add")
	}
	if !l.Add(22) || l.length != 2 {
		t.Fatal("invalid add")
	}
	if !l.Add(21) || l.length != 3 {
		t.Fatal("invalid add")
	}
	if l.Add(21) {
		t.Fatal(l.length, " invalid add")
	}
}

func TestStringSet(t *testing.T) {
	x := NewString()
	if !x.Add("111") || x.Len() != 1 {
		t.Fatal("invalid")
	}
	if !x.Add("222") || x.Len() != 2 {
		t.Fatal("invalid")
	}
	if x.Add("111") || x.Len() != 2 {
		t.Fatal("invalid")
	}
	if !x.Contains("111") || !x.Contains("222") {
		t.Fatal("invalid")
	}
	if !x.Remove("111") || x.Len() != 1 {
		t.Fatal("invalid")
	}
	if !x.Remove("222") || x.Len() != 0 {
		t.Fatal("invalid")
	}
}

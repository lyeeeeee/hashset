package hashset

import (
	"fmt"
	"reflect"
	"testing"
)

func Example() {
	l := NewInt64()

	for _, v := range []int{10, 12, 15} {
		l.Add(int64(v))
	}

	if l.Contains(10) {
		fmt.Println("hashset contains 10")
	}

	l.Range(func(value int64) bool {
		fmt.Println("hashset range found ", value)
		return true
	})

	l.Remove(15)
	fmt.Printf("hashset contains %d items\r\n", l.Len())
}

func TestIntSet(t *testing.T) {
	// Correctness.
	l := NewInt64()
	if l.Len() != 0 {
		t.Fatal("invalid length")
	}
	if l.Contains(0) {
		t.Fatal("invalid contains")
	}

	if l.Add(0); l.Len() != 1 {
		t.Fatal("invalid add")
	}
	if !l.Contains(0) {
		t.Fatal("invalid contains")
	}
	if l.Remove(0); l.Len() != 0 {
		t.Fatal("invalid remove")
	}

	if l.Add(20); l.Len() != 1 {
		t.Fatal("invalid add")
	}
	if l.Add(22); l.Len() != 2 {
		t.Fatal("invalid add")
	}
	if l.Add(21); l.Len() != 3 {
		t.Fatal("invalid add")
	}
	if l.Add(21); l.Len() != 3 {
		t.Fatal(l.Len(), " invalid add")
	}
}

func TestIntSetIntersection(t *testing.T) {
	l1 := NewInt64()
	l1.Add(99)
	l1.Add(88)
	l1.Add(77)

	l2 := NewInt64()
	l2.Add(77)
	l2.Add(66)
	l2.Add(55)

	expected := NewInt64()
	expected.Add(77)

	if res := l1.Intersection(l2); !reflect.DeepEqual(expected, res){
		t.Fatal(" invalid intersection")
	}
}

func TestIntSetUnion(t *testing.T) {
	l1 := NewInt64()
	l1.Add(99)
	l1.Add(88)
	l1.Add(77)

	l2 := NewInt64()
	l2.Add(77)
	l2.Add(66)
	l2.Add(55)

	expected := NewInt64()
	expected.Add(99)
	expected.Add(88)
	expected.Add(77)
	expected.Add(66)
	expected.Add(55)

	if res := l1.Union(l2); !reflect.DeepEqual(expected, res){
		t.Fatal(" invalid intersection")
	}
}

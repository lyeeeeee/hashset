# hashset

## Introduction
In this repository, we implemented one foundational data structure: Set based on Map in golang. We have:  
`Add(value int64)`: Adds the specified element to this set.  
`Contains(value int64) bool`: Returns true if this set contains the specified element.  
`Remove(value int64)`: Removes the specified element from this set.  
`Range(f func(value int64) bool)`: Function f executes by taking element in the set as parameter sequentially until f returns false  
`Len() int`: Returns the number of elements of this set.  

We made two experiments in order to measure the overall performance of the new hashset:  
1. the chosen value's type: empty struct vs. bool  
2. the impact of checking the existence of the key before add/remove an item  

## Features
- The API of hashset is totally compatible with [skipset](https://github.com/zhangyunhao116/skipset/)
- Usually, developers implement the set in golang by setting the value of <key,value> pair to `bool` or `int`. However, We proved \
that using empty struct is more space efficiency and slightly time efficiency. 


## When to use hashset
Hashset **doesnt** guarantee concurrent safe. If you do need a concurrent safe set, go for [skipset](https://github.com/zhangyunhao116/skipset/)

## Quickstart
```
package main

import (
	"fmt"
    "github.com/lyeeeeee/hashset"
)

func main() {
	l := NewInt()

	for _, v := range []int{10, 12, 15} {
		if l.Add(v) {
			fmt.Println("hashset add", v)
		}
	}

	if l.Contains(10) {
		fmt.Println("hashset contains 10")
	}

	l.Range(func(value int) bool {
		fmt.Println("hashset range found ", value)
		return true
	})

	l.Remove(15)
	fmt.Printf("hashset contains %d items\r\n", l.Len())
}
```

## Benchmark
go version: go1.15.10 linux/amd64  
CPU: Intel(R) Xeon(R) Platinum 8260 CPU @ 2.40GHz (4C8T)  
OS: Debian 4.14.81.bm.15  
MEMORY: 16G  

```
$ go test -run=None -bench=. -benchtime=1000000x -benchmem -count=10 -cpu=4 1000000x20x4.txt
$ benchstat 1000000x20x4.txt
name                             time/op
ValueAsBool-4                    325ns ±19%
ValueAsEmptyStruct-4             312ns ± 8%
AddAfterContains-4               362ns ±13%
AddWithoutContains-4             317ns ± 8%
RemoveAfterContains_Missing-4    177ns ± 3%
RemoveWithoutContains_Missing-4  148ns ±21%
RemoveAfterContains_Hitting-4    207ns ± 6%
RemoveWithoutContains_Hitting-4  128ns ±21%

name                             alloc/op
ValueAsBool-4                    54.0B ± 0%
ValueAsEmptyStruct-4             49.0B ± 0%
AddAfterContains-4               49.0B ± 0%
AddWithoutContains-4             49.0B ± 0%
RemoveAfterContains_Missing-4    0.00B
RemoveWithoutContains_Missing-4  0.00B
RemoveAfterContains_Hitting-4    0.00B
RemoveWithoutContains_Hitting-4  0.00B
```

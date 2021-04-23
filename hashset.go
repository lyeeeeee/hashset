package hashset

type Int64Set map[int64]struct{}

// NewInt64 returns an empty int64 set
func NewInt64() Int64Set {
	return make(map[int64]struct{})
}

// NewInt64WithSize returns an empty int64 set initialized with specific size
func NewInt64WithSize(size int) Int64Set {
	return make(map[int64]struct{}, size)
}

//Add adds the specified element to this set
//Always returns true due to the build-in map doesnt indicate caller whether the given element already exists
//Reserves the return type for future extension
func (s Int64Set) Add(value int64) bool {
	s[value] = struct{}{}
	return true
}

//Contains returns true if this set contains the specified element
func (s Int64Set) Contains(value int64) bool {
	if _, ok := s[value]; ok {
		return true
	}
	return false
}

//Remove removes the specified element from this set
//Always returns true due to the build-in map doesnt indicate caller whether the given element already exists
//Reserves the return type for future extension
func (s Int64Set) Remove(value int64) bool {
	delete(s, value)
	return true
}

//Function f executes by taking element in the set as parameter sequentially until f returns false
func (s Int64Set) Range(f func(value int64) bool) {
	for k := range s {
		if !f(k) {
			break
		}
	}
}

//Returns the number of elements of this set
func (s Int64Set) Len() int {
	return len(s)
}

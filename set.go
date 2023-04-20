package helper

// A set stores elements uniquely, without duplication
type Set map[interface{}]struct{}

// Add an element to the set.
//
// Complexity: O(1) average, O(n) worst case.
func (s Set) Add(e interface{}) {
	s[e] = struct{}{}
}

// Remove an element from the set.
//
// Complexity: O(1) average, O(n) worst case.
func (s Set) Remove(e interface{}) {
	delete(s, e)
}

// Remove an element from the set.
//
// Complexity: O(1) average, O(n) worst case.
func (s Set) Contains(e interface{}) bool {
	_, ok := s[e]
	return ok
}

// Size returns the number of elements in the set.
func (s Set) Size() int {
	return len(s)
}

// All returns all the elements of the set as a slice.
func (s Set) All() []interface{} {
	ems := make([]interface{}, 0, len(s))
	for e := range s {
		ems = append(ems, e)
	}
	return ems
}

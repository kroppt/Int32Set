package set

import (
	"sort"
	"strconv"
	"strings"
)

// Set is an integer set
type Set struct {
	m map[int32]struct{}
}

// NewSet returns a pointer to an empty set
func NewSet() Set {
	m := make(map[int32]struct{})
	return Set{m}
}

// NewSetInit returns a pointer to a set with the passed integer
func NewSetInit(i int32) (ns Set) {
	ns = NewSet()
	ns.add(i)
	return ns
}

// Add inserts the given integer into the set if it does not already exist
func (s Set) Add(i int32) {
	if s.Contains(i) {
		return
	}
	s.add(i)
}

// Remove deletes the given integer from the set if it exists
func (s Set) Remove(i int32) {
	if !s.Contains(i) {
		return
	}
	delete(s.m, i)
}

// Union merges the second set onto the first and returns whether the set changed
func (s Set) Union(os Set) bool {
	b := false
	for i := range os.m {
		if s.Contains(i) {
			continue
		}
		s.add(i)
		b = true
	}
	return b
}

// Copy returns a duplicate set
func (s Set) Copy() (ns Set) {
	ns = NewSet()
	for i := range s.m {
		ns.add(i)
	}
	return ns
}

// Range executes the ranging function for each element
func (s Set) Range(f func(int32) bool) {
	for i := range s.m {
		if !f(i) {
			break
		}
	}
}

// Contains returns whether the set contains the integer
func (s Set) Contains(i int32) bool {
	_, ok := s.m[i]
	return ok
}

// IsEmpty returns true if the set size is zero
func (s Set) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns the number of items in the set
func (s Set) Size() int32 {
	return int32(len(s.m))
}

// Equals returns whether the sets are equal
func (s Set) Equals(os Set) bool {
	if s.Size() != os.Size() {
		return false
	}
	equal := true
	s.Range(func(i int32) bool {
		if !os.Contains(i) {
			equal = false
			return false
		}
		return true
	})
	return equal
}

type int32s []int32

func (a int32s) Len() int           { return len(a) }
func (a int32s) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int32s) Less(i, j int) bool { return a[i] < a[j] }

// Print returns a string representation of the set
func (s Set) Print() string {
	ints := make([]int32, 0)
	for i := range s.m {
		ints = append(ints, i)
	}
	sort.Sort(int32s(ints))
	var sb strings.Builder
	sb.WriteString("{ ")
	for _, i := range ints {
		sb.WriteString(strconv.Itoa(int(i)))
		sb.WriteRune(' ')
	}
	sb.WriteString("}")
	return sb.String()
}

func (s Set) add(i int32) {
	s.m[i] = struct{}{}
}

package set_test

import (
	"fmt"
	"testing"

	"alon.kr/x/set"
	"github.com/stretchr/testify/assert"
)

func Example() {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Remove(2)

	if !s.Contains(2) {
		fmt.Println("does not contains 2")
	}

	if s.Contains(3) {
		fmt.Println("contains 3")
	}

	if len(s) == 2 {
		fmt.Println("length is 2")
	}

	// Output:
	// does not contains 2
	// contains 3
	// length is 2
}

func TestToSlice(t *testing.T) {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Remove(2)

	assert.ElementsMatch(t, s.ToSlice(), []int{1, 3})
}

func TestUnion(t *testing.T) {
	s1 := set.New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.New[int]()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)

	union := s1.Union(s2)
	assert.ElementsMatch(t, union.ToSlice(), []int{1, 2, 3, 4, 5})

	// Test empty set union
	empty := set.New[int]()
	assert.ElementsMatch(t, s1.Union(empty).ToSlice(), s1.ToSlice())
	assert.ElementsMatch(t, empty.Union(s1).ToSlice(), s1.ToSlice())
}

func TestDifference(t *testing.T) {
	s1 := set.New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.New[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	diff := s1.Difference(s2)
	assert.ElementsMatch(t, diff.ToSlice(), []int{1})

	// Test empty set difference
	empty := set.New[int]()
	assert.ElementsMatch(t, s1.Difference(empty).ToSlice(), s1.ToSlice())
	assert.Empty(t, empty.Difference(s1).ToSlice())
}

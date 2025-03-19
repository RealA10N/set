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

func TestFromSlice(t *testing.T) {
	// Test basic slice conversion
	slice := []int{1, 2, 3, 4}
	s := set.FromSlice(slice)
	assert.Equal(t, 4, len(s))
	for _, v := range slice {
		assert.True(t, s.Contains(v))
	}

	// Test with duplicates
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	sDup := set.FromSlice(duplicates)
	assert.Equal(t, 4, len(sDup))
	assert.ElementsMatch(t, sDup.ToSlice(), []int{1, 2, 3, 4})

	// Test with empty slice
	empty := []int{}
	sEmpty := set.FromSlice(empty)
	assert.Equal(t, 0, len(sEmpty))
	assert.Empty(t, sEmpty.ToSlice())

	// Test with string slice
	strSlice := []string{"a", "b", "c"}
	sStr := set.FromSlice(strSlice)
	assert.Equal(t, 3, len(sStr))
	for _, v := range strSlice {
		assert.True(t, sStr.Contains(v))
	}
}

func TestCopy(t *testing.T) {
	original := set.New[int]()
	original.Add(1)
	original.Add(2)
	original.Add(3)

	copied := original.Copy()

	assert.ElementsMatch(t, copied.ToSlice(), original.ToSlice())

	original.Add(4)
	assert.NotContains(t, copied.ToSlice(), 4)

	copied.Add(5)
	assert.NotContains(t, original.ToSlice(), 5)
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

func TestIntersection(t *testing.T) {
	s1 := set.New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.New[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	assert.ElementsMatch(t, s1.Intersection(s2).ToSlice(), []int{2, 3})
	assert.ElementsMatch(t, s2.Intersection(s1).ToSlice(), []int{2, 3})

	// Test empty intersection
	s3 := set.New[int]()
	s3.Add(5)
	s3.Add(6)
	assert.Empty(t, s1.Intersection(s3).ToSlice())
	assert.Empty(t, s3.Intersection(s1).ToSlice())

	// Test with empty set
	empty := set.New[int]()
	assert.Empty(t, s1.Intersection(empty).ToSlice())
	assert.Empty(t, empty.Intersection(s1).ToSlice())

	// Test with identical sets
	s4 := s1.Copy()
	assert.ElementsMatch(t, s1.Intersection(s4).ToSlice(), s1.ToSlice())
	assert.ElementsMatch(t, s4.Intersection(s1).ToSlice(), s1.ToSlice())
}

func TestEquals(t *testing.T) {
	s1 := set.New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	// Test equal sets
	s2 := set.New[int]()
	s2.Add(2)
	s2.Add(1)
	s2.Add(3)
	assert.True(t, s1.Equals(s2))
	assert.True(t, s2.Equals(s1))

	// Test unequal sets with same length
	s3 := set.New[int]()
	s3.Add(1)
	s3.Add(2)
	s3.Add(4)
	assert.False(t, s1.Equals(s3))
	assert.False(t, s3.Equals(s1))

	// Test unequal sets with different lengths
	s4 := set.New[int]()
	s4.Add(1)
	s4.Add(2)
	assert.False(t, s1.Equals(s4))
	assert.False(t, s4.Equals(s1))

	// Test with empty sets
	empty1 := set.New[int]()
	empty2 := set.New[int]()
	assert.True(t, empty1.Equals(empty2))
	assert.False(t, s1.Equals(empty1))
	assert.False(t, empty1.Equals(s1))
}

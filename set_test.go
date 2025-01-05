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

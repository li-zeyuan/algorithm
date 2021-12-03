package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversedSlice(t *testing.T) {
	strList := []string{"A", "B", "C", "D"}
	sList := ReversedSlice(strList)
	fmt.Println(sList)
}

func TestReverseString(t *testing.T) {
	var l1 []byte
	reverseString(l1)
	assert.Equal(t, l1, []byte{})

	l2 := []byte{96, 97, 98}
	reverseString(l1)
	assert.Equal(t, l2, []byte{98, 97, 96})

	l3 := []byte{96, 97, 98, 99}
	reverseString(l1)
	assert.Equal(t, l3, []byte{99, 98, 97, 96})
}

func TestIntersection(t *testing.T) {
	l1 := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	assert.Equal(t, l1[0], 2)
}

func TestIntersect(t *testing.T) {
	l1 := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	assert.Equal(t, len(l1), 2)
	assert.Equal(t, l1[0], 2)
	assert.Equal(t, l1[1], 2)

	l2 := intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	assert.Equal(t, len(l2), 2)
	assert.Equal(t, l2[0], 4)
	assert.Equal(t, l2[1], 9)
}

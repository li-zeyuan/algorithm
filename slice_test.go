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

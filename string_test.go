package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"fmt"
)

func TestIsUniqueString(t *testing.T) {
	s := "aabc"
	fmt.Println(isUniqueString(s))
}

func TestReversedString(t *testing.T) {
	fmt.Println(ReversedString("lizeyuan"))
}

func TestIsRegroup(t *testing.T) {
	s1 := "lizeyuan"
	s2 := "liyuanza"
	fmt.Println(IsRegroup(s1, s2))
}

func TestReverseVowels(t *testing.T) {
	assert.Equal(t, reverseVowels("race car"), "race car")
	assert.Equal(t, reverseVowels("hello"), "holle")
	assert.Equal(t, reverseVowels("leetcode"), "leotcede")
}

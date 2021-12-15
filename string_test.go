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

func TestFirstUniqChar(t *testing.T) {
	assert.Equal(t, firstUniqChar("leetcode"), 0)
	assert.Equal(t, firstUniqChar("loveleetcode"), 2)
}

func TestFindTheDifference(t *testing.T) {
	assert.Equal(t, findTheDifference("abcd", "abcde"), byte('e'))
	assert.Equal(t, findTheDifference("", "y"), byte('y'))
}

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, isSubsequence("abc", "ahbgdc"), true)
	assert.Equal(t, isSubsequence("axc", "ahbgdc"), false)
	assert.Equal(t, isSubsequence("aaaaaa", "bbaaaa"), false)
}

func TestRev(t *testing.T) {
	assert.Equal(t, rev(""), 0)
	assert.Equal(t, rev("a"), 1)
	assert.Equal(t, rev("aa"), 2)
	assert.Equal(t, rev("aab"), 3)
	assert.Equal(t, rev("aabb"), 4)
	assert.Equal(t, rev("aabbcd"), 5)
	assert.Equal(t, rev("aabbcccd"), 7)
}

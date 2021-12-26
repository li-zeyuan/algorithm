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

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, longestPalindrome("a"), 1)
	assert.Equal(t, longestPalindrome( ""), 0)
	assert.Equal(t, longestPalindrome( "ab"), 1)
	assert.Equal(t, longestPalindrome( "abccccdd"), 7)
	assert.Equal(t, longestPalindrome( "ccd"), 3)
	assert.Equal(t, longestPalindrome( "vv"), 2)
}

func TestAddStrings(t *testing.T) {
	assert.Equal(t, addStrings("11", "123"), "134")
	assert.Equal(t, addStrings("456", "77"), "533")
	assert.Equal(t, addStrings("0", "0"), "0")
}

func TestRepeatedSubstringPattern(t *testing.T) {
	assert.Equal(t, repeatedSubstringPattern("ababac"), false)
	assert.Equal(t, repeatedSubstringPattern("babbabbabbabbab"), true)
	assert.Equal(t, repeatedSubstringPattern("abab"), true)
	assert.Equal(t, repeatedSubstringPattern("aba"), false)
	assert.Equal(t, repeatedSubstringPattern("abcabcabcabc"), true)
}
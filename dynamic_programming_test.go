package algorithm

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, uniquePaths(3, 7), 28)
	assert.Equal(t, uniquePaths(3, 2), 3)
}

func TestNumWays(t *testing.T) {
	assert.Equal(t, numWays(3), 3)
}

func TestLengthOfLIS(t *testing.T) {
	assert.Equal(t, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), 4)
}

func TestLongestCommonSubsequence(t *testing.T) {
	assert.Equal(t, longestCommonSubsequence("abcde", "ace"), 3)
	assert.Equal(t, longestCommonSubsequence("", "ace"), 0)
}

func TestWordBreak(t *testing.T) {
	assert.Equal(t, wordBreak("aaaaaaa", []string{"aaaa","aaa"}), true)
	assert.Equal(t, wordBreak("leetcode", []string{"leet", "code"}), true)
	assert.Equal(t, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), false)
}


func TestNumberOfLines(t *testing.T) {
	t.Log(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	t.Log(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

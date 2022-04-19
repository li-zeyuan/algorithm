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
	assert.Equal(t, wordBreak("aaaaaaa", []string{"aaaa", "aaa"}), true)
	assert.Equal(t, wordBreak("leetcode", []string{"leet", "code"}), true)
	assert.Equal(t, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), false)
}

func TestNumberOfLines(t *testing.T) {
	t.Log(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	t.Log(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

func TestJump(t *testing.T) {
	assert.Equal(t, jump([]int{2, 3, 1, 1, 4}), 2)
}

func TestJump2(t *testing.T) {
	assert.Equal(t, jump2([]int{2, 3, 1, 1, 4}), 2)
}

func TestShortestToChar(t *testing.T) {
	t.Log(shortestToChar("loveleetcode", 'e'))
}

func TestShortestToChar2(t *testing.T) {
	t.Log(shortestToChar2("loveleetcode", 'e'))
}

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	assert.Equal(t, uniquePathsWithObstacles(obstacleGrid), 2)

	obstacleGrid2 := [][]int{{0, 1}, {0, 0}}
	assert.Equal(t, uniquePathsWithObstacles(obstacleGrid2), 1)

	obstacleGrid3 := [][]int{{1}}
	assert.Equal(t, uniquePathsWithObstacles(obstacleGrid3), 0)
}

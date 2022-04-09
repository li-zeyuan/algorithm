package algorithm

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestUniquePaths(t *testing.T)  {
	assert.Equal(t, uniquePaths(3,7), 28)
	assert.Equal(t, uniquePaths(3,2), 3)
}

func TestWordBreak(t *testing.T) {
	assert.Equal(t, wordBreak("aaaaaaa", []string{"aaaa","aaa"}), true)
	assert.Equal(t, wordBreak("leetcode", []string{"leet", "code"}), true)
	assert.Equal(t, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), false)
}


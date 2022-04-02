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

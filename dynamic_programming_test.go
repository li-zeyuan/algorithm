package algorithm

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestUniquePaths(t *testing.T)  {
	assert.Equal(t, uniquePaths(3,7), 28)
	assert.Equal(t, uniquePaths(3,2), 3)
}

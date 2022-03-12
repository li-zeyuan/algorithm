package algorithm

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestDivide(t *testing.T)  {
	assert.Equal(t, divide(-2147483648,-1), 2147483647)
	assert.Equal(t, divide(1, 1), 1)
	assert.Equal(t, divide(10, 3), 3)
	assert.Equal(t, divide(7, -3), -2)
}
func TestMyPow(t *testing.T)  {
	assert.Equal(t, myPow(2.00000,  10), 1024.00000)
}

package algorithm

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-1)
	t.Log(stack.stack)

	t.Log(stack.GetMin())
	t.Log(stack.Top())
	stack.Pop()
	t.Log(stack.GetMin())
}

func TestIsPopOrder(t *testing.T) {
	assert.Equal(t, IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}), true)
	assert.Equal(t, IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}), false)
}

package algorithm

import "testing"

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

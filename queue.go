package algorithm

/*
232. 用栈实现队列
https://leetcode-cn.com/problems/implement-queue-using-stacks/
思路：
1、最新的加入的元素需要放在栈底
2、一个栈用来装最终元素的顺序
3、push元素时，用一个临时栈存放预计存在的元素
*/
type MyQueue struct {
	stack     []int
	tampStack []int
}

func Constructor2() MyQueue {
	return MyQueue{
		stack:     []int{},
		tampStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	for i := len(this.stack) - 1; i >= 0; i-- {
		this.tampStack = append(this.tampStack, this.stack[i])
	}
	this.stack = []int{x}

	for i := len(this.tampStack) - 1; i >= 0; i-- {
		this.stack = append(this.stack, this.tampStack[i])
	}
	this.tampStack = []int{}
}

func (this *MyQueue) Pop() int {
	if len(this.stack) == 0 {
		return 0
	}

	result := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	return result
}

func (this *MyQueue) Peek() int {
	if len(this.stack) == 0 {
		return 0
	}

	return this.stack[len(this.stack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

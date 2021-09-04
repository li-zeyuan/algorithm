package algorithm

/*
MinStack
最小栈
https://leetcode-cn.com/problems/min-stack/
思路：
1、每次push一个元素时，记录当前值和当前最小值
2、pop时直接删除当前值和最小值
*/
type MinStack struct {
	stack []Item
}

type Item struct {
	minItem int
	val     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]Item, 0),
	}
}

func (this *MinStack) Push(val int) {
	item := Item{}
	item.val = val
	if len(this.stack) > 0 && val > this.stack[len(this.stack)-1].minItem {
		item.minItem = this.stack[len(this.stack)-1].minItem
	} else {
		item.minItem = val
	}

	this.stack = append(this.stack, item)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	result := 0
	if len(this.stack) > 0 {
		result = this.stack[len(this.stack)-1].val
	}

	return result
}

func (this *MinStack) GetMin() int {
	result := 0
	if len(this.stack) > 0 {
		result = this.stack[len(this.stack)-1].minItem
	}

	return result
}

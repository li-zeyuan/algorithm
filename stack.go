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

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pushV int整型一维数组
 * @param popV int整型一维数组
 * @return bool布尔型
 */
/*
思路：
1、定义一个辅助栈
2、p1指向popV队列
3、遍历pushV队列，加入辅助栈
4、若值不等于p1指向值，
5、若辅助栈 栈顶元素 == p1指向值，出栈
6、遍历结束，辅助栈是否为空
 */
func IsPopOrder( pushV []int ,  popV []int ) bool {
	if len(pushV) != len(popV) {
		return false
	}

	pop := 0
	push := 0
	stack := make([]int, 0)
	for  {
		if len(stack) == 0 {
			stack = append(stack, pushV[push])
			push ++
			continue
		}
		if stack[len(stack) - 1] == popV[pop] {
			stack = stack[:len(stack) - 1]
			pop ++
		}else if push < len(pushV) {
			stack = append(stack,  pushV[push])
			push ++
		}else {
			push ++
		}

		if len(stack) > 0  && push > len(pushV) {
			break
		}
		if len(stack) == 0 && pop == len(popV) {
			break
		}

	}

	return len(stack) == 0
}
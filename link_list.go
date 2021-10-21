package algorithm

import (
	"fmt"
)

/*
单向链表
h->1->2->nil
*/

type Node struct {
	data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
	}
}

type LList struct {
	header *Node
}

func NewList() *LList {
	return &LList{}
}

// 表头增加节点
func (l *LList) Add(data interface{}) {
	n := NewNode(data)
	if l.header == nil {
		l.header = n
		n.Next = new(Node)
	} else {
		n.Next = l.header
		l.header = n
		n.Next = new(Node)
	}
}

func (l *LList) Append(data interface{}) {
	n := NewNode(data)
	if l.header == nil {
		l.header = n
		n.Next = new(Node)
	} else {
		curNode := l.header
		for curNode.Next != nil {
			curNode = curNode.Next
		}

		curNode.data = data
		curNode.Next = new(Node)
	}
}

func (l *LList) Insert(i int, data interface{}) {
	if i <= 0 {
		l.Add(data)
		return
	}
	if i >= l.Length() {
		l.Append(data)
		return
	}

	curNode := l.header
	curIdx := 0
	for curNode != nil {
		if curIdx == i {
			n := NewNode(data)
			n.Next = curNode.Next
			curNode.Next = n
			break
		}
		curIdx++
	}
}

// 翻转链表：遍历链表，把当前的node.next指向前一个节点
func (l *LList) Revert() {
	preNode := new(Node)

	curNode := l.header
	lastNode := new(Node)
	for curNode != nil {
		lastNode = curNode

		next := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = next
	}

	l.header = lastNode
}

func (l *LList) Delete(i int) {
	if i+1 > l.Length() {
		return
	}

	if i == 0 {
		l.header = l.header.Next
		return
	}

	curNode := l.header
	lastNode := l.header
	for curNode.Next != nil {
		if i == 0 {
			lastNode.Next = curNode.Next
			break
		}

		i--
		lastNode = curNode
		curNode = curNode.Next
	}
}

func (l *LList) Length() int {
	curNode := l.header

	length := 0
	for curNode.Next != nil {
		length++
		curNode = curNode.Next
	}

	return length
}

func (l *LList) Scan() {
	curNode := l.header

	for curNode != nil {
		fmt.Print(curNode.data, " ")
		curNode = curNode.Next
	}
	fmt.Println()
}

/*
参考
https://learnku.com/articles/44998
*/

/*
相交链表
https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
方式一：遍历
1、每取出L1一个元素，遍历L2中的所有元素
2、判断两个元素是否相等
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	itemA := headA
	for itemA != nil {
		itemB := headB
		for itemB != nil {
			if itemA == itemB {
				return itemA
			}
			itemB = itemB.Next
		}

		itemA = itemA.Next
	}

	return nil
}

/*
相交链表
https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
方式二：双指针
1、pA、pB分别从 lA、lB的头节点出发，每次移动一个位置
2、判断两个指针的元素是否相等，相等则返回该节点
3、pA移动到尾部时，重新指向lB头节点
4、pB移动到尾部时，重新指向lA头节点
*/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

/*
234. 回文链表
https://leetcode-cn.com/problems/palindrome-linked-list/
思路：
1、翻转链表
2、对比翻转后和当前链表
 */
func isPalindrome3(head *ListNode) bool {
	if head == nil {
		return false
	}

	var revList *ListNode
	curNode := head
	for curNode != nil {
		tempNode := &ListNode{
			Val: curNode.Val,
		}
		tempNode.Next = revList
		revList = tempNode

		curNode = curNode.Next
	}

	curNode2 := head
	revNode := revList
	for curNode2.Next != nil {
		if curNode2.Val != revNode.Val {
			return false
		}

		curNode2 = curNode2.Next
		revNode = revNode.Next
	}

	return true
}
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
	result := new(Node)
	for curNode != nil {
		result = curNode
		next := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = next
	}

	l.header = result
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
203. 移除链表元素
https://leetcode-cn.com/problems/remove-linked-list-elements/
思路：
1、快慢指针
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head // 初始都指向表头
	for fast != nil {
		if slow.Val == val { // 说明表头处为要移除的元素
			head = head.Next
			slow = head
			fast = head
			continue
		}

		if fast.Val == val {
			slow.Next = fast.Next
			fast = fast.Next
			continue
		}

		slow = fast
		fast = fast.Next
	}

	return head
}

/*
反转链表
https://leetcode-cn.com/problems/reverse-linked-list/
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var result = &ListNode{head.Val, nil}

	curNode := head.Next
	for curNode != nil {
		result = &ListNode{
			Val:  curNode.Val,
			Next: result,
		}
		curNode = curNode.Next
	}

	return result
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

/*
2. 两数相加
思路：
直接相加，满十进1
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := &ListNode{}
	sumPrt := sum
	hL1 := l1
	hL2 := l2
	isUp := 0
	for hL1 != nil || hL2 != nil {
		if hL1 != nil && hL2 != nil {
			s := hL1.Val + hL2.Val + isUp
			isUp = 0
			if s >= 10 {
				isUp = 1
			}
			sumPrt.Next = &ListNode{
				Val: (s) % 10,
			}
			sumPrt = sumPrt.Next
			hL1 = hL1.Next
			hL2 = hL2.Next
		} else if hL1 != nil && hL2 == nil {
			s2 := hL1.Val + isUp
			isUp = 0
			if s2 >= 10 {
				isUp = 1
			}

			sumPrt.Next = &ListNode{
				Val: (s2) % 10,
			}
			sumPrt = sumPrt.Next

			hL1 = hL1.Next
		} else if hL1 == nil && hL2 != nil {
			s3 := hL2.Val + isUp
			isUp = 0
			if s3 >= 10 {
				isUp = 1
			}
			sumPrt.Next = &ListNode{
				Val: (s3) % 10,
			}
			sumPrt = sumPrt.Next

			hL2 = hL2.Next
		}
	}

	if isUp == 1 {
		sumPrt.Next = &ListNode{
			Val: 1,
		}
	}

	return sum.Next
}

/*
19、删除链表第n个节点
思路1：反转链表
1、反转
2、删除节点
3、反转
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 1 {
		return nil
	}
	head = rev(head)
	show := head
	fast := head
	if n == 1 {
		head = head.Next
		return rev(head)
	}

	for fast != nil {
		if n == 1 {
			show.Next = fast.Next
			break
		}
		n--
		show = fast
		fast = fast.Next
	}

	return rev(head)
}

func rev(head *ListNode) *ListNode {
	var newList *ListNode
	cur := head
	for cur != nil {
		n := &ListNode{
			Val: cur.Val,
		}
		n.Next = newList
		newList = n
		cur = cur.Next
	}

	return newList
}

/*
19、删除链表第n个节点
思路2：队列
1、遍历链表，放入队列
2、删除队列节点
3、拼接成链表
*/

/*
19、删除链表第n个节点
思路：计数
1、遍历链表，记录链表长度
2、删除第l-n个接点
*/

/*
19、删除链表第n个节点
思路3：反转->删除节点->反转
*/

func gen(n int, strList []string, str string) {
	if n == 0 {
		return
	}

	nStrList := make([]string, 0)
	for _, s := range strList {
		nStrList = append(nStrList, s+str)
	}
	strList = nStrList
	n--

	gen(n, strList, "(")
	gen(n, strList, ")")
}

/*
间隔划分链表
*/
func splitLink(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	var firstLink *ListNode
	var secondLink *ListNode
	fPrt := firstLink
	sPrt := secondLink

	prt := head
	i := 0
	for prt != nil {
		tempNode := &ListNode{
			Val: prt.Val,
		}
		if i%2 == 0 {
			if firstLink == nil {
				firstLink = tempNode
				fPrt = firstLink
			} else {
				fPrt.Next = tempNode
				fPrt = fPrt.Next
			}
		} else {
			if secondLink == nil {
				secondLink = tempNode
				sPrt = secondLink
			} else {
				sPrt.Next = tempNode
				fPrt = sPrt.Next
			}
		}

		i++
		prt = prt.Next
	}

	return firstLink, secondLink
}

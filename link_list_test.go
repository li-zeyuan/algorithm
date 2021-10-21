package algorithm

import (
	"fmt"
	"testing"
)

func TestLList_Add(t *testing.T) {
	lList := NewList()
	lList.Add(1)
	lList.Add(2)

	lList.Scan()
}

func TestLList_Append(t *testing.T) {
	lList := NewList()

	lList.Append(1)
	lList.Append(2)
	lList.Scan()
}

func TestLList_Length(t *testing.T) {
	lList := NewList()

	lList.Append(1)
	lList.Append(2)
	fmt.Println(lList.Length())
}

func TestLList_Insert(t *testing.T) {
	lList := NewList()

	lList.Append(1)
	lList.Append(2)
	lList.Scan()

	lList.Insert(6, 0)
	lList.Scan()
}

func TestLList_Revert(t *testing.T) {
	lList := NewList()

	lList.Append(1)
	lList.Append(2)
	lList.Scan()

	lList.Revert()
	lList.Scan()
}
func TestLList_Delete(t *testing.T) {
	lList := NewList()

	lList.Append(1)
	lList.Append(2)
	lList.Append(3)
	lList.Delete(0)

	lList.Scan()
}

func TestLList_getIntersectionNode(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	node1.Next = node3
	lA := node1
	node2.Next = node3
	lB := node2
	t.Log(getIntersectionNode(lA, lB))
}

func TestLList_getIntersectionNode2(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	node1.Next = node3
	lA := node1
	node2.Next = node3
	lB := node2
	t.Log(getIntersectionNode2(lA, lB))
}

func TestLList_isPalindrome3(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 2}
	node4 := &ListNode{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	t.Log(isPalindrome3(node1))
}

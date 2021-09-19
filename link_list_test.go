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

func TestRemoveElements(t *testing.T) {
	head := &ListNode{1, &ListNode{
		2,
		&ListNode{
			6,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						&ListNode{
							6,
							nil,
						},
					},
				},
			},
		},
	}}

	//head := &ListNode{1, &ListNode{1, nil}}
	result := removeElements(head, 6)
	t.Log(result)
}

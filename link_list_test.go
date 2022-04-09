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

func TestLList_generateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(0))
}

func TestLList_gen(t *testing.T) {
	gen(0, []string{""}, "")
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

func TestLList_reverseList(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	node1.Next = node2
	node2.Next = node3

	reverseResult := reverseList(node1)
	t.Log(reverseResult)
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

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	l2 := &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}

	//head := &ListNode{1, &ListNode{1, nil}}
	result := addTwoNumbers(l1, l2)
	t.Log(result)
}

func TestRemoveNthFromEnd(t *testing.T) {
	l1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}

	result := removeNthFromEnd(l1, 3)
	t.Log(result)
}

func TestSplitLink(t *testing.T) {
	l1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}

	first, second := splitLink(l1)
	t.Log(first)
	t.Log(second)
}

func TestMvLink(t *testing.T) {
	l1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}

	result := MvLink(l1, 1)
	t.Log(result)

	l2 := &ListNode{
		1,
		nil,
	}

	result2 := MvLink(l2, 1)
	t.Log(result2)
}

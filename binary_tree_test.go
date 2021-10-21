package algorithm

import (
	"fmt"
	"testing"
)

func TestBinaryTree_Search(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(4)

	fmt.Println(tree.Search(4))
}

func TestBinaryTree_MiddleShow(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(8)

	tree.MiddleShow()
}

func TestNewBinaryTree(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(8)

	tree.FirstShow()
}

func TestBinaryTree_LevelShow(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)

	tree.LevelShow()
}

func TestBinaryTree_FirstShow(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(8)

	tree.Reverse()
	tree.MiddleShow()
}

func TestInvertTree(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			2,
			&TreeNode{
				1,
				nil,
				nil,
			},
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
		Right: &TreeNode{
			7,
			&TreeNode{
				6,
				nil,
				nil,
			},
			&TreeNode{
				9,
				nil,
				nil,
			},
		}}
	tree := invertTree(root)
	t.Log(tree)
}

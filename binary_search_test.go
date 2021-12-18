package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearchObj_BinarySearch(t *testing.T) {
	binaryS := NewBSO([]int{1, 2, 3}, 2)
	fmt.Println(binaryS.BinarySearch())
}

func Test1(t *testing.T) {
	A()
}

func A() {
	p := Person{}
	p.howOld(2)
	p.howOld2("lizeyuan")
}

type Person struct {
	name string
	age  int
}

func (p *Person) howOld(i int) {
	p.age = i
}

func (p *Person) howOld2(s string) {
	p.name = s
}

func TestTwoSum2(t *testing.T) {
	t.Log(twoSum2([]int{2, 7, 11, 15}, 9))
}

func TestTwoSum3(t *testing.T) {
	t.Log(twoSum3([]int{2, 7, 11, 15}, 9))
}

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	p := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	q := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	node := lowestCommonAncestor(root, p, q)
	t.Log(node)
}

func TestBinaryTreePaths(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		//Left:  &TreeNode{
		//	Val:   2,
		//	Left:  nil,
		//	Right: &TreeNode{
		//		Val:   5,
		//		Left:  nil,
		//		Right: nil,
		//	},
		//},
		//Right: &TreeNode{
		//	Val:   3,
		//	Left:  nil,
		//	Right: nil,
		//},
	}

	t.Log(binaryTreePaths(root))
}

func TestAddDigits(t *testing.T) {
	t.Log(addDigits(38))
}

func TestBinaryTreePaths2(t *testing.T) {
	t.Log(binaryTreePaths2(nil))
}

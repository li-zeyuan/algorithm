package algorithm

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		1,
		&TreeNode{
			Val: 2,
		},
		&TreeNode{
			Val: 3,
		},
	}
	fmt.Println(inorderTraversal(root))
}

func TestIsSameTree(t *testing.T) {
	q := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			1,
			nil,
			nil,
		},
	}
	p := &TreeNode{
		1,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			2,
			nil,
			nil,
		},
	}

	fmt.Println(isSameTree(q, p))
}

func TestIsSymmetric(t *testing.T) {
	q := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
	}

	fmt.Println(isSymmetric(q))
}

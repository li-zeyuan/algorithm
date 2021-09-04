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

func TestMaxDepth(t *testing.T) {
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

	fmt.Println(maxDepth(q))
}

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	node := sortedArrayToBST(nums)
	t.Log(node)
}

func TestIsBalanced(t *testing.T) {
	q := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				&TreeNode{
					4,
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
				3,
				nil,
				nil,
			},
		},
		&TreeNode{
			2,
			nil,
			nil,
		},
	}

	t.Log(isBalanced(q))
}

func TestMinDepth(t *testing.T) {
	q := &TreeNode{
		2,
		nil,
		&TreeNode{
			3,
			nil,
			&TreeNode{
				4,
				nil,
				&TreeNode{
					5,
					nil,
					&TreeNode{
						6,
						nil,
						nil,
					},
				},
			},
		},
	}

	t.Log(minDepth(q))
}

//[5,4,8,11,null,13,4,7,2,null,null,null,1],
func TestHasPathSum(t *testing.T) {
	r := &TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{
				11,
				&TreeNode{
					7,
					nil,
					nil,
				},
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
			nil,
		},
		&TreeNode{
			8,
			&TreeNode{
				13,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
	}

	t.Log(hasPathSum(r, 22))
}

func TestPreorderTraversal(t *testing.T) {
	q := &TreeNode{
		2,
		nil,
		&TreeNode{
			3,
			nil,
			&TreeNode{
				4,
				nil,
				&TreeNode{
					5,
					nil,
					&TreeNode{
						6,
						nil,
						nil,
					},
				},
			},
		},
	}

	t.Log(preorderTraversal(q))
}
func TestPostorderTraversal(t *testing.T) {
	q := &TreeNode{
		2,
		nil,
		&TreeNode{
			3,
			nil,
			&TreeNode{
				4,
				nil,
				&TreeNode{
					5,
					nil,
					&TreeNode{
						6,
						nil,
						nil,
					},
				},
			},
		},
	}

	t.Log(postorderTraversal(q))
}

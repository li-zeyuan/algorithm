package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestSumOfLeftLeaves(t *testing.T) {
	q := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}

	assert.Equal(t, sumOfLeftLeaves(q), 24)

	q2 := &TreeNode{
		1,
		nil,
		nil,
	}

	assert.Equal(t, sumOfLeftLeaves(q2), 0)
}

func TestSumOfLeftLeaves2(t *testing.T) {
	q := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}

	assert.Equal(t, sumOfLeftLeaves2(q), 24)

	q2 := &TreeNode{
		1,
		nil,
		nil,
	}

	assert.Equal(t, sumOfLeftLeaves2(q2), 0)
}

func TestGetMinimumDifference(t *testing.T) {
	q := &TreeNode{
		4,
		&TreeNode{
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
		&TreeNode{
			6,
			nil,
			nil,
		},
	}

	assert.Equal(t, getMinimumDifference(q), 1)

	q2 := &TreeNode{
		1,
		&TreeNode{
			0,
			nil,
			nil,
		},
		&TreeNode{
			48,
			&TreeNode{
				12,
				nil,
				nil,
			},
			&TreeNode{
				49,
				nil,
				nil,
			},
		},
	}

	assert.Equal(t, getMinimumDifference(q2), 1)

	q3 := &TreeNode{
		236,
		&TreeNode{
			104,
			nil,
			&TreeNode{
				227,
				nil,
				nil,
			},
		},
		&TreeNode{
			701,
			nil,
			&TreeNode{
				911,
				nil,
				nil,
			},
		},
	}

	assert.Equal(t, getMinimumDifference(q3), 9)
}

func TestDiameterOfBinaryTree(t *testing.T) {
	q := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	assert.Equal(t, diameterOfBinaryTree(q), 3)

	q2 := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		nil,
	}

	assert.Equal(t, diameterOfBinaryTree(q2), 1)
}

func TestMostSumOfBinaryTree(t *testing.T) {

	q2 := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	assert.Equal(t, mostSumOfBinaryTree(q2), 6)

	q3 := &TreeNode{
		-10,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}

	assert.Equal(t, mostSumOfBinaryTree(q3), 42)
}

func TestLowestCommonAncestor2(t *testing.T) {
	p := &TreeNode{
		Val: 4,
	}
	q := &TreeNode{
		Val: 5,
	}

	root := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	assert.Equal(t, lowestCommonAncestor2(root, p, q).Val, 2)
}

package algorithm

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T)  {
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

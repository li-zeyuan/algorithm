package algorithm

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}

/*
相同的树
https://leetcode-cn.com/problems/same-tree/
思路：
１．中序遍历树
２．对比树节点的值
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil{
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}

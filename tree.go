package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
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
	if p == nil && q == nil {
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

/*
对称二叉树
https://leetcode-cn.com/problems/symmetric-tree/
对称二叉树的条件：
1、左右两颗子树互为镜像

递归思路：
1、对比左右两子树的值
2、定义两指针，一个指针向左移，一个指针向右移
3、对比两个指针的只是否相等

迭代思路：
1、定义一个队列
2、初始放入两个root节点
3、每次从队头取出两个节点最对比
4、两个结点的左右子结点按相反的顺序插入队列中
*/
func isSymmetric(root *TreeNode) bool {
	return check(root, root) // 初始化两指针都指向根
}
func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val &&
		check(l.Left, r.Right) && // l指针向左移，r指针向右移
		check(l.Right, r.Left) // // l指针向右移，r指针向左移
}

/*
二叉树的最大深度
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
思路：
1、迭代终止条件：当前节点为nil
2、当前节点不为空，则返回1 + max(左节点深度，右节点深度）
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return 1+leftDepth
	}

	return 1 + rightDepth
}
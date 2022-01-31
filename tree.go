package algorithm

import (
	"math"
)

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
		return 1 + leftDepth
	}

	return 1 + rightDepth
}

/*
将有序数组转化成 高度平衡 二叉树搜索树
https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/submissions/
高度平衡二叉树（平衡二叉树）：每个节点子树高度<=1
思路：
1、二叉树的中序遍历就是有序数组
2、截取数组中间位（或偏左）为根节点
3、分别多左右两个子数组 递归操作（2）
4、拼接成一颗树
*/
func sortedArrayToBST(nums []int) *TreeNode {
	mid := (len(nums) - 1) / 2
	node := &TreeNode{Val: nums[mid]}
	if mid == 0 && len(nums) == 2 { // 当有两元素，需要继续把递归右侧，返回左元素
		node.Right = sortedArrayToBST(nums[mid+1:])
		return node
	}
	if mid == 0 && len(nums) < 2 { // 当只有一个元素，直接返回即可
		return node
	}

	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}

/*
判断是否平衡二叉树
https://leetcode-cn.com/problems/balanced-binary-tree/
思路：
1、问题转化成判断每个节点的左右两颗子树的最大高度差是否超过1
2、从跟节点开始，递归判断左右两个节点
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 求节点的最大深度
func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(depth(node.Left), depth(node.Right)) + 1
}

/*
二叉树最小深度
https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
思路：

*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	minLeft, minRight := 0, 0
	if root.Left != nil {
		minLeft = minDepth(root.Left)
	}

	if root.Right != nil {
		minRight = minDepth(root.Right)
	}

	return min(minLeft, minRight) + 1
}

func min(x, y int) int { // 最小非零值
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}

	if x > y {
		return y
	}
	return x
}

/*
路径总和
https://leetcode-cn.com/problems/path-sum/
思路：
1、边界条件：root == nil && targetSum > 0 返回 false
2、边界条件：root != nil && targetSum == 0 返回 false
3、递归查找树的左边
4、递归查找树的右边
5、只要有一边是返回true，则结果就是true
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val-targetSum == 0 {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

/*
二叉树的前序遍历
https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
*/
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

/*
二叉树的后序遍历
https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
*/
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

/*
404. 左叶子之和
https://leetcode-cn.com/problems/sum-of-left-leaves/
思路：深度优先
1、isLeft标识是否未左节点
*/
func sumOfLeftLeaves(root *TreeNode) int {
	return nodeOfLeftLeave(root.Left, true) + nodeOfLeftLeave(root.Right, false)
}

func nodeOfLeftLeave(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil && isLeft {
		return node.Val
	}

	return nodeOfLeftLeave(node.Left, true) + nodeOfLeftLeave(node.Right, false)
}

/*
404. 左叶子之和
https://leetcode-cn.com/problems/sum-of-left-leaves/
思路：广度优先
*/
func sumOfLeftLeaves2(root *TreeNode) int {
	result := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		} else if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return result
}

/*
530. 二叉搜索树的最小绝对差
https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
思路：
1、二叉树中序便利为有序数组
2、有序数组的最小距离为相邻两个元素的差值
3、pre为前一个节点的值，用当前值-pre 和 结果做比较
*/
func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

/*
543. 二叉树的直径
https://leetcode-cn.com/problems/diameter-of-binary-tree/
思路：
1、不一定经过根结点，定义一个变量保存最大路径
2、最大左边节点路径 + 最大右边节点路径 = 当前节点最大值
3、递归所有的节点，比1中的变量大则更新值
*/
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0 // 保存最大路径和、

	var dfs func(node *TreeNode) int // 求该节点的最大深度
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := dfs(node.Left)
		rh := dfs(node.Right)
		result = max4(lh+rh, result)

		return max4(lh, rh) + 1
	}
	dfs(root)

	return result
}

func max4(x, y int) int {
	if x > y {
		return x
	}
	return y
}

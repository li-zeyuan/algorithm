package algorithm

import (
	"strconv"
)

/*
二分法通过对折的方式查找一个数据，条件是必须是一个有序的数组。数组底层是顺序链表，可以随机获取一个值 O (log n)
参考
- https://learnku.com/articles/45836
*/

type BinarySearchObj struct {
	l      []int
	target int
}

func NewBSO(l []int, target int) *BinarySearchObj {
	return &BinarySearchObj{
		l:      l,
		target: target,
	}
}

func (b *BinarySearchObj) BinarySearch() int {
	l, r := 0, len(b.l)-1
	for l <= r {
		mid := (l + r) / 2
		if b.l[mid] == b.target {
			return 0
		}
		if b.target > b.l[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

/*
两数之和 II - 输入有序数组
https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
方式一：二分查找
1、依次遍历数组，得第一个元素
2、第二个元素为target - 第一个元素，二分查找第二个元素
*/
func twoSum2(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	for i := 0; i < len(numbers); i++ {
		start, end := i+1, len(numbers)-1
		for start <= end {
			mid := (start + end) / 2
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return nil
}

/*
方式二：双指针
1、一个指针指向第一个元素，一个指针指向最后一个元素
2、若两指针之和大与target，则右指针左移
3、若两指针之和小与target，则左指针右移
*/
func twoSum3(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}
	start, end := 0, len(numbers)-1
	for start < end {
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		} else if numbers[start]+numbers[end] > target {
			end--
		} else {
			start++
		}
	}

	return nil
}

/*
235. 二叉搜索树的最近公共祖先
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}
	if root.Val <= p.Val && root.Val >= q.Val {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return nil
}

/*
257. 二叉树的所有路径
https://leetcode-cn.com/problems/binary-tree-paths/
方式一：深度优先（递归）
*/
var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = make([]string, 0)
	deepenSearch(root, "")
	return paths
}

func deepenSearch(root *TreeNode, prePath string) {
	if root != nil {
		prePath += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil { // 叶子节点
			paths = append(paths, prePath)
		} else { // 非叶子节点则继续递归
			prePath += "->"
			deepenSearch(root.Left, prePath)
			deepenSearch(root.Right, prePath)

		}
	}
}

/*
257. 二叉树的所有路径
https://leetcode-cn.com/problems/binary-tree-paths/
方式二：广度优先（队列存储节点）
*/
func binaryTreePaths2(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
	var nodeQueue []*TreeNode
	var pathQueue []string
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))

	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}

/*
258. 各位相加
https://leetcode-cn.com/problems/add-digits/
*/
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	sum := 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}

	return addDigits(sum)
}

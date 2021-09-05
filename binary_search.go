package algorithm

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

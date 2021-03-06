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

/*
33. 搜索旋转排序数组
https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
思路：二分搜索
1、二分后，总有一半是有序的
*/
func revolveArraySearch(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		// 中间值即为target,直接返回mid
		if target == nums[mid] {
			return mid
		}
		// 此时前半部分有序
		if nums[0] <= nums[mid] {
			// 此时target落在前半部分有序区间内
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				// 此时target落在后半部分无序区间内
				l = mid + 1
			}
		} else {
			// 此时后半部分有序
			// 此时target落在后半部分有序区间内
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				// 此时target落在前半部分无序区间内
				r = mid - 1
			}
		}
	}
	// 循环结束没有找到目标值target，返回-1
	return -1
}

/*
34. 在排序数组中查找元素的第一个和最后一个位置
https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
思路：二分查找
1、找出target
2、双指针找出起初和终点
*/
func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	}

	left := 0
	right := l - 1
	for left <= right {
		m := (left + right) / 2
		if nums[m] == target {
			return sRange(nums, m)
		} else if target < nums[m] {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return []int{-1, -1}
}

func sRange(nums []int, index int) []int {
	start, end := index, index
	sRange := make([]int, 2)
	sRange[0] = start
	sRange[1] = end
	for {
		isBreak := true
		if start >= 0 && nums[start] == nums[index] {
			sRange[0] = start
			start--
			isBreak = false
		}
		if end < len(nums) && nums[end] == nums[index] {
			sRange[1] = end
			end++
			isBreak = false
		}

		if isBreak {
			break
		}
	}

	return sRange
}

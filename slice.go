package algorithm

import (
	"math"
	"sort"
)

/*
反转字符列表
使得{"A", "B", "C", "D"} -> {"D", "C", "B", "A"}

思路
	1、以切片中心为轴，翻转切片
*/
func ReversedSlice(strList []string) []string {
	if len(strList) == 0 {
		return []string{}
	}

	for i := 0; i < len(strList)/2; i++ {
		strList[i], strList[len(strList)-1-i] = strList[len(strList)-1-i], strList[i]
	}

	return strList
}

/*
344. 反转字符串
https://leetcode-cn.com/problems/reverse-string/
*/
func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}

	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

/*
349. 两个数组的交集
https://leetcode-cn.com/problems/intersection-of-two-arrays/
思路：map
*/
func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] = struct{}{}
	}

	nums2Map := make(map[int]struct{})
	for i := 0; i < len(nums2); i++ {
		nums2Map[nums2[i]] = struct{}{}
	}

	result := make([]int, 0)
	for k := range nums2Map {
		if _, ok := nums1Map[k]; ok {
			result = append(result, k)
		}
	}

	return result
}

/*
350. 两个数组的交集 II
https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
*/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	nums1Map := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] += 1
	}

	nums2Map := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		nums2Map[nums2[i]] += 1
	}

	result := make([]int, 0)
	for k, v := range nums1Map {
		amount, ok := nums2Map[k]
		if !ok {
			continue
		}

		maxAmount := int(math.Min(float64(v), float64(amount)))
		for maxAmount > 0 {
			result = append(result, k)
			maxAmount--
		}
	}

	return result
}

/*
414. 第三大的数
https://leetcode-cn.com/problems/third-maximum-number/
*/
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	if len(nums) < 3 {
		return nums[0]
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			j++
		}
		if j == 2 {
			return nums[i]
		}
	}

	return nums[0]
}

/*
448. 找到所有数组中消失的数字
https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
思路：
1、构造map
2、遍历nums
*/
func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	numsMap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = struct{}{}
	}

	result := make([]int, 0)
	for j := 1; j <= len(nums); j++ {
		if _, ok := numsMap[j]; !ok {
			result = append(result, j)
		}
	}

	return result
}

/*
思路：
1、第一次遍历，把num放回顺序的位置，并+n
2、第二次遍历，没有大于n的num为缺失num
*/
func findDisappearedNumbers2(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}

/*
453. 最小操作次数使数组元素相等
https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/
思路：
1、“每次操作将会使 n - 1 个元素增加 1” 等同于 “每次将1个元素减1”
2、最少的操作即将所有的元素减到数组的最小值
*/
func minMoves(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}

	result := 0
	for _, n2 := range nums {
		result = result + (n2 - min)
	}

	return result
}

/*
455. 分发饼干
https://leetcode-cn.com/problems/assign-cookies/
思路：
1、排序两数组
2、双指针遍历，若s[j] >= g[i];则result ++
*/
func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}

	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	i, j, result := 0, 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			result++
			i++
			j++
		} else {
			j++
		}
	}

	return result
}

/*
11. 盛最多水的容器
https://leetcode-cn.com/problems/container-with-most-water/
思路：暴力（超时）
1、指针1：以该位置为起点
2、指针2：以该位置为终点
 */
func maxArea(height []int) int {
	l := len(height)
	if l <= 1 {
		return 0
	}

	result := 0
	for start := 0; start < l ; start ++ {
		for end := start + 1; end < l ; end ++ {
			sum := 0
			if height[start] > height[end] {
				sum = height[end] * (end - start)
			}else {
				sum = height[start] * (end - start)
			}

			if sum > result {
				result = sum
			}
		}
	}

	return result
}
/*
11. 盛最多水的容器
https://leetcode-cn.com/problems/container-with-most-water/
思路：双指针（没懂）
1、指针1：起初指向0
2、指针2：起初指向len()
3、哪个指针小，则移动哪个指针
*/
func maxArea2(height []int) int {
	l := len(height)
	if l <= 1 {
		return 0
	}

	p1 := 0
	p2 := l - 1
	result := 0
	for p2 > p1 {
		ans := 0
		if height[p1] > height[p2] {
			ans = (p2 - p1) * height[p2]
			p2 --
		}else {
			ans = (p2 - p1) * height[p1]
			p1 ++
		}

		if ans > result {
			result = ans
		}
	}

	return result
}
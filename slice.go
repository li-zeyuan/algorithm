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

	if len(nums) <3 {
		return nums[0]
	}
	j := 0
	for i := 1 ; i < len(nums); i ++ {
		if nums[i] < nums[i -1] {
			j ++
		}
		if j == 2 {
			return nums[i]
		}
	}

	return nums[0]
}

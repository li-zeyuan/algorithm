package algorithm

import (
	"fmt"
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
	for start := 0; start < l; start++ {
		for end := start + 1; end < l; end++ {
			sum := 0
			if height[start] > height[end] {
				sum = height[end] * (end - start)
			} else {
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
			p2--
		} else {
			ans = (p2 - p1) * height[p1]
			p1++
		}

		if ans > result {
			result = ans
		}
	}

	return result
}

/*
17. 电话号码的字母组合
思路：暴力
1、遍历输入字符串，取出一个字符，把对应的字母输入加入结果
2、遍历结果，并入当前按键的字母
*/
var phoneMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return []string{}
	}

	result := make([]string, 0)
	for i := 0; i < l; i++ {
		tempResult := make([]string, 0)
		aList, ok := phoneMap[string(digits[i])]
		if !ok {
			continue
		}
		for _, item := range result {
			for _, phone := range aList {
				tempResult = append(tempResult, item+phone)
			}
		}
		if len(result) == 0 {
			tempResult = append(tempResult, aList...)
		}

		result = tempResult
	}

	return result
}

/*
48.旋转图像
思路：归纳法
旋转次数：n * n / 4
旋转元素：[x][y] --> [y][n - 1 - x]
*/
func rotate(matrix [][]int) {
	n := len(matrix)

	if n <= 1 {
		return
	}

	//旋转次数
	count := n * n / 4

	x := 0
	y := 0
	z := 0

	var x1, y1, x2, y2 int

	for i := 0; i < count; i++ {
		if z >= (n-1)-2*x {
			x += 1
			z = 0
		}
		y = z + x
		z += 1

		x1 = x
		y1 = y

		for j := 0; j < 3; j++ {
			x2 = n - 1 - y1
			y2 = x1

			matrix[x1][y1] = matrix[x1][y1] ^ matrix[x2][y2]
			matrix[x2][y2] = matrix[x1][y1] ^ matrix[x2][y2]
			matrix[x1][y1] = matrix[x1][y1] ^ matrix[x2][y2]

			x1 = x2
			y1 = y2
		}
	}
}

/*
49、字母异位词分组
思路：排序
1、排序strs元素，构造map
2、遍历map,加入结果集
*/
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	sortMap := make(map[string][]string)
	for _, str := range strs {
		bSlice := []byte(str)
		sort.Slice(bSlice, func(i, j int) bool {
			return bSlice[i] < bSlice[j]
		})

		sortMap[string(bSlice)] = append(sortMap[string(bSlice)], str)
	}

	result := make([][]string, 0)
	for _, v := range sortMap {
		result = append(result, v)
	}

	return result
}

/*
54.螺旋矩阵
思路：模拟轨迹m*n
1、设定左右上下边界
2、分别左移、下移、右移、上移，碰壁则终止并移动边界
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	x, y := 0, 0
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1 // 设定左右上下边界
	result := make([]int, 0)
	for avoid(x, y, left, right, up, down) {
		for y = left; avoid(x, y, left, right, up, down); y++ { // 左
			result = append(result, matrix[x][y])
		}

		up++                                                  // 移动边界
		y--                                                   // 归位
		for x = up; avoid(x, y, left, right, up, down); x++ { // 下
			result = append(result, matrix[x][y])
		}

		right--
		x--
		for y = right; avoid(x, y, left, right, up, down); y-- { // 右
			result = append(result, matrix[x][y])
		}

		down--
		y++
		for x = down; avoid(x, y, left, right, up, down); x-- { // 上
			result = append(result, matrix[x][y])
		}
		left++
		x++
		y = left
	}
	return result
}

func avoid(x, y, left, right, up, down int) bool {
	return x >= up && x <= down && y >= left && y <= right
}

/*
55. 跳跃游戏
https://leetcode-cn.com/problems/jump-game/
思路：贪心
1、遍历列表，维护一个最大可到达 下标
2、判断最大可到达下标是否 >= 数组最后一个元素
*/
func canJump(nums []int) bool {
	maxIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 && maxIndex <= i {
			break
		}
		maxIndex = max5(maxIndex, i+nums[i])
	}

	return maxIndex >= len(nums)-1
}

func max5(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
56. 合并区间
https://leetcode-cn.com/problems/merge-intervals/
思路：
1、按照区间start排序
2、遍历intervals，判断当前区间和前一个区间是否有重合，重合则合并
*/
func merge3(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(intervals[i]) != 2 {
			continue
		}

		if len(result) == 0 {
			result = append(result, intervals[i])
			continue
		}

		pre := result[len(result)-1]
		if intervals[i][0] <= pre[1] { // 合并
			pre[1] = max5(intervals[i][1], pre[1])
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

/*
73、矩阵置零
思路：遍历
1、外层遍历找到为0的元素
2、内层遍历置0所在的行和列
*/
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	row := len(matrix)
	col := len(matrix[0])
	zeroItems := make([][]int, 0)
	for rowIdx := 0; rowIdx < row; rowIdx++ {
		for colIdx := 0; colIdx < col; colIdx++ {
			if matrix[rowIdx][colIdx] != 0 {
				continue
			}

			items := []int{rowIdx, colIdx}
			zeroItems = append(zeroItems, items)
		}
	}

	for i := 0; i < len(zeroItems); i++ {
		for rowI := 0; rowI < row; rowI++ { // 列置0
			matrix[rowI][zeroItems[i][1]] = 0
		}
		for colI := 0; colI < col; colI++ { // 行置0
			matrix[zeroItems[i][0]][colI] = 0
		}
	}

	fmt.Println(matrix)
}

/*
75、颜色分类
思路：
1、统计各个颜色出现的次数
2、重写数组
*/
func swapColors(colors []int) {
	count0 := 0
	count1 := 0
	for _, c := range colors {
		if c == 0 {
			count0++
		}
		if c == 1 {
			count1++
		}
	}

	for i := 0; i < len(colors); i++ {
		if i < count0 {
			colors[i] = 0
		} else if i-count0 < count1 {
			colors[i] = 1
		} else {
			colors[i] = 2
		}
	}

	fmt.Println(colors)
}

/*
75、颜色分类
思路：三指针
1、P0-1指向最后一个0的位置，p2+1指向最后一个2的位置，i指向但前位置
*/
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	p0, p2 := 0, len(nums)-1
	for i := 0; i < p2; i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		} else if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
			i--
		}
	}

	fmt.Println(nums)
}

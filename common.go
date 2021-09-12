package algorithm

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 单例
var (
	single *Singleton
	sOne   sync.Once
)

type Singleton struct{}

func GetSingleton() *Singleton {
	sOne.Do(func() {
		single = new(Singleton)
	})

	return single
}

// ===========================

// 装饰器
func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func coolFunc(a func()) { // go 通过函数传参实现装饰器
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

// 闭包

// ==================================

// iota :一个const中，iota初始值为0（第一行为0），可跳过，可占位
const (
	a = iota
	b = iota
)
const (
	name = "menglu" // 占位
	c    = iota
	d    = iota
)

func GetIota() {
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 1
	fmt.Println(d) // 2
}

/*
求阶乘和
1+2！+3！+4！
*/
// 求一个数阶乘
func Fac(i int) int {
	if i == 1 {
		return i
	}

	return i * Fac(i-1)
}

func Sum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += Fac(i)
	}

	return sum
}

/*
斐波纳契数列，又称黄金分割数列，指的是这样一个数列：1、1、2、3、5、8、13、21、……
在数学上，斐波纳契数列以如下被以递归的方法定义：F0=0，F1=1，Fn=F(n-1)+F(n-2)（n>=2，n∈N*）
*/
func PriFib(num int) {
	ch := make(chan int)

	go func(ch chan int, n int) {
		pre, cur := 0, 1
		for i := 0; i < n; i++ {
			ch <- cur
			pre, cur = cur, pre+cur
		}

		close(ch)
	}(ch, num)

	for i := range ch {
		fmt.Print(i)
	}
}

/*
回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

输入：x = 121
输出：true
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
输入：x = -101
输出：false

思路：反转后半数字，和前一半数字做比较
1、负数、和最后一位为0的数，不是回文数
2、x重复/10，为取末位数做准备
3、result 上一次的值 * 10，再加上x值的末位
4、若该回文数的长度为奇数，后半数字得 /10 去掉一位再和前半数字做比较
*/
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	result := 0
	for x > result {
		result = result*10 + x%10 // result 上一次的值 * 10，再加上x值的末位
		x = x / 10                // x重复/10，为取末位数做准备
	}

	return result == x || x == result/10
}

/*
给定一组非负整数 nums，重新排列它们每个数字的顺序（每个数字不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"

链接：https://leetcode-cn.com/problems/largest-number
*/
func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	o := strings.Join(ss, "")
	if o[0] == '0' {
		return "0"
	}
	return o
}

/*
求两数只和
https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}

	return result
}

/*
搜索插入位置
https://leetcode-cn.com/problems/search-insert-position/

思路
1、二分查找有序s数组
2、找出第一个>= target的数
*/
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid // target小于中间指针值，则将min赋值给result
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

/*
外观数列
https://leetcode-cn.com/problems/count-and-say/

思路
1、对countAndSay(n-1)返回的结果进行描述 则得到countAndSay(n)
2、递归终止条件：countAndSay(n-1) = "1"
*/
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	latest := countAndSay(n - 1)

	pre := latest[0]
	var count byte = '1'
	result := make([]byte, 0)
	for i := 1; i < len(latest); i++ {
		if latest[i] == pre {
			count += 1
			pre = latest[i]
		} else {
			result = append(result, count, pre)
			pre = latest[i]
			count = '1'
		}
	}

	result = append(result, count, pre)

	return string(result)
}

/*
求最大子序列和
https://leetcode-cn.com/problems/maximum-subarray/
*/
// 动态规划
/*
思路：
滚动数组
*/
func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] { // 若当前数字 + 前一个数字 > 当前数字，
			// 说明当前数字有增益效果，则滚动；否则，跳过
			nums[i] = nums[i] + nums[i-1]
		}

		if nums[i] > result { // 数组滚动后，和当前最大结果比较，取最大值
			result = nums[i]
		}
	}

	return result
}

/*
最后一个单词长度
https://leetcode-cn.com/problems/length-of-last-word/
思路
1、string转[]byte
2、倒序遍历
*/
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	strByte := []byte(s)
	for i := len(strByte) - 1; i >= 0; i-- {
		if string(strByte[i]) != " " {
			result++
			continue
		}

		if result != 0 && string(strByte[i]) == " " {
			break
		}
	}

	return result
}

/*
数组最后一个数加一
https://leetcode-cn.com/problems/plus-one/
*/
func plusOne(digits []int) []int {
	result := make([]int, 0)
	if len(digits) == 0 {
		return result
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] = digits[i] + 1
		}
		if digits[i] >= 10 && i > 0 {
			digits[i-1] = digits[i-1] + 1
			digits[i] = digits[i] - 10
		}

		if digits[i] >= 10 && i == 0 {
			digits[i] = digits[i] - 10
			digits = append([]int{1}, digits...)
		}
	}

	return digits
}

/*
二进制求和
https://leetcode-cn.com/problems/add-binary/
思路
1、两二进制尾部对其
2、定义一个变量存储是否进位
3、从尾部遍历最长的str，每位：(temp + str1[i] + str2[i]) % 2; temp = (temp + str1[i] + str2[i]) / 2
*/
func addBinary(a string, b string) string {
	maxStr := func(s1, s2 string) string {
		if len(s1) > len(s2) {
			return s1
		}
		return s2
	}(a, b)

	temp := 0
	result := ""
	for i := 0; i < len(maxStr); i++ {
		aItem := 0
		bItem := 0
		if len(a) > i {
			aItem = int(a[len(a)-1-i] - '0')
		}
		if len(b) > i {
			bItem = int(b[len(b)-1-i] - '0')
		}

		result = strconv.Itoa((temp+aItem+bItem)%2) + result
		temp = (temp + aItem + bItem) / 2
	}

	if temp == 1 {
		result = "1" + result
	}

	return result
}

/*
x 的平方根
https://leetcode-cn.com/problems/sqrtx/
思路：
1、定义一个数n从0到x
2、使得：n的平方 <= x && (n+1)的平方 > x
*/
func mySqrt(x int) int {
	if x < 0 {
		return 0
	}
	for n := 0; n <= x; n++ {
		if n*n <= x && (n+1)*(n+1) > x {
			return n
		}
	}

	return 0
}

/*
爬楼梯
https://leetcode-cn.com/problems/climbing-stairs/
思路
1、递归
2、f(x) = f(x-1) + f(x-2)
3、边界条件：f(0) = 0, f(1) = 1
*/
// 方式一：递归，会超时
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

// 方式二：可以看成是个斐波那契数列，滚动数组
/*
      *
(0 1) 1 2 3 5 8 13 21
*/
func climbStairs2(n int) int {
	a := 0      // 初始 前第二个数
	b := 1      // 初始 前第一个数
	result := 0 // 初始 当前值
	for i := 0; i < n; i++ {
		result = a + b
		a = b
		b = result
	}

	return result
}

/*
删除排序链表中的重复元素
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
思路：
	1、遍历列表，若当前节点 等于 前一个节点，则删除当前节点
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curNode := head // 1 、// 这一步：因为head是指针，curNode和head的内存地址是一样的
	for curNode.Next != nil {
		if curNode.Val == curNode.Next.Val {
			curNode.Next = curNode.Next.Next // 3、对curNode.Next赋值，即对head元素Next赋值，因为是同一块内存
		} else {
			curNode = curNode.Next // 2、curNode被赋值head的元素内存地址
		}
	}

	return head
}

/*
合并两个有序数组
https://leetcode-cn.com/problems/merge-sorted-array/
*/
/*
方式一：直接合并后排序
思路：
直接合并后排序
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

/*
方式二：双指针
*/
func merge2(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			result = append(result, nums2[p2:n]...)
			break
		}
		if p2 == n {
			result = append(result, nums1[p1:m]...)
			break
		}

		if nums1[p1] <= nums2[p2] {
			result = append(result, nums1[p1])
			p1++
		} else {
			result = append(result, nums2[p2])
			p2++
		}
	}

	copy(nums1, result)
}

/*
杨辉三角形
https://leetcode-cn.com/problems/pascals-triangle/
*/
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		inner := make([]int, 0)
		inner = append(inner, 1)
		if len(result) > 0 {
			lastInner := result[len(result)-1]
			j, k := 0, 1
			for k < len(lastInner) {
				inner = append(inner, lastInner[j]+lastInner[k])
				j++
				k++
			}
			inner = append(inner, 1)
		}
		result = append(result, inner)
	}

	return result
}

/*
杨辉三角 II
https://leetcode-cn.com/problems/pascals-triangle-ii/
思路：
1、通过递归的方式获取上一层结果
2、在上层结果的基础上基础本层结果
3、递归的终止条件：rowIndex == 0
*/
func getRow(rowIndex int) []int {
	if rowIndex <= 0 {
		return []int{1}
	}

	lastArray := getRow(rowIndex - 1)
	result := make([]int, 0)
	result = append(result, 1)

	i, j := 1, 0
	for i < len(lastArray) {
		result = append(result, lastArray[j]+lastArray[i])
		i++
		j++
	}

	result = append(result, 1)
	return result
}

/*
买卖股票的最佳时期
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/121-mai-mai-gu-piao-de-zui-jia-shi-ji-by-leetcode-/
思路
1、在最低点买入
2、买入后遍历，找受益最高值
*/
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, i := range prices {
		if i < minPrice {
			minPrice = i
		} else if i-minPrice > maxProfit {
			maxProfit = i - minPrice
		}
	}

	return maxProfit
}

/*
买卖股票最佳时间（可多次买卖）
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
方式一：动态规划
dp0[i]表示第i天手里没有股票的最大收益
dp1[i]表示第i天手里有股票的最大收益

第i天手里没有股票：dp0[i] = max{dp0[i-1], dp1[i-1]+prices[i]}
第i天手里有股票：dp1[i] = max{dp1[i-1], dp0[i-1]-prices[i]}
*/
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp0, dp1 := 0, -prices[0] // dp1以0为持仓成本
	for i := 1; i < len(prices); i++ {
		dp0 = max2(dp0, dp1+prices[i])
		dp1 = max2(dp1, dp0-prices[i])
	}

	return dp0 // 最后一天，手里没有股票，利润肯定是最大的
}

func max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
验证回文串
https://leetcode-cn.com/problems/valid-palindrome/
*/
func isPalindrome2(s string) bool {
	if len(s) <= 1 {
		return true
	}

	start, end := 0, len(s)-1
	for start < end {
		if !isAlNum(s[start]) {
			start++
			continue
		}
		if !isAlNum(s[end]) {
			end--
			continue
		}

		if strings.ToLower(string(s[start])) != strings.ToLower(string(s[end])) {
			return false
		}
		start++
		end--
	}

	return true
}

func isAlNum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') ||
		(ch >= 'a' && ch <= 'z') ||
		(ch >= '0' && ch <= '9')
}

/*
只出现一次的数字
https://leetcode-cn.com/problems/single-number/
异或运算（^）：两个数的二进制位不同时,当前位才置1 否则置0
	1、a ^ 0 = a
	2、a ^ a = 0
	3、满足结合律、交换律
思路
	1、遍历异或
	2、利用结合律、交换律的性质
*/
func singleNumber(nums []int) int {
	single := 0
	for i := 0; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}

/*
Excel表列名称
https://leetcode-cn.com/problems/excel-sheet-column-title/
思路：
	1、26进制转化
*/
func convertToTitle(columnNumber int) string {
	result := make([]byte, 0)
	for columnNumber > 0 {
		a := columnNumber % 26
		if a == 0 {
			a = 26                           // 整除时，余数强行等于26
			columnNumber = columnNumber - 26 // 向上接一位
		}
		result = append([]byte{byte(a) + 64}, result...)
		columnNumber = columnNumber / 26
	}

	return string(result)
}

/*
Excel 表列序号
https://leetcode-cn.com/problems/excel-sheet-column-number/
*/
func titleToNumber(columnTitle string) int {
	var result float64
	for i := 0; i < len(columnTitle); i++ {
		n := float64(columnTitle[i] - 64)
		result = result + n*math.Pow(26, float64(len(columnTitle)-1-i))
	}

	return int(result)
}

/*
多数元素
https://leetcode-cn.com/problems/majority-element/
*/
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	eMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		eMap[nums[i]] = eMap[nums[i]] + 1
	}

	maxKey := 0
	maxVal := 0
	for k, v := range eMap {
		if v > maxVal {
			maxVal = v
			maxKey = k
		}
	}

	return maxKey
}

func majorityElementV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	return nums[len(nums)/2]
}

/*
阶乘后的零
https://leetcode-cn.com/problems/factorial-trailing-zeroes/
*/
func trailingZeroes(n int) int {
	var ans int
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans

}

package algorithm

/*
动态规划算法
https://juejin.cn/post/6951922898638471181
*/

/*
62. 不同路径
https://leetcode-cn.com/problems/unique-paths/
思路：动态规划
1、到达某个格子路径f(i,j) = f(i-1,j) + j(i,j-1)
2、边界条件：第一行f(i,0) = 1，第一列f(0,j) = 1
3、填充m*n网格第一行，第一列的值，再填充其他值
4、返回f(i-1,j) + j(i,j-1)
*/
func uniquePaths(m int, n int) int { // m: 行；7：列
	if m == 1 || n == 1 {
		return 1
	}

	pathArray := make([][]int, 0)
	for mIndex := 0; mIndex < m; mIndex++ {
		row := make([]int, n)
		if mIndex == 0 {
			for nIndex := 0; nIndex < n; nIndex++ {
				row[nIndex] = 1
			}
		} else {
			row[0] = 1
		}
		pathArray = append(pathArray, row)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			pathArray[i][j] = pathArray[i-1][j] + pathArray[i][j-1]
		}
	}

	return pathArray[m-1][n-2] + pathArray[m-2][n-1]
}

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 10 级的台阶总共有多少种跳法。
思路：自顶向下
1、f（n） = f（n-1）+f(n-2)
2、带备忘录
*/
func numWays(n int) int {

	tempMap := make(map[int]int) // 备忘录
	var dp func(num int) int
	dp = func(num int) int {
		if num == 0 {
			return 1
		}
		if num <= 2 {
			return num
		}

		if result, ok := tempMap[num]; ok {
			return result
		} else {
			tempMap[num] = (dp(num-1) + dp(num-2)) % 1000000007
			return tempMap[num]
		}
	}
	return dp(n)
}

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
思路：自底向上
1、定义dp数组记录 i处最长递增子序列
2、定义一个数记录结果
3、遍历nums[0,i],若后数比前数大，dp[i] = Math.max(dp[i], dp[j] + 1); 0<=j<i
4、返回结果
*/
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	result := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		result = max(result, dp[i])
	}

	return result
}

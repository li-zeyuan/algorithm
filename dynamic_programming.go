package algorithm

import "math"

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
63. 不同路径 II
https://leetcode-cn.com/problems/unique-paths-ii/
思路：动态规划

时间复杂度：m * n
空间复杂度：m * n
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	dp := make([][]int, 0)
	dp = append(dp, make([]int, len(obstacleGrid[0])+1))
	for i := 1; i <= len(obstacleGrid); i++ {
		row := make([]int, len(obstacleGrid[0])+1)
		for j := 1; j <= len(obstacleGrid[0]); j++ {
			if i == 1 && j == 1 && obstacleGrid[i-1][j-1] != 1 {
				row[j] = 1
			} else if obstacleGrid[i-1][j-1] == 1 {
				row[j] = 0
			} else {
				row[j] = row[j-1] + dp[i-1][j]
			}
		}
		dp = append(dp, row)
	}

	return dp[len(obstacleGrid)][len(obstacleGrid[0])]
}

/*
64. 最小路径和
https://leetcode-cn.com/problems/minimum-path-sum/
思路：动态规划
1、构造dp二维数组
2、f(i,j) = min(f(i-1,j), j(i,j-1)) + f(i,j)

时间复杂度：m * n
空间复杂度：m * n
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := make([][]int, len(grid)+1)
	firstRow := make([]int, len(grid[0])+1)
	for i2 := 0; i2 < len(grid[0])+1; i2++ {
		if i2 == 1 {
			firstRow[i2] = 0
		} else {
			firstRow[i2] = math.MaxInt32
		}
	}
	dp[0] = firstRow

	for i := 1; i <= len(grid); i++ {
		dpRow := make([]int, len(grid[0])+1)
		dpRow[0] = math.MaxInt32
		for j := 1; j <= len(grid[0]); j++ {
			dpRow[j] = minPathSumMin(dpRow[j-1], dp[i-1][j]) + grid[i-1][j-1]
		}
		dp[i] = dpRow
	}

	return dp[len(grid)][len(grid[0])]
}

func minPathSumMin(x, y int) int {
	if x > y {
		return y
	}

	return x
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

/*
1143. 最长公共子序列
https://leetcode-cn.com/problems/longest-common-subsequence/
思路：
1、构建二维数组dp
2、自底向上遍历，text1用i指针，text2用j指针
3、若text1[i] == text2[j],dp[i][j] = dp[i-1][j-1] + 1。既斜对角加1
4、若text1[i] != text2[j],dp[i][j] = max(dp[i-1][j],dp[i][j-1])。既相邻最大值
5、返回dp[i][j]
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := make([][]int, len(text1)+1)
	dp[0] = make([]int, len(text2)+1)
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if j == 0 {
				// todo
				dp[i+1] = make([]int, len(text2)+1)
			}

			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

/*
139. 单词拆分
https://leetcode-cn.com/problems/word-break/
思路：动态规划
index：	0	1	2	3	4	5	6	7	8
s：			l	e	e	t	c	o	d	e
dp：   true		   	   true			   true

dp记录之前的状态
*/
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

/*
806. 写字符串需要的行数
https://leetcode-cn.com/problems/number-of-lines-to-write-string/
思路：直接遍历
*/
func numberOfLines(widths []int, s string) []int {
	const maxWidth = 100
	lines, width := 1, 0
	for _, c := range s {
		need := widths[c-'a']
		width += need
		if width > maxWidth {
			lines++
			width = need
		}
	}
	return []int{lines, width}
}

/*
45. 跳跃游戏 II
https://leetcode-cn.com/problems/jump-game-ii/
思路：动态规划
1、dp数组保存该下标下，最小的步数
2、遍历nums数组，取出元素，向前修改dp数组

时间复杂度：n^2
空间复杂度：1
*/
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] > 1 {
			return 1
		} else {
			return 0
		}
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		step := nums[i]
		for j := i + 1; step > 0 && j < len(nums); j++ {
			step--
			if i == 0 {
				dp[j] = 1
				continue
			}
			if dp[j] == 0 {
				dp[j] = dp[i] + 1
			}
		}
	}

	return dp[len(nums)-1]
}

/*
思路：正向查找可到达的最大位置
https://leetcode-cn.com/problems/jump-game-ii/solution/tiao-yue-you-xi-ii-by-leetcode-solution/
时间复杂度：n
空间复杂度：1
*/
func jump2(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = jumpMax(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func jumpMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
821. 字符的最短距离
https://leetcode-cn.com/problems/shortest-distance-to-a-character/
思路：双指针

时间复杂度：n^2
空间复杂度：n
*/
func shortestToChar(s string, c byte) []int {
	if len(s) == 0 {
		return nil
	}

	result := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			result[i] = 0
			continue
		}

		j := i
		for {
			if (2*i-j) >= 0 && s[2*i-j] == c { // 前移指针
				result[i] = j - i
				break
			}
			if j < len(s) && s[j] == c { // // 后移指针
				result[i] = j - i
				break
			}

			j++
		}
	}

	return result
}

/*
思路2：两次遍历
1、s[i] 到其左侧最近的字符 cc 的距离
2、 s[i]s[i] 到其右侧最近的字符 cc 的距离

时间复杂度：n
空间复杂度：1
https://leetcode-cn.com/problems/shortest-distance-to-a-character/solution/zi-fu-de-zui-duan-ju-chi-by-leetcode-sol-2t49/
*/
func shortestToChar2(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)

	idx := -n
	for i, ch := range s {
		if byte(ch) == c {
			idx = i
		}
		ans[i] = i - idx
	}

	idx = n * 2
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		ans[i] = shortestToCharMin(ans[i], idx-i)
	}
	return ans
}

func shortestToCharMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

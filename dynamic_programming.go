package algorithm

/*
动态规划算法
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

	for i := 1 ; i < m ; i ++ {
		for j := 1; j < n; j ++ {
			pathArray[i][j] = pathArray[i-1][j] + pathArray[i][j - 1]
		}
	}

	return pathArray[m-1][n-2] + pathArray[m-2][n-1]
}

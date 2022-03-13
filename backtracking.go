package algorithm

/*
回溯算法
*/
/*
46. 全排列
https://leetcode-cn.com/problems/permutations/
思路：回溯法
	res: 结果集
	path: 经过的路径节点
	visited：记录可用的节点
*/
func permute(nums []int) [][]int {
	var res [][]int
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}

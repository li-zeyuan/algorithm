package algorithm

/*
22. 括号生成
https://leetcode-cn.com/problems/generate-parentheses/
思路：回溯
选择
	1、要么选 ( 要么选 )
约束条件
	1、有（ 优先选 (
	2、）数量大于 ( ，才能选 ）
目标
	1、长度等于2n时，加入结果集

题解：https://leetcode-cn.com/problems/generate-parentheses/solution/shou-hua-tu-jie-gua-hao-sheng-cheng-hui-su-suan-fa/
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	result := make([]string, 0)
	var dfs func(left, right int, s string)
	dfs = func(left, right int, s string) {
		if len(s) == 2*n { // 长度等于2n时，加入结果集
			result = append(result, s)
			return
		}

		if left > 0 { // 有（ 优先选 (
			dfs(left-1, right, s+"(")
		}
		if right > left { // ）数量大于 ( ，才能选 ）
			dfs(left, right-1, s+")")
		}
	}
	dfs(n, n, "")

	return result
}

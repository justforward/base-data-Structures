package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	if n == 0 {
		return result
	}

	// 执行深度优先遍历，搜索可能的结果
	dfs("", n, n, &result)
	return result
}

/**
 * @param curStr 当前递归得到的结果
 * @param left   左括号还有几个可以使用
 * @param right  右括号还有几个可以使用
 * @param res    结果集
 */
func dfs(curStr string, left, right int, res *[]string) {
	// 因为每一次尝试，都使用新的字符串变量，所以无需回溯
	// 在递归终止的时候，直接把它添加到结果集即可
	if left == 0 && right == 0 {
		*res = append(*res, curStr)
		return
	}

	// 剪枝（如图，左括号可以使用的个数严格大于右括号可以使用的个数，才剪枝）
	// 左括号可
	if left > right {
		return
	}

	if left > 0 {
		dfs(curStr+"(", left-1, right, res)
	}

	if right > 0 {
		dfs(curStr+")", left, right-1, res)
	}
}

func main() {
	n := 3
	result := generateParenthesis(n)
	fmt.Println(result)
}

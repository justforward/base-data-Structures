package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfs("", n, n, &res)
	return res
}

func dfs(value string, left, right int, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, value)
		return
	}
	if left > right {
		return
	}
	if left > 0 {
		dfs(value+"(", left-1, right, res)
	}
	if right > 0 {
		dfs(value+")", left, right-1, res)
	}
}

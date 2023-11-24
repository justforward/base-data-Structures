package main

var (
	path []string
	res  [][]string
)

// aab截取
func partition(s string) [][]string {

}

func dfs(s string, index int) {
	if index >= len(s) {
		// 查看是否为
	}

}

func isValid(s string) bool {
	left, right := 0, 0
	for left < right {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

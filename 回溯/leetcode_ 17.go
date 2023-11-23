package main

//
//var (
//	m    []string
//	path []byte
//	res  []string
//)
//
//// 本题目中每个树节点代表的是不同的集合 那么for循环中需要遍历的东西还是需要改变的
//func letterCombinations(digits string) []string {
//
//	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
//	path = make([]byte, 0)
//	res = make([]string, 0)
//	if digits == "" {
//		return res
//	}
//	dfs(digits, 0)
//	return res
//}
//
//func dfs(digits string, start int) {
//	if len(path) == len(digits) {
//		tmp := string(path)
//		res = append(res, tmp)
//		return
//	}
//
//	digit := int(digits[start] - '0')
//	str := m[digit-2]
//	for i := 0; i < len(str); i++ {
//		path = append(path, str[i])
//		dfs(digits, start+1)
//		path = path[:len(path)-1]
//	}
//
//}

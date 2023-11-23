package main

/*var (
	path []int
	res  [][]int
)

func combine(n int, k int) [][]int {
	path = make([]int, 0)
	res = make([][]int, 0)
	backed(n, k, 1)
	return res
}

func backed(n, k, start int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path) // 这步是为了确保以后对tmp slice的操作不会影响到最后结果
		res = append(res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		// 剩下的数量个数，已经不满足条件 为什么要+1 因为这个是一个左闭的集合，从
		if n-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		backed(n, k, i+1)
		path = path[:len(path)-1]
	}
}

func main() {

}
*/

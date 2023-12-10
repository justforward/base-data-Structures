package main

// 需要进行一个去重的过程，go中不存在set集合，需要使用map进行去重的操作
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	tmpMap := make(map[int]int, 0)

	for _, v := range nums1 {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v]++
		}
	}

	for _, m := range nums2 {
		if _, ok := tmpMap[m]; ok {
			res = append(res, m)
			delete(tmpMap, m)
		}
	}
	return res
}
func main() {

}

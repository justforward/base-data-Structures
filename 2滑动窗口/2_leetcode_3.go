package main

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		m[c]++
		for m[c] > 1 {
			d := s[left]
			left++
			m[d]--
		}
		res = max(res, right-left)

	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

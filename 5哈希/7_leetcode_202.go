package main

func isHappy(n int) bool {
	m := make(map[int]bool)
	for !m[n] {
		m[n] = true
		n = sum(n)
		if n == 1 {
			return true
		}
	}
	return false
}

func sum(m int) int {
	res := 0
	for m > 0 {
		n := m % 10
		res += n * n
		m /= 10
	}
	return res
}

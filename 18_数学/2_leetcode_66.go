package main

func plusOne(digits []int) []int {

	res := make([]int, len(digits))
	plus := 0
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + plus
		if i == len(digits)-1 {
			num++
		}
		plus = num / 10
		res[i] = num % 10
	}
	if plus != 0 {
		res = append([]int{plus}, res...)
	}

	return res
}

func main() {
	dig := []int{9}
	plusOne(dig)
}

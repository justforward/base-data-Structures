package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			// 注意：最新的一个数组出现在栈顶，所以需要num1=-2 num2=-1
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			tmpVal := 0
			switch token {
			case "+":
				tmpVal = num1 + num2
			case "-":
				tmpVal = num1 - num2
			case "*":
				tmpVal = num1 * num2
			case "/":
				tmpVal = num1 / num2
			}
			stack = append(stack, tmpVal)
		}
	}
	return stack[0]
}

package main

import "math"

/*

 快速幂+递归
 每次都根据结果再次
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == math.MaxInt32 {
		return myPow(1/x, -(n+1)) / x
	}

	if n < 0 {
		return myPow(1/x, -n)
	}
	if n%2 == 1 {
		return (x * myPow(x, n-1))
	} else {
		sub := myPow(x, n/2)
		return (sub * sub)
	}
}

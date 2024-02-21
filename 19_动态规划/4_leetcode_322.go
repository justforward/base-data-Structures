package main

import "math"

// dp[j] 凑足amout的最小硬币数量
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1) //
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

package _背包问题

import "math"

/*
	如何转换成背包问题？
	假设加法的总和为x， 那么减法对应的总和是sum-x
    我们要求 x-（sum-x）=target
    x=（target+sum）/2hnj

 	这个题目转成装满 背包容量=x的时候，有多少种方法？
    1）dp[j] 装满容量为j 有dp[j]种方法
	2）dp[j]+=dp[j-num[i]] 将所有的情况都加起来
	3）初始化 dp[0]=1 ? 为什么等于1 装满0容量的 也就是一种方法，【0】
*/

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if abs(target) > sum {
		return 0
	}

	if (target+sum)%2 != 0 {
		return 0
	}
	// 计算背包大小
	bag := (sum + target) / 2
	// 定义dp数组
	dp := make([]int, bag+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			//推导公式
			dp[j] += dp[j-nums[i]]
			//fmt.Println(dp)
		}
	}
	return dp[bag]
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

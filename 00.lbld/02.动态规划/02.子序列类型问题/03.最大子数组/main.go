package main

import (
	"fmt"
	"math"
)

func main() {

	// 01 最大子序和。 给定一个整数数组nums, 找到一个具有最大和的连续子数组（至少包含一个元素），返回其最大和
	fmt.Println(maxSubArray([]int{1, 2, -3, 4, -6, 9, -8, 5, 5}))
}

func maxSubArray(nums []int) int {

	var dp = make([]int, len(nums))

	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {

		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	var res = math.MinInt
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

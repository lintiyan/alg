package main

import "fmt"

func main() {

	// 01. 给定一个只包含正整数的非空数组，是否可以将这个数组分割成两个子集，使得这两个子集的元素和相等。
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5, 5, 6, 100}))
}

func canPartition(nums []int) bool {

	var sum int
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	var target = sum / 2

	// 此时，问题就变为了，能否从nums数组中选出和为target的子数组
	res1 := getSubArrayTargetFromTop(nums, target, 0)
	res2 := getSubArrayTargetFromBottom(nums, target)
	fmt.Println("res1=", res1, "res2=", res2)

	return res1 || res2
}

// 自顶向下的解法
func getSubArrayTargetFromTop(nums []int, target int, i int) bool {

	if target == 0 {
		return true
	}
	if i >= len(nums) {
		return false
	}

	return getSubArrayTargetFromTop(nums, target-nums[i], i+1) || getSubArrayTargetFromTop(nums, target, i+1)
}

func getSubArrayTargetFromBottom(nums []int, target int) bool {

	// dp数组的含义是 对于nums[...i], 是否有子数组的和值为target
	var dp [][]bool
	for i := 0; i < len(nums); i++ {
		item := make([]bool, target+1)
		dp = append(dp, item)
	}

	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}

	for t := 0; t <= target; t++ {
		if t == nums[0] {
			dp[0][t] = true
		}
	}

	for i := 1; i < len(nums); i++ {
		for t := 0; t <= target; t++ {
			var temp bool
			if t >= nums[i] {
				temp = dp[i-1][t] || dp[i-1][t-nums[i]]
			} else {
				temp = dp[i-1][t]
			}
			dp[i][t] = temp
		}
	}

	return dp[len(nums)-1][target]
}

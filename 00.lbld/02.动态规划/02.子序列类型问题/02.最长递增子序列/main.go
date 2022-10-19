package main

import (
	"fmt"
	"sort"
)

func main() {

	// 01. 输入一个无序的整数数组，请你找到其中最长的严格递增子序列的长度
	fmt.Println(lengthOfLIS([]int{}))

	// 02. 俄罗斯套娃信封问题。 给你一个二维整数数组envelopes，其中envelopes[i] = [Wi,Hi],表示第i个信封的宽度和长度
	// 当另一个信封的宽度和长度都比这个信封大时，这个信封就可以被放进另一个信封里。
	// 计算最多能有多少个信封能组成一组“俄罗斯套娃”信封
	fmt.Println(maxEnvelopes([][]int{[]int{1, 2}, {2, 3}, {2, 3}, {2, 4}, {2, 5}, {4, 6}}))
}

func lengthOfLIS(nums []int) int {

	// 定义状态
	// 状态是i，数组的下标。 i从0到len(nums)变化对应不同的递增子序列长度

	// 定义dp
	// dp[i]表示，以nums[i]结尾的最长子序列长度

	// 明确选择
	// 从dp[0]到dp[i-1]能推出dp[i]么？可以的， 只要找到nums[0]到nums[i]中，值比nums[i]小的数对应的dp值最大的加1就可以

	if len(nums) == 0 {
		return 0
	}

	var dp = make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		var max int
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j] > max {
				max = dp[j]
			}
		}

		dp[i] = max + 1
	}

	var res int
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

func maxEnvelopes(envelopes [][]int) int {

	// 宽度按从小到达排序，然后从高度的数组中查找最长递增子序列就行。
	// 为了避免出现宽度相同、高度递增的信封被选中，将宽度相同的信封按照高度的降序排列

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] > envelopes[j][0] {
			return false
		} else {
			if envelopes[i][1] > envelopes[j][1] {
				return true
			} else {
				return false
			}
		}
	})

	var nums = make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		nums[i] = envelopes[i][1]
	}

	return lengthOfLIS(nums)
}

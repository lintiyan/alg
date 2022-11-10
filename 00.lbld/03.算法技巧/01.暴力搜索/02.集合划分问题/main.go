package main

import "fmt"

func main() {

	// 01. 排列组合中的回溯思想
	fmt.Println(canPartitionKSubSets([]int{1,5,3,3,3,3},2))
}

func canPartitionKSubSets(nums []int, k int) bool {

	if len(nums) <= 0 {
		return false
	}

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%k != 0 {
		return false
	}

	target := sum / k

	var kSum = make([]int, k)
	return canPartitionKSubSetFromNum(0, nums, kSum, target)
}

// 回溯算法的关键：
// 路径 - 已经走过的路径
// 选择 - 有哪些选择，如何选择和取消选择
// 结束 - 到达决策树底层，没有选择的时候，或者说选择结束该看结果的时候到了

// 回溯算法与动态规划的区别是，动态规划只需要走满足条件的分支，而回溯需要全量遍历一遍

// 从数字的角度讲，每个数都有K种选择
func canPartitionKSubSetFromNum(i int, nums []int, kSum []int, target int) bool {

	// 选择结束，该看结果了
	if i == len(nums) {
		// k个桶中的值都等于target才算成功
		for j := 0; j < len(kSum); j++ {
			if kSum[j] != target {
				return false
			}
		}
		return true
	}

	// 每个数字有K中选择，选择谁都可以，无非是遍历而已。
	// “选谁都能继续走下去，但是不是选谁都能得到想要的结果”
	for j := 0; j < len(kSum); j++ {

		// 剪枝
		if kSum[j] + nums[i] > target {
			continue
		}
		// 第i个数选择放入第j个桶
		kSum[j] += nums[i]
		if canPartitionKSubSetFromNum(i+1, nums, kSum, target){
			return true
		}
		// 撤销选择
		kSum[j] -= nums[i]
	}

	return false
}

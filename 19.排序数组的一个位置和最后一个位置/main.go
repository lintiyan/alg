package main

import "fmt"

/*

	给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

	如果数组中不存在目标值 target，返回 [-1, -1]。

	你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

	 

	示例 1：

	输入：nums = [5,7,7,8,8,10], target = 8
	输出：[3,4]
	示例 2：

	输入：nums = [5,7,8,8,8,8,10], target = 6
	输出：[-1,-1]
	示例 3：

	输入：nums = [], target = 0
	输出：[-1,-1]
	 

	提示：

	0 <= nums.length <= 105
	-109 <= nums[i] <= 109
	nums 是一个非递减数组
	-109 <= target <= 109

*/

func searchRange(nums []int, target int) []int {

	left := searchFirst(nums, target)
	right := searchFirst(nums, target + 1)

	if left == len(nums)  || nums[left] != target {
		return []int{-1,-1}
	}
	return []int{left, right -1}
}

// 查找大于等于target的第一个值
func searchFirst(nums []int, target int)int {
	var left = 0
	var right = len(nums)

	for left < right {
		temp := (left + right) /2
		if nums[temp] >= target {
			right = temp
		} else if nums[temp] < target {
			left = temp + 1
		}
	}

	return left
}

func main() {

	//var nums = []int{5,7,8,8,8,8,10}
	var nums = []int{5,7,7,8,8,10}
	//var nums = []int{1, 1,3,3}
	//var nums = []int{2,2}
	fmt.Println(searchRange(nums,8))
}

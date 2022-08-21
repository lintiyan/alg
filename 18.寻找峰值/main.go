package main

import "fmt"

/*

	峰值元素是指其值严格大于左右相邻值的元素。

	给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

	你可以假设 nums[-1] = nums[n] = -∞ 。

	你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

	 

	示例 1：

	输入：nums = [1,2,3,1]

	输出：2
	解释：3 是峰值元素，你的函数应该返回其索引 2。
	示例 2：

	输入：nums = [1,2,1,3,5,6,4]
	输出：1 或 5
	解释：你的函数可以返回索引 1，其峰值元素为 2；
	     或者返回索引 5， 其峰值元素为 6。

*/


func findPeakElement(nums []int) int {
	var left = 0
	var right = len(nums) - 1

	for left <= right {
		temp := (left + right) / 2
		fmt.Println("temp=",temp, "left=", left, "right=", right)

		// 如果temp元素大于它左右两边的元素，那么这个元素就是峰值
		if ((temp+1 > len(nums) -1) || nums[temp] > nums[temp+1]) && ( temp-1 < 0 || nums[temp] > nums[temp -1]) {
			return temp
		}

		if nums[temp] < nums[temp+1] {
			left = temp + 1
		} else if nums[temp] > nums[temp+1] {
			right = temp
		}

	}

	return 0
}

func findPeakElement2(nums []int) int {
	var left = 0
	var right = len(nums) - 1

	for left < right {
		temp := (left + right) >> 1
		fmt.Println("temp=",temp, "left=", left, "right=", right)

		if nums[temp] < nums[temp+1] {
			left = temp + 1
		} else if nums[temp] > nums[temp+1] {
			right = temp
		}
	}
	return left
}

func main() {

	nums := []int{8,6,4,9,2,3,4,1,9,76,66,55,22}
	//nums := []int{1,2,1,3,5,6,4}
	//nums:= []int{11,2}
	fmt.Println(findPeakElement(nums))
	fmt.Println(findPeakElement2(nums))
}

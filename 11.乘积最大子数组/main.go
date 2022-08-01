package main

import "fmt"

/*

	给你一个整数数组 nums，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

	测试用例的答案是一个32-位 整数。

	子数组 是数组的连续子序列。


	示例 1:

	输入: nums = [2,3,-2,4]
	输出: 6
	解释:子数组 [2,3] 有最大乘积 6。
	示例 2:

	输入: nums = [-2,0,-1]
	输出: 0
	解释:结果不能为 2, 因为 [-2,-1] 不是子数组。

*/

func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var max = make([]int, len(nums))
	var min = make([]int, len(nums))

	max[0] = nums[0]
	min[0] = nums[0]

	var maxNum = max[0]
	for i := 1 ; i < len(nums) ; i++ {
		if nums[i] > 0 {

			if max[i-1] > 0 {
				max[i] = nums[i] * max[i-1]
			}else if max[i-1] <= 0{
				max[i] = nums[i]
			}

			if min[i-1] < 0 {
				min[i] = nums[i] * min[i-1]
			} else if min[i-1] >= 0 {
				min[i] = nums[i]
			}

		}else if nums[i] < 0 {
			if min[i-1] < 0 {
				max[i] = nums[i] * min[i-1]
			}else if min[i-1] >= 0 {
				max[i] = nums[i]
			}

			if max[i-1] > 0 {
				min[i] = nums[i] * max[i-1]
			}else if max[i-1] <= 0 {
				min[i] = nums[i]
			}
		}else {
			max[i] = 0
			min[i] = 0
		}

		if max[i] > maxNum {
			maxNum = max[i]
		}

		fmt.Println("i=",i,"max=",max)
		fmt.Println("i=",i,"min=",min)
	}

	return maxNum
}

func maxProduct2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var maxPre = nums[0]
	var minPre = nums[0]
	var maxNum = nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		var tempMax = maxPre
		var tempMin = minPre
		if nums[i] > 0 {

			if tempMax > 0 {
				maxPre = nums[i] * tempMax
			}else if tempMax <= 0{
				maxPre = nums[i]
			}

			if tempMin < 0 {
				minPre = nums[i] * tempMin
			} else if tempMin >= 0 {
				minPre = nums[i]
			}

		}else if nums[i] < 0 {
			if tempMin < 0 {
				maxPre = nums[i] * tempMin
			}else if tempMin >= 0 {
				maxPre = nums[i]
			}

			if tempMax > 0 {
				minPre = nums[i] * tempMax
			}else if tempMax <= 0 {
				minPre = nums[i]
			}
		}else {
			maxPre = 0
			minPre = 0
		}

		if maxPre > maxNum {
			maxNum = maxPre
		}
	}

	return maxNum
}

func main() {

	//var nums = []int{-2,0,-1}
	//var nums = []int{2,3,-2,4}
	var nums = []int{2,-3,-2,-4}
	//var nums = []int{-2,3,-4}
	fmt.Println(maxProduct2(nums))
}
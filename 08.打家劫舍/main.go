package main

import (
	"fmt"
	"math"
)

/*

	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
	如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

	示例 1：

	输入：[1,2,3,1]
	输出：4
	解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。 偷窃到的最高金额 = 1 + 3 = 4 。


	示例 2：

	输入：[2,7,9,3,1]
	输出：12
	解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。 偷窃到的最高金额 = 2 + 9 + 1 = 12 。


	思路：
		先设定一个数组arr， arr[i]代表抢到第i家时，能获取到的最大利益。分析下arr[i]的取值，
		可以抢第i家，则arr[i]的取值是 抢到i-2家的最大收益加上当前第i家的财富值，即arr[i] = arr[i-2] + nums[i]
		可以不抢第i家，arr[i] = arr[i-1]
		那么arr[i] = max(arr[i-2] + nums[i], arr[i-1])

*/

func main() {

	var nums1 = []int{2,7,9,3,1}
	fmt.Println(rob(nums1))
	fmt.Println(rob2(nums1))
	var nums2 = []int{1,2,3,1}
	fmt.Println(rob(nums2))
	fmt.Println(rob2(nums2))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var robRes = make([]int, len(nums))
	robRes[0] = nums[0]
	robRes[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i:= 2 ; i < len(nums) ; i++ {
		robRes[i] = int(math.Max( float64(robRes[i-2]+ nums[i]), float64(robRes[i-1])))
	}

	var max int
	for _, item := range robRes {
		if item > max {
			max = item
		}
	}

	return max
}


func rob2(nums []int)int{
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var pre1, pre2 , max int

	pre1 = nums[0]
	pre2 = int(math.Max(float64(nums[0]), float64(nums[1])))
	max = int(math.Max(float64(pre1), float64(pre2)))

	for i:= 2 ; i < len(nums) ; i++ {
		max = int(math.Max( float64(pre1 + nums[i]), float64(pre2)))
		pre1 = pre2
		pre2 = max
	}
	return max
}
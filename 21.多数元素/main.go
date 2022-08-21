package main

import "fmt"

/*

	给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

	你可以假设数组是非空的，并且给定的数组总是存在多数元素。

	 

	示例 1：

	输入：nums = [3,2,3]
	输出：3
	示例 2：

	输入：nums = [2,2,1,1,1,2,2]
	输出：2
	 

	提示：
	n == nums.length
	1 <= n <= 5 * 104
	-109 <= nums[i] <= 109
	 

	进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题


*/

func majorityElement(nums []int) int {

	// 候选值，及其剩余票数
	var candidate, times int

	for i := 0; i < len(nums); i++ {

		// 如果票数为0，证明上一个候选人被淘汰，产生新的候选人
		if times == 0 {
			candidate = nums[i]
			times++
		}else {
			// 票数不为0，证明当前候选人目前为止未被淘汰，剩余times票
			if nums[i] == candidate {
				// 当前nums[i] = candidate, 候选值的票数加1
				times++
			}else {
				// 当前nums[i] != candidate, 候选人票数减1
				times--
			}
		}
	}
	return candidate
}

func main() {

	//var nums = []int{2,2,1,1,1,2,2}
	//var nums = []int{3,2,3}
	var nums = []int{3,3,4}
	fmt.Println(majorityElement(nums))
}

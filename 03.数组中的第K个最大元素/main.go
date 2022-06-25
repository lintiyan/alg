package main

import "fmt"

/**

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

 

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
 

提示：

1 <= k <= nums.length <= 104
-104 <= nums[i] <= 104


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/

func findKthLargest(nums []int, k int) int {
	return QuickSelect(nums, 0, len(nums) - 1, k)
}

// 基于快排思想
func QuickSelect(sli []int, left, right int,  k int) int {

	// 进行一轮快排，得到这一轮快排确定位置的元素的下标
	index := QuickPartition(sli, left, right)

	// 如果这轮确定位置的元素下标+1 后等于k，直接返回下标对应的值，加1是因为下标+1后才是排名
	if index + 1 == k {
		return sli[index]
	}

	// 如果这轮确定位置的元素下标小于k， 则要找的元素在该下标的右侧，针对该下标右侧的元素数组进行QuickSelect
	if index < k {
		return QuickSelect(sli, index + 1, right, k)
	}

	// 如果这轮确定位置的元素下标大于k， 则要找的元素在该下标的左侧，针对该下标左侧的元素数组进行QuickSelect
	return QuickSelect(sli, left, index - 1, k)
}

// 对sli进行快排调整操作，返回当前调整元素的下标
func QuickPartition(sli []int, left, right int) int {

	if left >= right {
		return left
	}

	t := sli[left]
	var i, j = left, right

	for i < j {

		for i < j {
			if sli[j] > t {
				sli[i] = sli[j]
				i++
				break
			}else{
				j--
			}
		}

		for i < j {
			if sli[i] < t {
				sli[j] = sli[i]
				j--
				break
			}else{
				i++
			}
		}
	}

	// 找到t的位置
	sli[i] = t
	return i
}

func main() {
	var arr = []int{3,2,3,1,2,4,5,5,6}
	fmt.Println(findKthLargest(arr, 4))

	var arr2 = []int{3,2,1,5,6,4}
	fmt.Println(findKthLargest(arr2, 2))
}

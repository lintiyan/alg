package main

import "fmt"

func main() {

	nums := []int{1, 2, 2, 2, 3, 5, 6, 8, 9}
	// 01. 寻找有序数组中的一个数
	fmt.Println(binarySearch(nums, 8))

	// 02. 寻找有序数组中的一个数，如果有多个，返回最左边那个数的下标
	fmt.Println(leftBound(nums, 2))
	fmt.Println(leftBound(nums, 0))
	fmt.Println(leftBound(nums, 10))

	// 03. 寻找有序数组中的一个数，如果有多个，返回最右边那个数的下标
	fmt.Println(rightBound(nums, 2))
	fmt.Println(rightBound(nums, 0))
	fmt.Println(rightBound(nums, 10))
}

func binarySearch(nums []int, target int) int {

	// 左闭右闭区间
	var left = 0
	var right = len(nums) - 1

	for left <= right {
		var mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	return -1
}

// 1,2,2,2,2,2,3
func leftBound(nums []int, target int) int {

	var left = 0
	var right = len(nums) - 1

	for left <= right {

		var mid = left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	// 如果target比所有数小，或者target比所有数大，则返回-1
	if left >= len(nums) || nums[left] != target {
		return -1
	}

	return left
}

// 1,2,2,2,2,2,3
func rightBound(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1

	for left <= right {

		var mid = left + (right-left)/2

		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	// target比数组中所有的数都小，或者target比数组中所有的数都大
	if  right < 0  || nums[right] != target{
		return -1
	}

	return left - 1
}

package main

import "fmt"

/*

	给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

	请你设计并实现时间复杂度O(n) 的算法解决此问题。


	示例 1：

	输入：nums = [100,4,200,1,3,2]
	输出：4
	解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
	示例 2：

	输入：nums = [0,3,7,2,5,8,4,6,0,1]
	输出：9

*/

func longestConsecutive(nums []int) int {

	var ma = make(map[int]int, len(nums))
	var max int

	for _, num := range nums {

		if _, ok := ma[num]; ok {
			continue
		}

		left, isLeftExist := ma[num-1]
		right, isRightExist := ma[num+1]

		ma[num] = left + right + 1
		if ma[num] > max {
			max = ma[num]
		}

		if isLeftExist {
			ma[num-left] = ma[num]
		}
		if isRightExist {
			ma[num+right] = ma[num]
		}
		fmt.Println(ma)
	}
	return max
}

func main() {
	// var nums = []int{0,3,7,2,5,8,4,6,0,1}
	//var nums = []int{100,4,200,1,3,2}
	// var nums = []int{100,4,200,1,3,2,6,5,9,8}
	var nums = []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(nums))
}

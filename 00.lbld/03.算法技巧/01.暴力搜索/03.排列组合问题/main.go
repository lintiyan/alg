package main

import "fmt"

func main() {

	// 01. 给一个无重复元素的数组nums，每个元素最多使用一次，请返回nums所有的子集
	fmt.Println(SubSets([]int{1, 2, 3}))

	// 02. 给定两个整数n 和 K，返回范围[1,n]中，所有可能的K个数的组合
	fmt.Println(Combine(3,2))
}

type Result struct {
	ResAll [][]int
}

func SubSets(nums []int) [][]int {
	var result = &Result{}
	SubSetsBasic(nums, 0, []int{}, result)
	return result.ResAll
}

/*
					 []
			   /      |	     \
		   [1]     	  [2]  	   [3]
		  /   \        |
	  [1,2]   [1,3]   [2,3]
     /
[1,2,3]

*/
func SubSetsBasic(nums []int, start int, res []int, result *Result) {

	// 树上的每个节点都是子集
	result.ResAll = append(result.ResAll, copyRes(res))

	// 有j到len(nums)-1种选择
	for j := start; j < len(nums); j++ {
		// 选择nums[j]
		res = append(res, nums[j])

		// 此时只有j+1到len(nums)-1种选择
		SubSetsBasic(nums, j+1, res, result)

		// 撤销nums[j]
		res = res[:len(res)-1]
	}
}

func copyRes(src []int) []int {
	var dst = make([]int, len(src))
	copy(dst, src)
	return dst
}

func Combine(n int, k int) [][]int {
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	var result = &Result{}
	CombineBasic(k, nums, 0, []int{}, result)
	return result.ResAll
}
func CombineBasic(k int, nums []int, start int, res []int, result *Result) {
	if len(res) == k {
		result.ResAll = append(result.ResAll, copyRes(res))
	}

	for i := start; i < len(nums); i++ {
		res = append(res, nums[i])
		CombineBasic(k, nums, i+1, res, result)
		res = res[:len(res)-1]
	}
}

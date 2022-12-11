package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	// 01. 给一个无重复元素的数组nums，每个元素最多使用一次，请返回nums所有的子集
	fmt.Println(SubSets([]int{1, 2, 3}))

	// 02. 给定两个整数n 和 K，返回范围[1,n]中，所有可能的K个数的组合
	fmt.Println(Combine(3, 2))

	// 03. 给定一个不重复数组的数组，返回其所有可能的全排列
	fmt.Println(permute([]int{1,2,3}))
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
	result.ResAll = append(result.ResAll, basicStruct.CopyRes(res))

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
		result.ResAll = append(result.ResAll, basicStruct.CopyRes(res))
	}

	for i := start; i < len(nums); i++ {
		res = append(res, nums[i])
		CombineBasic(k, nums, i+1, res, result)
		res = res[:len(res)-1]
	}
}

func permute(nums []int) [][]int {
	var result = &Result{}
	permuteBasic(nums, []int{}, result)
	return result.ResAll
}
func permuteBasic(nums []int, res []int, result *Result) {

	if len(res) == len(nums){
		result.ResAll = append(result.ResAll, basicStruct.CopyRes(res))
		return
	}

	for i:=0 ;i<len(nums);i++{
		if !basicStruct.ContainsItem(res, nums[i]){
			res = append(res, nums[i])
			permuteBasic(nums, res, result)
			res = res[:len(res)-1]
		}
	}
}

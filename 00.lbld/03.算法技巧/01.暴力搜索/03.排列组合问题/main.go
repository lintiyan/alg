package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
	"sort"
)

func main() {

	// 元素无重不可复选

	// 01. 给一个无重复元素的数组nums，每个元素最多使用一次，请返回nums所有的子集
	fmt.Println(SubSets([]int{1, 2, 3}))

	// 02. 给定两个整数n 和 K，返回范围[1,n]中，所有可能的K个数的组合
	fmt.Println(Combine(3, 2))

	// 03. 给定一个不重复数组的数组，返回其所有可能的全排列
	fmt.Println(permute([]int{1, 2, 3}))

	// 元素可重不可复选
	// 04. 给你一个整数数组，其中可能包含重复元素，请你返回改数组所有可能的子集
	fmt.Println(SubSetsWithDup([]int{1, 2, 2}))

	// 05. 给你一个可包含重复数字的序列，返回所有可能的全排列
	fmt.Println(permuteUnique([]int{1, 2, 2}))
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
	var used = make([]bool, len(nums)) // 判断元素是否已使用
	permuteBasic(nums, []int{}, used, result)
	return result.ResAll
}
func permuteBasic(nums []int, res []int, used []bool, result *Result) {

	if len(res) == len(nums) {
		result.ResAll = append(result.ResAll, basicStruct.CopyRes(res))
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			res = append(res, nums[i])
			used[i] = true
			permuteBasic(nums, res, used, result)
			res = res[:len(res)-1]
			used[i] = false
		}
	}
}

func SubSetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var result = &Result{}
	SubSetsWithDupBasic(nums, 0, []int{}, result)
	return result.ResAll
}
func SubSetsWithDupBasic(nums []int, start int, res []int, result *Result) {

	result.ResAll = append(result.ResAll, basicStruct.CopyRes(res))

	for i := start; i < len(nums); i++ {
		// i > start 原因是 避免数组越界，还有是对于某个节点下得分支进行剪枝 [1,2,2] 这种是允许的
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		res = append(res, nums[i])
		SubSetsWithDupBasic(nums, i+1, res, result)
		res = res[:len(res)-1]
	}
}

func permuteUnique(nums []int) [][]int {
	var result = &Result{}
	used := make([]bool, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	permuteUniqueBasic(nums, []int{}, used, result)
	return result.ResAll
}
func permuteUniqueBasic(nums []int, res []int, used []bool, result *Result) {

	if len(res) == len(nums) {
		result.ResAll = append(result.ResAll, basicStruct.CopyRes(res))
		return
	}

	for i := 0; i < len(nums); i++ {
		// 已使用过当然得过滤掉
		// 如果nums[i]=nums[i-1]，则表示分支有重复，那么如果nums[i-1]没有被使用得话，nums[i]也不能被使用，这样去保证相同元素的相对位置
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		used[i] = true
		res = append(res, nums[i])
		permuteUniqueBasic(nums, res, used, result)
		res = res[:len(res)-1]
		used[i] = false
	}
}

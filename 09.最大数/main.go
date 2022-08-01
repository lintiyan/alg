package main

import (
	"fmt"
	"sort"
)

/*
	给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

	注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。


	示例 1：

	输入：nums = [10,2]
	输出："210"
	示例2：

	输入：nums = [3,30,34,5,9]
	输出："9534330"
*/

func main() {
	nums := []int{0,0,0,0,0}
	fmt.Println(largestNumber(nums))
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return fmt.Sprintf("%d%d", nums[i], nums[j]) > fmt.Sprintf("%d%d", nums[j], nums[i])
	})

	if nums[0] == 0 {
		return "0"
	}

	var res string
	for _, item := range nums{
		res += fmt.Sprintf("%d", item)
	}

	return res
}
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int)[][]int{

	var res = make([][]int,0)
	if len(nums) < 3 {
		return res
	}

	// 排序
	sort.Ints(nums)
	fmt.Println(nums)
	for i:=0;i<len(nums)-2;i++{
		if i>0&& nums[i] == nums[i-1]{
			continue
		}
		left := i+1
		right := len(nums)-1

		for left < right {
			if nums[i] + nums[left] + nums[right] < 0 {
				left++
			}else if nums[i] + nums[left] + nums[right] > 0 {
				right--
			}else {
				if right < len(nums)-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				resItem := []int{nums[i], nums[left], nums[right]}
				fmt.Println("i=",i,"left=",left, "right=", right)
				fmt.Println(resItem)
				res = append(res, resItem)
				right--
			}
		}
	}

	return res
}



func main(){
	res := threeSum([]int{-9,-9,-9,-8,-6,-4,-2,-1,-1,-1,1,1,1,1,2,3,5,6,7,8,9,10,10,10})
	fmt.Println(res)
}

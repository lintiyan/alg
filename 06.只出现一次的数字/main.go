package main

import "fmt"

func main() {
	var arr = []int {1,2,3,2,3,4,5,6,4,5,6,1,8}
	fmt.Println(singleNumber(arr))
}


func singleNumber(nums []int) int {

	var res int
	for _, item := range nums {
		res ^= item
	}

	return res
}
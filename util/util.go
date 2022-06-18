package util

import "fmt"

// 遍历map
func TraverseMap(target map[int]int){
	for k, v:= range target{
		fmt.Println("key = ",k, "value=",v)
	}
}

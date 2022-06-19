package _0_基础排序算法

import "testing"

func TestHeapSort(t *testing.T){
	arr := []int{10,2,8,18,27,36,24}

	HeapSort(arr)
	t.Log(arr)
}


func TestQuickSort(t *testing.T){
	arr := []int{10,2,8,18,27,36,24}
	QuickSort(arr)
	t.Log(arr)
}
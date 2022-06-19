package _0_基础排序算法

import "fmt"

func QuickSort(sli []int){
	left := 0
	right := len(sli) -1

	QuickSort0(sli, left, right)
}

func QuickSort0(sli []int, left, right int){

	if left >= right {
		return
	}

	var i, j = left, right
	var t = sli[left]
	for left < right {
		// 右边，碰到比t小的则 交换位置，并停止
		for left < right {

			if sli[right] < t {
				sli[left] = sli[right]
				left++
				break
			}else{
				right--
			}
		}

		// 左边，碰到比t大的则 交换位置，并停止
		for left < right {
			if sli[left] > t {
				sli[right] = sli[left]
				right--
				break
			}else {
				left++
			}
		}
	}

	// left 有没有可能比right大呢？
	if left == right {
		sli[left] = t
	}
	fmt.Println(sli, i, left -1 , right + 1, j)
	QuickSort0(sli, i, left-1)
	QuickSort0(sli, right+1, j)
}

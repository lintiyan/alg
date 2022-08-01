package main

import "fmt"

/*

	给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

	示例 1：

	输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
	输出：3
	解释：长度最长的公共子数组是 [3,2,1] 。
	示例 2：

	输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
	输出：5

*/

func findLength(nums1 []int, nums2 []int) int {

	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	var row = len(nums1)
	var col = len(nums2)
	var matrix [][]int

	for i := 0; i < row; i++ {
		var item = make([]int, col)
		matrix = append(matrix, item)
	}

	// matrix[i][j] 表示以nums1的第i个元素,nums2的第j个元素结尾的最长重复子数组的长度
	var max int
	for i := 0; i < col; i++ {
		if nums2[i] == nums1[0] {
			matrix[0][i] = 1
		}
		if matrix[0][i] > max {
			max = matrix[0][i]
		}
	}

	for i := 0; i < row; i++ {
		if nums1[i] == nums2[0] {
			matrix[i][0] = 1
		}
		if matrix[i][0] > max {
			max = matrix[i][0]
		}
	}



	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if nums1[i] == nums2[j] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			}else {
				matrix[i][j] = 0
			}

			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
	}

	return max
}

func main() {

	//var nums1 = []int{1,2,3,2,1,5}
	//var nums2 = []int{3,2,1,5,4,7}
	var nums1 = []int{1,2,3,2,8}
	var nums2 = []int{5,6,1,4,7}
	fmt.Println(findLength(nums1, nums2))
}

package main

import "fmt"

/*

	在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

	 

	示例 1：

	输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
	输出：4

	示例 2：

	输入：matrix = [["0","1"],["1","0"]]
	输出：1

	示例 3：

	输入：matrix = [["0"]]
	输出：0
	 

	提示：

	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 300
	matrix[i][j] 为 '0' 或 '1'

*/

func maximalSquare(matrix [][]byte) int {

	var row = len(matrix)
	var col = len(matrix[0])

	// dp[i][j]代表以matrix[i][j]为右下角的最大正方形的边长
	var dp [][]int

	for i := 0; i < row; i++ {
		var item = make([]int, col)
		dp = append(dp, item)
	}
	var maxLength int

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			temp := matrix[i][j] - '0'
			if temp == 1 {
				if i == 0 || j == 0 {
					// 靠边的正方形最大变成为1
					dp[i][j] = 1
				} else {
					// 以dp[i][j]为右下角的最大正方形的边长 = 以dp[i-1][j-1]为右下角的最大正方形边长 + 1
					// 或者是以 dp[i-1][j]为右下角的最大正方形边长 + 1
					// 或者是以 dp[i][j-1]为右下角的最大正方形边长 + 1
					// 取三者最小值
					dp[i][j] = Min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				}
				if maxLength < dp[i][j] {
					maxLength = dp[i][j]
				}
			}
		}
	}
	return maxLength * maxLength
}

func Min(a, b, c int) int {
	var min = a
	if a > b {
		min = b
	}

	if min > c {
		min = c
	}
	return min
}

func main() {

	var matrix = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	fmt.Println(maximalSquare(matrix))
}

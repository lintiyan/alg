package main

import "fmt"

func main() {

	// 01. N皇后问题
	NQueue(8)
	fmt.Println(res)
	fmt.Println(len(res))
}

func NQueue(n int) {

	var board [][]int
	for i := 0; i < n; i++ {
		var item = make([]int, n)
		board = append(board, item)
	}

	NQueueBasic(0, board)
}

var res [][][]int

func NQueueBasic(i int, board [][]int) {

	if i >= len(board) {
		res = append(res, copyTwoDimensionSlice(board))
		return
	}

	// 有n中选择，每种都进行遍历
	for j := 0; j < len(board); j++ {
		// 检测选择是否有效
		if !isValid(i, j, board) {
			continue
		}
		// 记录选择
		board[i][j] = 1
		// 选择后，递归下一条
		NQueueBasic(i+1, board)
		// 取消选择
		board[i][j] = 0
	}
}

func isValid(i, j int, board [][]int) bool {

	// 同一列
	for k := 0; k < i; k++ {
		if board[k][j] == 1 {
			return false
		}
	}
	// 同一行
	for k := 0; k < j; k++ {
		if board[i][k] == 1 {
			return false
		}
	}

	// 左上
	for k, m := i-1, j-1; k >= 0 && m >= 0; {
		if board[k][m] == 1 {
			return false
		}
		k--
		m--
	}

	// 右上
	for k, m := i-1, j+1; k >= 0 && m < len(board[0]); {
		if board[k][m] == 1 {
			return false
		}
		k--
		m++
	}
	return true
}

func copyTwoDimensionSlice(src [][]int) [][]int {
	var dst [][]int
	for i := 0; i < len(src); i++ {
		var item = make([]int, len(src))
		dst = append(dst, item)
	}

	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src[0]); j++ {
			dst[i][j] = src[i][j]
		}
	}
	return dst
}

package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
	"math"
)

func main() {

	// 01. 最小路径和。给你一个二维数组，其中的元素都是非负整数，现在你站在左上角，只能向右或者向下移动，需要到达右下角。
	// 现在请你计算，经过的路径和最小是多少？
	var grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSumTop(grid,0,0))
	fmt.Println(minPathSumFromBottom(grid))
}

func minPathSumTop(grid [][]int, i, j int) int {
	return minPathSumFromTopBasic(grid, i, j)
}

func minPathSumFromTopBasic(grid [][]int, i, j int) int {

	if i >= len(grid) {
		return math.MaxInt
	}

	if j >= len(grid[0]) {
		return math.MaxInt
	}

	// base case
	if i == len(grid)-1 && j == len(grid[0])-1 {
		return grid[i][j]
	}

	return basicStruct.GetMin(minPathSumFromTopBasic(grid, i+1, j), minPathSumFromTopBasic(grid, i, j+1)) + grid[i][j]
}

func minPathSumFromBottom(grid [][]int) int {

	// dp[i][j]的含义是，从左上角位置走到dp[i][j]的最短路径
	var dp [][]int
	for i := 0; i < len(grid); i++ {
		item := make([]int, len(grid[0]))
		dp = append(dp, item)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = basicStruct.GetMin(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

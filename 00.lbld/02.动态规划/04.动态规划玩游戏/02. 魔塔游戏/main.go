package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
	"math"
)

func main() {

	// 01. 魔塔。输入一个存储整数的二维数组grid，
	// 如果grid[i][j]>0, 说明这个格子存放的是血量
	// 如果grid[i][j]=0, 说明是个空格子
	// 如果grid[i][j]<0, 说明这个格子存放的是妖怪，会掉血
	// 如果你是勇士，将会出现在左上角，公主被困在右下角，你只能向右或者向左移动，请问勇士初始生命值至少为多少时，才能成功拯救公主？
	var grid = [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	fmt.Println(calculateMinHP(grid))
}

func calculateMinHP(gird [][]int) int {

	return calculateMinHPFromTop(gird, 0, 0)
}

// 自顶向下的思路
func calculateMinHPFromTop(grid [][]int, i, j int) int {

	if i >= len(grid) {
		return math.MaxInt
	}
	if j >= len(grid[0]) {
		return math.MaxInt
	}

	if i == len(grid)-1 && j == len(grid[0])-1 {
		if grid[i][j] >= 0 {
			return 1
		} else {
			return 0 - grid[i][j] + 1
		}
	}

	min := basicStruct.GetMin(calculateMinHPFromTop(grid, i+1, j), calculateMinHPFromTop(grid, i, j+1))
	res := grid[i][j] - min
	if res >= 0 {
		return 0
	} else {
		return 0 - res
	}
}

// 自底向上
func calculateMinHPFromBottom(grid [][]int) int {
	return 0
}

package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

/*

	编辑距离。 给定两个字符串s1和s2,计算将s1转换成s2所使用的最少操作数。你可以，插入、删除、替换一个字符

*/

func main() {

	// 01. 递归解法--自顶向下
	fmt.Println(miniDistanceRecursion("horse", "ros"))
	fmt.Println(miniDistanceRecursion("intention", "execution"))

	// 02. 带备忘录的递归解法--自顶向下
	//fmt.Println(miniDistanceRecursionMemo("horse", "ros"))
	fmt.Println(miniDistanceRecursionMemo("intention", "execution"))

	// 03. dp解法--自底向上
	fmt.Println(miniDistanceDP("horse", "ros"))
	fmt.Println(miniDistanceDP("intention", "execution"))
}

func miniDistanceRecursion(s1, s2 string) int {
	return miniDistanceRecursionBasic(s1, 0, s2, 0)
}

func miniDistanceRecursionBasic(s1 string, i int, s2 string, j int) int {

	if i == len(s1) {
		return len(s2) - j
	}
	if j == len(s2) {
		return len(s1) - i
	}

	if s1[i] == s2[j] {
		return miniDistanceRecursionBasic(s1, i+1, s2, j+1) // s1[i] == s2[j]，直接跳过计算s1[i+1:]到s2[j+1:]
	} else {
		d1 := miniDistanceRecursionBasic(s1, i+1, s2, j) + 1   // 得到s1[i+1:] => s2[j:]的编辑距离后，再加一次删除操作
		d2 := miniDistanceRecursionBasic(s1, i, s2, j+1) + 1   // 得到s1[i:]=>s2[j+1:]的编辑距离后，再加一次插入操作
		d3 := miniDistanceRecursionBasic(s1, i+1, s2, j+1) + 1 // 得到s1[i+1:]=>s2[j+1:]的编辑距离后，再加一次替换操作
		return basicStruct.GetMin(basicStruct.GetMin(d1, d2), d3)
	}
}

func miniDistanceRecursionMemo(s1, s2 string) int {
	return miniDistanceRecursionMemoBasic(s1, 0, s2, 0)
}

var memo = make(map[string]int)

func miniDistanceRecursionMemoBasic(s1 string, i int, s2 string, j int) int {

	if i == len(s1) {
		return len(s2) - j
	}
	if j == len(s2) {
		return len(s1) - i
	}

	if s1[i] == s2[j] {
		if val, ok := memo[fmt.Sprintf("%d%d", i, j)]; ok {
			return val
		} else {
			temp := miniDistanceRecursionBasic(s1, i+1, s2, j+1)
			memo[fmt.Sprintf("%d%d", i, j)] = temp
			return temp
		}

	} else {
		var d1, d2, d3 int
		if val1, ok := memo[fmt.Sprintf("%d%d", i+1, j)]; ok {
			d1 = val1
		} else {
			temp := miniDistanceRecursionBasic(s1, i+1, s2, j) + 1
			memo[fmt.Sprintf("%d%d", i+1, j)] = temp
			d1 = temp
		}

		if val2, ok := memo[fmt.Sprintf("%d%d", i, j+1)]; ok {
			d2 = val2
		} else {
			temp := miniDistanceRecursionBasic(s1, i, s2, j+1) + 1
			memo[fmt.Sprintf("%d%d", i, j+1)] = temp
			d2 = temp
		}

		if val3, ok := memo[fmt.Sprintf("%d%d", i+1, j+1)]; ok {
			d3 = val3
		} else {
			temp := miniDistanceRecursionBasic(s1, i+1, s2, j+1) + 1
			memo[fmt.Sprintf("%d%d", i+1, j+1)] = temp
			d3 = temp
		}

		return basicStruct.GetMin(basicStruct.GetMin(d1, d2), d3)
	}
}

func miniDistanceDP(s1, s2 string) int {

	var dp [][]int
	for i := 0; i < len(s1); i++ {
		var item = make([]int, len(s2))
		dp = append(dp, item)
	}

	// i=0
	var i, j = 0, 0
	for j = 0; j < len(s2); j++ {
		if s1[i] == s2[j] {
			dp[i][j] = j
		} else {
			dp[i][j] = j + 1
		}
	}
	i, j = 0, 0
	for i = 0; i < len(s1); i++ {
		if s1[i] == s2[j] {
			dp[i][j] = i
		} else {
			dp[i][j] = i + 1
		}
	}

	for i = 1; i < len(s1); i++ {
		for j = 1; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = basicStruct.GetMin(basicStruct.GetMin(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len(s1)-1][len(s2)-1]
}

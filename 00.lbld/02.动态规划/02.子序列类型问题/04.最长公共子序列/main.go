package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	// 01. 最长公共子序列。给你两个字符串，找出他们最长公共子序列，返回这个子序列的长度
	fmt.Println(longestCommonSubsequence("zabcde", "acez"))
	fmt.Println(dpLCSFromBottom("zabcde", "acez"))

	// 02.字符串的删除操作。给定两个单词s1,s2 找到使得s1和s2相同所需的最小步数，每步可以删除任意一个字符串中的一个字符
	fmt.Println(minDistance("seater", "eat"))

	// 03. 两个字符串的最小ASCII删除和

}

func longestCommonSubsequence(s1, s2 string) int {
	return dpLCSFromTop(s1, 0, s2, 0)
}

// 自顶向下
var memo = make(map[string]int)

func dpLCSFromTop(s1 string, i int, s2 string, j int) int {

	// base case
	if i == len(s1) || j == len(s2) {
		return 0
	}

	if s1[i] == s2[j] {
		var val int
		var ok bool
		if val, ok = memo[fmt.Sprintf("%d%d", i+1, j+1)]; !ok {
			val = dpLCSFromTop(s1, i+1, s2, j+1)
			memo[fmt.Sprintf("%d%d", i, j)] = val
		}
		return val + 1
	} else {

		var val1, val2, val3 int

		if val, ok := memo[fmt.Sprintf("%d%d", i, j+1)]; !ok {
			val1 = dpLCSFromTop(s1, i, s2, j+1)
			memo[fmt.Sprintf("%d%d", i, j+1)] = val1
		} else {
			val1 = val
		}

		if val, ok := memo[fmt.Sprintf("%d%d", i+1, j)]; !ok {
			val2 = dpLCSFromTop(s1, i+1, s2, j)
			memo[fmt.Sprintf("%d%d", i+1, j)] = val2
		} else {
			val2 = val
		}

		if val, ok := memo[fmt.Sprintf("%d%d", i+1, j+1)]; !ok {
			val3 = dpLCSFromTop(s1, i+1, s2, j+1)
			memo[fmt.Sprintf("%d%d", i+1, j+1)] = val3
		} else {
			val3 = val
		}

		return basicStruct.GetMax(basicStruct.GetMax(val1, val2), val3)
	}
}

// 自底向上
func dpLCSFromBottom(s1 string, s2 string) int {
	// 建立dp数组，含义是 以s1[i]， s2[j] 结尾的字符串最长公共子序列的长度
	var dp [][]int
	for i := 0; i < len(s1); i++ {
		var item = make([]int, len(s2))
		dp = append(dp, item)
	}

	// base case，第一行和第一列
	for j := 0; j < len(s2); j++ {
		if s1[0] == s2[j] {
			dp[0][j] = 1
		}
		if j > 0 && dp[0][j-1] == 1 {
			dp[0][j] = 1
		}
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[0] {
			dp[i][0] = 1
		}
		if i > 0 && dp[i-1][0] == 1 {
			dp[i][0] = 1
		}
	}

	var res int

	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = basicStruct.GetMax(dp[i-1][j], basicStruct.GetMax(dp[i][j-1], dp[i-1][j-1]))
			}

			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

func minDistance(s1, s2 string) int {
	return dpMDFromTop(s1, 0, s2, 0)
}

// 自顶向下
func dpMDFromTop(s1 string, i int, s2 string, j int) int {
	if i == len(s1) {
		return len(s2) - j
	}
	if j == len(s2) {
		return len(s1) - i
	}

	if s1[i] == s2[j] {
		return dpMDFromTop(s1, i+1, s2, j+1)
	} else {
		return basicStruct.GetMin(dpMDFromTop(s1, i+1, s2, j), basicStruct.GetMin(dpMDFromTop(s1, i, s2, j+1), dpMDFromTop(s1, i+1, s2, j+1))) + 1
	}
}

func minimumDeleteSum(s1, s2 string) int {
	return dpMinimumDeleteSum(s1, 0, s2, 0)
}

func dpMinimumDeleteSum(s1 string, i int, s2 string, j int) int {
	// 和最小删除数一样，不过求值修改下
	return 0
}

package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

// 子序列问题一般都是两个模板，要么是一维数组，要么是二维数组，往题目上套就行，套不出来就换，怎样好推算状态转移方程就用哪个
// 一般一个字符串或者一个数组的dp问题都是用一维，两个字符串的问题用二维，也有部分一个字符串的用二维，比如今天的两道题

func main() {
	// 01. 最长回文子序列。 返回一个最长回文子序列的长度
	fmt.Println(longestPalindromeSubSeq("aeccda"))
	fmt.Println(longestSubSeqFromBottom("aeccda"))

	// 02. 任意位置插入任意字符使得字符串称为回文串的最小插入次数。
	fmt.Println(minInsertions("acba"))
}

func longestPalindromeSubSeq(s string) int {
	return longestSubSeqFromTop(s, 0, len(s)-1)
}

// 自顶向下的解法
func longestSubSeqFromTop(s string, i, j int) int {

	if i == j {
		return 1
	}
	if i > j {
		return 0
	}

	if s[i] == s[j] {
		return longestSubSeqFromTop(s, i+1, j-1) + 2
	} else {
		return basicStruct.GetMax(longestSubSeqFromTop(s, i+1, j-1), basicStruct.GetMax(longestSubSeqFromTop(s, i+1, j), longestSubSeqFromTop(s, i, j-1)))
	}
}

// 自底向上的解法
func longestSubSeqFromBottom(s string) int {

	var dp [][]int
	for i := 0; i < len(s); i++ {
		var item = make([]int, len(s))
		dp = append(dp, item)
	}

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= 0; j-- {
			if i == j {
				dp[i][j] = 1
			}
		}
	}
	fmt.Println(dp)

	var res int
	//[
	//[1 1 1 1 1 2]
	//[0 1 1 1 1 1]
	//[0 0 1 2 2 2]
	//[0 0 0 1 1 1]
	//[0 0 0 0 1 1]
	//[0 0 0 0 0 1]
	//]
	// aeccda
	for j := 1; j < len(s); j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = basicStruct.GetMax(dp[i+1][j-1], basicStruct.GetMax(dp[i+1][j], dp[i][j-1]))
			}

			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	fmt.Println(dp)
	return res
}

func minInsertions(s string) int {
	return minInsertionsFromTop(s, 0, len(s)-1)
}
func minInsertionsFromTop(s string, i, j int) int {
	if i == j {
		return 0
	}

	if i > j {
		return 0
	}

	if s[i] == s[j] {
		return minInsertionsFromTop(s, i+1, j-1)
	} else {
		return basicStruct.GetMin(minInsertionsFromTop(s, i, j-1), minInsertionsFromTop(s, i+1, j)) + 1
	}
}

package main

import (
	"fmt"
)

func main() {

	// 01. 给定不同面额的硬币和一个总金额，写出函数来计算可以凑成总金额的硬币组合数。假设每种硬币有无限个
	fmt.Println(getChangeNum([]int{1, 2, 5}, 15))
	fmt.Println(getChangeNumFromBottom([]int{1, 2, 5}, 15))
}

func getChangeNum(coins []int, amount int) int {
	//sort.Slice(coins, func(i, j int) bool {
	//	return coins[i] > coins[j]
	//})

	return getChangeNumFromTop(coins, amount, 0)
}

// 自顶向下
func getChangeNumFromTop(coins []int, amount int, i int) int {

	// 如果amount为0，则表示找到了一种方式能组合成原始amount
	if amount == 0 {
		return 1
	}
	// 如果i到达数组末尾还没找到，则返回0
	if i >= len(coins) {
		return 0
	}

	var res int

	// 分别计算取0-j个coins[i]时，能组合成amount的数量，求和即得到结果
	for j := 0; j <= amount/coins[i]; j++ {
		res += getChangeNumFromTop(coins, amount-j*coins[i], i+1)
	}

	return res
}

// 自底向上
func getChangeNumFromBottom(coins []int, amount int) int {

	// dp[i][a]的含义是，使用coins[...i], 能拼出amount的总数，那么我们要求的就是dp[len(nums)-1][amount]
	var dp [][]int
	for i := 0; i < len(coins); i++ {
		item := make([]int, amount+1)
		dp = append(dp, item)
	}

	// 当amount=0时，作为base case，此时dp[i][0]应该为1。因为当其他情况向下寻找时，如果amount=0表示能找到一种方法能拼出amount
	for i := 0; i < len(coins); i++ {
		dp[i][0] = 1
	}
	// 当只有一个币种时，那么能对这个币值取余为0的金额能被拼出来
	for a := 0; a <= amount; a++ {
		if a%coins[0] == 0 {
			dp[0][a] = 1
		}
	}

	for i := 1; i < len(coins); i++ {
		for a := 1; a <= amount; a++ {
			// 分别计算，使用0到j个coins[i]时，使用coins[...i-1]能拼出余额的方法数
			for j := 0; j <= a/coins[i]; j++ {
				dp[i][a] += dp[i-1][a-j*coins[i]]
			}
		}
	}

	return dp[len(coins)-1][amount]
}

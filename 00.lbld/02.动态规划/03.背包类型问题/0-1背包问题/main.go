package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	// 01. 给你一个可装载重量为W的背包和N个物品，每个物品有重量和价值两个属性。
	// 其中，第i个物品的重量为wt[i]，价值为val[i]，现在让你用这个背包装物品，最多能装的价值是多少？
	fmt.Println(knapsack(4, 3, []int{2, 1, 3}, []int{4, 2, 3}))
	fmt.Println(knapsackFromBottom(4, 3, []int{2, 1, 3}, []int{4, 2, 3}))
}

func knapsack(W int, N int, wt []int, val []int) int {
	return knapsackFromTop(W, 0, N, wt, val)
}

// 自顶向下的动态规划问题，都可以用备忘录进行优化，这里就不写了
func knapsackFromTop(w int, i int, N int, wt []int, val []int) int {

	// 背包再无空余
	if w <= 0 {
		return 0
	}

	// 再无物品可装
	if i >= N {
		return 0
	}

	var res1, res2 int

	// 只有能装下第i个的时候才有选择权，否则只能选择不装第i个
	if w-wt[i] >= 0 {
		pick := knapsackFromTop(w-wt[i], i+1, N, wt, val) + val[i]
		notPick := knapsackFromTop(w, i+1, N, wt, val)
		res1 = basicStruct.GetMax(pick, notPick)
	} else {
		res2 = knapsackFromTop(w, i+1, N, wt, val)
	}

	return basicStruct.GetMax(res1, res2)
}

// 自底向上的dp
func knapsackFromBottom(W int, N int, wt []int, val []int) int {

	var dp [][]int
	for i := 0; i < N; i++ {
		item := make([]int, W+1)
		dp = append(dp, item)
	}

	for w := 0; w <= W; w++ {
		if w >= wt[0] {
			dp[0][w] = val[0]
		}
	}

	var res int


	for i := 1; i < N; i++ {
		for w := 0; w <= W; w++ {
			var target1, target2 int
			if w >= wt[i] {
				target1 = basicStruct.GetMax(dp[i-1][w-wt[i]]+val[i], dp[i-1][w])
			} else {
				target2 = dp[i-1][w]
			}

			dp[i][w] = basicStruct.GetMax(target1, target2)

			if dp[i][w] > res {
				res = dp[i][w]
			}
		}
	}

	return res
}

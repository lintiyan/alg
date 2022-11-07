package main

import "fmt"

func main() {

	// 01. 你和你的朋友面前有一排石头，用一个数组piles表示，piles[i]表示第i堆石子有多少个。
	// 你们轮流拿走最左边或者最右边的石头堆。所有石头被拿完后，谁拥有的石头多，谁获胜。
	// 假设两个人都很聪明，请你设计一个算法，返回先手和后手的最后得分之差。
	fmt.Println(StoneGame([]int{1, 2, 3, 100, 6, 3}))
	fmt.Println(stoneGameFromBottom([]int{1, 2, 3, 100, 6, 3}))
}

func StoneGame(piles []int) int {
	first, second := stoneGameFromTop(piles, 0, len(piles)-1)
	fmt.Println(first, second)
	return first - second
}

// 自顶向下。 这里没有用备忘录去重，简单加个map就行，就不写了
// 这个递归函数的含义是，对于piles[i...j]，返回先手和后手分别能取得的最大石头数
func stoneGameFromTop(piles []int, i, j int) (int, int) {

	if i == j {
		return piles[i], 0
	}
	var first int
	var second int

	// 假如先手选择左边，那剩余的石头堆里，先手会变成后手。相应的后手会变成先手
	// 只要计算出先手选择左边piles[i]后，与在剩下的石堆里后手能拿到的最大值leftSecond 之和
	// 与先手选择右边piles[j]后，与剩下的石堆后手能拿到的最大值rightSecond 之和 相比较
	// 就可以得出对于piles[i...j]先手能拿到的最大石头数
	// 对于piles[i...j]后手能拿到的最大石头数依赖于先手的选择
	// 如果先手选择左边，那么后手能拿到的最大值就是对于piles[i+1...j]先手能拿到的最大值
	// 如果先手选择右边，那么后手能拿到的最大值就是对于piles[i...j-1]先手能拿到的最大值

	// 对于piles[i+1...j] 先手与后手能拿到的最大值
	leftFirst, leftSecond := stoneGameFromTop(piles, i+1, j)

	// 对于piles[i...j-1] 先手与后手能拿到的最大值
	rightFirst, rightSecond := stoneGameFromTop(piles, i, j-1)

	// 如果先手选择左边 比先手选择右边获得石头多, 那先手必然选择左边，后手只能取piles[i+1...j]中先手的最大值
	if leftSecond+piles[i] > rightSecond+piles[j] {
		first = leftSecond + piles[i]
		second = leftFirst
	} else {
		first = rightSecond + piles[j]
		second = rightFirst
	}

	return first, second
}

// 自底向上
func stoneGameFromBottom(piles []int) int {

	var dp [][][]int
	for i := 0; i < len(piles); i++ {

		item := make([][]int, len(piles))
		for j := range item {
			item[j] = make([]int, 2)
		}
		dp = append(dp, item)
	}

	// base case. 当i==j时，先手总是最优，后手总是吃亏为0
	for i := 0; i < len(piles); i++ {
		for j := 0; j < len(piles); j++ {
			if i == j {
				dp[i][j][0] = piles[i]
			}
		}
	}

	for j := 1; j < len(piles); j++ {
		for i := j - 1; i >= 0; i-- {

			// 先手选左边
			leftFirst := dp[i+1][j][0]
			leftSecond := dp[i+1][j][1]

			// 先手选右边
			rightFirst := dp[i][j-1][0]
			rightSecond := dp[i][j-1][1]

			if piles[i] + leftSecond > piles[j] + rightSecond {
				dp[i][j][0] = piles[i] + leftSecond
				dp[i][j][1] = leftFirst
			} else {
				dp[i][j][0] = piles[j] + rightSecond
				dp[i][j][1] = rightFirst
			}
		}
	}

	return dp[0][len(piles)-1][0] - dp[0][len(piles)-1][1]
}

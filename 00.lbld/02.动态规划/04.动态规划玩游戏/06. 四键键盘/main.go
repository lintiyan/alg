package main

import "fmt"

func main() {
	//  假设你有一个特殊的键盘，包含下面的按键：
	// key1 - A : 在屏幕上打印出一个 'A'
	// key2 - Ctrl + A : 选中整个屏幕
	// key3 - Ctrl + C : 复制选中区域到缓冲区
	// key4 - Ctrl + V : 将缓冲区的内容输出到上次输入的结束位置，并显示在屏幕上
	// 现在你只可以按键N次，请问屏幕上最多可以显示多少个A？
	fmt.Println(maxAFromTop(6))
	fmt.Println(maxAFromBottom(7))
}

// 自顶向下的解法，优点像二叉树的后续遍历哈
func maxAFromTop(n int) (int, int) {
	if n == 0 {
		return 0, 0
	}

	if n == 1 || n == 2 || n == 3 {
		return n, 0
	}

	// 当只有n次机会时，推理一下将问题拆成子问题的思路
	// case1:如果得到n-1的结果，那么最后一次机会选择打印一个A即可得到n的结果
	// case2:如果得到n-1的结果，那么最后一次机会选择粘贴一次缓冲区的内容，即可得到n的结果，当然缓冲区有可能为空
	// case3:如果得到n-3的结果，那么最后三次机会选择CA、CC、CV即可得到n的结果
	// 取三者最大值即可。
	// 对于case2，因为不知道缓冲区中A的个数，所以无法进行值得比较，所以必须知道n-1时，缓冲区得大小
	// 也就是子问题必须返回当前缓冲区的大小。
	// 考虑下，当选择最后一次操作是打印A或者执行CV操作时，缓冲区的大小是一样的
	// 当最后选择执行CA、CC、CV操作时，缓冲区的大小为n-3的值。
	// 比较各种选择后取最大的那个选择，并返回对应的缓冲区大小就可以
	// 这种解法有点后序遍历那味道了，因为大问题依赖小问题的解。
	// 最后，肯定也可以用map来减少计算，这里就不写了

	// 假设最后一步选择打印一个A
	printA, buf := maxAFromTop(n - 1)
	res1 := printA + 1
	// 假设最后一步选择执行CV
	res2 := printA + buf
	// 假设最后一步的结果为n-3的值 再执行CA\CC\CV，此时缓冲区的A个数为cv
	cv, _ := maxAFromTop(n - 3)
	res3 := cv * 2

	if res1 > res2 {
		if res1 > res3 {
			// 最后一步选择打印A最优，此时缓冲区中A的个数为buf
			return res1, buf
		} else {
			// 最后一步的结果是n-3的结果再执行复制粘贴，此时缓冲区中A的个数为n-3的结果数
			return res3, cv
		}
	} else {
		if res2 > res3 {
			// 最后一步执行粘贴最优，此时缓冲区中A的个数为buf
			return res2, buf
		} else {
			// 最后一步的结果是n-3的结果再执行复制粘贴，此时缓冲区中A的个数为n-3的结果数
			return res3, cv
		}
	}
}

// 自底向上
func maxAFromBottom(n int) int {
	// dp[i]的含义是 只有i次机会时，能打印的A的最大个数
	var dp = make([]int, n+1)
	// 初始化
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	var buf int
	for i := 4; i <= n; i++ {
		res1 := dp[i-1] + 1
		res2 := dp[i-1] + buf
		res3 := dp[i-3] * 2

		if res1 > res2 {
			if res1 > res3 {
				dp[i] = res1
			} else {
				dp[i] = res3
				buf = dp[i-3]
			}
		} else {
			if res2 > res3 {
				dp[i] = res2
			} else {
				dp[i] = res3
				buf = dp[i-3]
			}
		}
	}

	return dp[n]
}

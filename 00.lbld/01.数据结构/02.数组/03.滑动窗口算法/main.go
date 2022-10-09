package main

import "fmt"

func main() {

	// 01. 最小覆盖子串，给你一个字符串s，一个字符串t，返回s中涵盖t的所有字符的最小子串。不存在返回""
	//midstSubStringTest()

	// 02. 字符串的排列。给你两个字符串s1,s2 写一个函数来判断s2是否包含s1的排列
	checkInclusionTest()
}

func midstSubStringTest() {
	fmt.Println(midstSubString("abccdeefgh", "ccdf"))
}

// 滑动窗口问题需要考虑：
// 1. 窗口中包含目标字符的map 什么时候更新？ 在窗口扩大及缩小的时候需要更新
// 2. 窗口什么时候缩小？ 在窗口中已经包含了所有目标字符的时候需要缩小以达到最小子字符串的目的
// 3. 窗口缩小时有什么需要更新？ 窗口缩小时，当前最小覆盖子串的结果需要更新，窗口中所含目标字符个数也需要更新
func midstSubString(src string, target string) string {

	left, right := 0, 0 // 左右边界

	// 包含所有目标字符的map
	var neededMap = make(map[byte]int, len(target))
	for i := range target {
		neededMap[target[i]] += 1
	}

	// 左右边界形成的窗口的子字符串，包含目标字符的信息，key为目标字符，val为窗口中目标字符的个数
	var windowMap = make(map[byte]int, len(target))
	var valid int // 窗口中达到目标字符串个数要求的字符数

	var res string        // 结果字符串
	var minLen = len(src) // 最小覆盖子串的长度

	// 右边界到达末尾遍历结束
	for right < len(src) {

		//当前右边界上的字符是目标字符，则更新窗口信息
		if _, ok := neededMap[src[right]]; ok {
			windowMap[src[right]] += 1
			if windowMap[src[right]] == neededMap[src[right]] {
				valid++
			}
		}
		right++
		fmt.Println("left=", left, "right=", right, "window_map=", windowMap, "valid=", valid)
		// 窗口中包含了所有目标字符，则可以尝试缩小窗口，以达到最小的目的
		for valid == len(neededMap) {

			if right-left < minLen {
				minLen = right - left
				res = src[left:right]
			}
			fmt.Println("left=", left, "right=", right, "res=", res)
			// 因为左边界需要右移，在移动前将窗口信息更新，也可放在左边界移动之后更新
			// 如果左边界上的字符是目标字符，且窗口左边界右移后窗口中该目标字符个数正好等于目标字符串中目标字符的个数，则valid--
			if count, ok := neededMap[src[left]]; ok && windowMap[src[left]] == count {
				windowMap[src[left]] -= 1
				valid--
			}
			left++
		}
	}

	return res
}


func checkInclusionTest(){

	fmt.Println(checkInclusion("acbad", "acd"))
}
func checkInclusion(src string, target string) bool {
	var left, right = 0, 0

	var neededMap = make(map[byte]int, len(target))
	for i := 0; i < len(target); i++ {
		neededMap[target[i]] += 1
	}

	var windowMap = make(map[byte]int, len(target))
	var valid int

	for right < len(src) {

		if _, ok := neededMap[src[right]]; ok {
			windowMap[src[right]] += 1
			if windowMap[src[right]] == neededMap[src[right]] {
				valid++
			}
		}
		right++

		for valid == len(neededMap) {

			if right-left == len(target) {
				return true
			}

			if count, ok := neededMap[src[left]]; ok && windowMap[src[left]] == count {
				valid--
				windowMap[src[left]] -= 1
			}
			left++
		}

	}

	return false
}

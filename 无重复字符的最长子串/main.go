package main

import "fmt"

/*

	给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

	示例 1:

		输入: s = "abcabcbb"
		输出: 3
		解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

	示例 2:

		输入: s = "bbbbb"
		输出: 1
		解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

	示例 3:

		输入: s = "pwwkew"
		输出: 3
		解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
    	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 

	提示：

		0 <= s.length <= 5 * 104
		s 由英文字母、数字、符号和空格组成

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func main() {
	length := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(length)
}

// dp思路
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	// 以i位置结尾的字符串 无重复最长长度
	var endOfILength = make([]int, len(s))
	// 以第一个字符结尾的子串无重复长度为1
	endOfILength[0] = 1
	//fmt.Println(endOfILength)
	for i := 1; i < len(s); i++ {
		lastCharLength := endOfILength[i-1]
		// 判断从前一个字符开始往前推lastCharLength个字符，
		// 如果没有与当前字符重复的，则endOfILength[i] = endOfILength[i-1]+1
		// 如果有，记为j,则以当前字符结尾的最大无重复子串为该位置(j,i]，不包含j，包含i

		// abcabcbb
		var isInterrupted bool
		for j := i - 1; j >= i-lastCharLength; j-- {
			if s[j] == s[i] {
				endOfILength[i] = i-j
				isInterrupted = true
				break
			}
		}

		if !isInterrupted{
			endOfILength[i] = endOfILength[i-1] + 1
		}
		//fmt.Println(endOfILength)
	}

	// 结束后获得以各个字符结尾的最长子串长度列表
	var maxLength  int
	for k:= 0 ; k < len(endOfILength) ; k++{
		if endOfILength[k] > maxLength {
			maxLength = endOfILength[k]
		}
	}

	return maxLength
}

/**

	时间复杂度分析
		i从0到len(s)-1, j 从i往前到m,s[m]=s[i]停止，m最小到0，所以m取值为0-i

		故最坏情况复杂度为 0 + 1 + 2 + 3 + ... + n-1， 分别对应 i从0到n-1时的计算量，
		即 (0 + n-1) * n = n^2 - n  => O(n^2)

		最好情况就是，每次j无需遍历，只需计算i的计算量，即为O(n)

	空间复杂度即为endOfILength的长度，可优化
*/


// dp + 滑动窗口思路
func lengthOfLongestSubstringV2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	// 以i位置结尾的字符串 无重复最长长度
	var endOfILength = make([]int, len(s))
	var locMap = make(map[byte]int)
	// 以第一个字符结尾的子串无重复长度为1
	endOfILength[0] = 1
	locMap[s[0]] = 0


	var k = 0
	for i := 1; i < len(s); i++ {
		// 判断从前一个字符开始往前推lastCharLength个字符，
		// 如果没有与当前字符重复的，则endOfILength[i] = endOfILength[i-1]+1
		// 如果有，记为j,则以当前字符结尾的最大无重复子串为该位置(j,i]，不包含j，包含i

		// abcabcbb
		j, ok := locMap[s[i]]
		if ok {
			// 存在
			// 删除locMap中 k到j之间位置出现的字符
			for m:=k; m<=j;m++ {
				delete(locMap,s[m])
			}
			k = j+1
			endOfILength[i] = i - k + 1

		}else {
			// 不存在
			endOfILength[i] = endOfILength[i-1]+1
		}
		locMap[s[i]] = i

	}

	// 结束后获得以各个字符结尾的最长子串长度列表
	var maxLength  int
	for k:= 0 ; k < len(endOfILength) ; k++{
		if endOfILength[k] > maxLength {
			maxLength = endOfILength[k]
		}
	}

	return maxLength
}

/**

	时间复杂度分析： 这种解法循环内部依旧有循环，时间复杂度依旧是最坏O(n^2),最好O(n)
	空间复杂度：因为引入了map，空间复杂度升高
*/

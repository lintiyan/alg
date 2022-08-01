package main

import "fmt"

/*

	编写一个函数来查找字符串数组中的最长公共前缀。

	如果不存在公共前缀，返回空字符串 ""。

	 

	示例 1：

	输入：strs = ["flower","flow","flight"]
	输出："fl"
	示例 2：

	输入：strs = ["dog","racecar","car"]
	输出：""
	解释：输入不存在公共前缀。

*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	if strs[0] == "" {
		return ""
	}

	var index int

	for i := 0; i < len(strs[0]); i++ {
		var isBreak bool
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 || strs[j][i] != strs[0][i] {
				isBreak = true
				if i == 0 {
					return ""
				}
				break
			}
		}
		if isBreak {
			break
		}
		index = i
	}

	return strs[0][:index+1]
}

func main() {

	//var strs = []string{"flower","flow","flight"}
	//var strs = []string{"flower","flow","floght"}
	var strs = []string{"dog", "racecar", "car"}
	// var strs = []string{"ab", "a"}
	//var strs = []string{"a", "a", "b"}
	//var strs = []string{"aaa","aa","aaa"}

	fmt.Println(longestCommonPrefix(strs))
}

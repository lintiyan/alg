package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
	"math"
)

func main() {

	// 01. 输入一个字符串ring代表圆盘上的字符，再输入一个字符串key代表你需要拨动圆盘输入的字符串
	// 返回输入这个key至少进行多少次操作（拨动一次和按下按钮都算一次操作）
	fmt.Println(findRotateSteps("godding", "gd"))

}

func findRotateSteps(ring string, key string) int {

	var charIndex = make(map[byte][]int)
	for i, c := range ring {
		charIndex[byte(c)] = append(charIndex[byte(c)], i)
	}

	return findRotateStepsFromTop(ring, key, 0, 0, charIndex)
}

func findRotateStepsFromTop(ring string, key string, i, j int, charIndex map[byte][]int) int {
	if j == len(key) {
		return 0
	}

	// 如果i指向的字符刚好等于j指向的字符，则只需要按下按钮，i不变，j需要加1
	indexs := charIndex[key[j]]
	var res = math.MaxInt
	for _, index := range indexs {

		delta := abs(i - index)
		delta = basicStruct.GetMin(len(ring)-index+i, delta)

		res = basicStruct.GetMin(res, findRotateStepsFromTop(ring, key, index, j+1, charIndex)+delta+1)
	}
	return res

}

func abs(num int) int {
	if num < 0 {
		return 0 - num
	}
	return num
}

package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
	"math"
)

func main() {

	// 01. 现在有n个城市,分别用0,1...n-1这些序号来表示,城市之间的航线用三元组[from, to, price]来表示
	// 请计算再k次中转之内,从src到dst所需要的最小花费是多少?如果无法到达,返回-1
	fmt.Println(findCheapestPrice(3,[][]int{{0,1,100},{1,2,100},{0,2,500}}, 0,2,0))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	// 构建某个能到达某个城市的 关联城市及价格信息
	var toMap = make(map[int][][]int)
	for _, flight := range flights {
		var toItem = []int{flight[0], flight[2]}
		toMap[flight[1]] = append(toMap[flight[1]], toItem)
	}
	res := findCheapestPriceFromTop(n, src, dst, k, toMap)
	if res == math.MaxInt{
		return -1
	}
	return res
}

func findCheapestPriceFromTop(n int, src int, dst int, k int, toMap map[int][][]int) int {

	if src == dst {
		return 0
	}

	if k <= 0 {
		return -1
	}

	dstMsg, ok := toMap[dst]
	if !ok {
		return -1
	}

	var res = math.MaxInt
	for _, item := range dstMsg {
		from := item[0]
		price := item[1]
		fromRes := findCheapestPriceFromTop(n, src, from, k-1, toMap)
		if fromRes < 0 {
			continue
		}else {
			res = basicStruct.GetMin(fromRes + price, res)
		}
	}

	return res
}

package main

import (
	"fmt"
	"sort"
)

func main() {

	// 01. 优势洗牌。 给你两个长度相等的数组nums1和nums2，请你重新组织nums1中元素的位置，使得nums1的优势最大化。
	// 如果nums1[i] > nums2[i]，就是说nums1在索引i上对nums2有优势。优势最大化就是说让你重新组织nums1，让尽可能多的nums1[i]>nums2[i]
	advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11})
}

func advantageCount(nums1, nums2 []int) {

	// 将nums1按升序排序
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] > nums1[j]
	})


	// 构建一个结构，以值和值的索引组成，然后按照值排序，得到nums2中马匹的战力排行，以及各个马匹对应的索引位置
	var items ItemSlice
	for i := 0; i < len(nums2); i++ {
		item := Item{nums2[i], i}
		items = append(items, item)
	}
	sort.Sort(items)

	var numsMap = make(map[int]int, len(nums1))	// 存储结果
	var min = 0						// 左指针，指向当前剩余最大值
	var max = len(nums1) - 1		// 右指针，指向当前剩余最小值

	var q = 0	// items当前遍历位置
	for q < len(nums2) {
		if nums1[min] <= items[q].value {			// 如果nums1[min] <= items[q].value, 意味着nums1中的上等马也比不过nums2的上等马
			numsMap[items[q].index] = nums1[max]	// 将nums1[max] 这个下等马放在items[q].index位置对线nums2的上等马，必输，所以摆烂
			max--									// 下等马用了一只
		}else if nums1[min] > items[q].value {		// 如果nums1[min] > items[q].value，意味着nums1中的最大值比nums2的最大值值大
			numsMap[items[q].index] = nums1[min]	// 那么这个上等马就放在items[q].index位置，来对线对方的上等马，能赢那就赢
			min++									// 上等马用了一只
		}
		q++											// 不论怎样，nums2必定要消耗一匹马
	}

	// 按照map构建nums1
	for i := 0; i < len(nums1); i++ {
		nums1[i] = numsMap[i]
	}

	fmt.Println(nums1)
	fmt.Println(nums2)
}

type Item struct {
	value int
	index int
}

type ItemSlice []Item

func (a ItemSlice) Less(i, j int) bool {
	return a[i].value > a[j].value
}

func (a ItemSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ItemSlice) Len() int {
	return len(a)
}

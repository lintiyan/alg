package basicStruct

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Traverse() {
	cur := l
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func SwapInt(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func SwapBytes(nums []byte, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func GetMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func GetMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

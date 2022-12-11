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

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	fmt.Print(node.Val, " ")
	node.Left.Traverse()
	node.Right.Traverse()
}

func CopyRes(src []int) []int {
	var dst = make([]int, len(src))
	copy(dst, src)
	return dst
}

func ContainsItem(nums []int, target int)bool {
	for _, item := range nums{
		if item == target {
			return true
		}
	}
	return false
}

func SliceSum(nums []int) int {
	var res int
	for _, item := range nums {
		res += item
	}
	return res
}
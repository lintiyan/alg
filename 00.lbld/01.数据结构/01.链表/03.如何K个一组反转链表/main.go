package main

import (
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	//reverseABTest()

	// 0.1 K个一组翻转链表
	reverseKGroupTest()
	reverseKGroupNewTest()
}

func reverseKGroupTest() {
	node1 := &basicStruct.ListNode{Val: 1}
	node2 := &basicStruct.ListNode{Val: 2}
	node3 := &basicStruct.ListNode{Val: 3}
	node4 := &basicStruct.ListNode{Val: 4}
	node5 := &basicStruct.ListNode{Val: 5}
	node6 := &basicStruct.ListNode{Val: 6}
	node7 := &basicStruct.ListNode{Val: 7}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	newHead := reverseKGroup(node1, 3)
	newHead.Traverse()
}

func reverseKGroup(list *basicStruct.ListNode, k int) *basicStruct.ListNode {

	if list == nil {
		return list
	}
	var cur = list

	// 不足k个不翻转
	for i := 1; i < k; i++ {
		cur = cur.Next
		if cur == nil {
			return list
		}
	}

	// 返回翻转A到B的节点后，并返回翻转部分的头和尾
	newHead, tempTail := reverseAB(list, cur)
	// 递归对子链表执行K个一组翻转
	tempHead := reverseKGroup(tempTail.Next, k)
	// 翻转后的链表与之前的翻转的尾部连接起来
	tempTail.Next = tempHead

	return newHead
}

func reverseABTest() {
	node1 := &basicStruct.ListNode{Val: 1}
	node2 := &basicStruct.ListNode{Val: 2}
	node3 := &basicStruct.ListNode{Val: 3}
	node4 := &basicStruct.ListNode{Val: 4}
	node5 := &basicStruct.ListNode{Val: 5}
	node6 := &basicStruct.ListNode{Val: 6}
	node7 := &basicStruct.ListNode{Val: 7}
	node8 := &basicStruct.ListNode{Val: 8}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8

	newHead, tail := reverseAB(node1, node5)
	newHead.Traverse()
	tail.Traverse()
}

// 此方法执行的是左闭右闭区间的翻转，事实上，执行左闭右开区间的翻转更简单
func reverseAB(list *basicStruct.ListNode, target *basicStruct.ListNode) (*basicStruct.ListNode, *basicStruct.ListNode) {

	var pre *basicStruct.ListNode
	cur := list

	targetNext := target.Next

	for cur != targetNext {

		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	list.Next = cur

	return target, list
}


func reverseKGroupNewTest() {
	node1 := &basicStruct.ListNode{Val: 1}
	node2 := &basicStruct.ListNode{Val: 2}
	node3 := &basicStruct.ListNode{Val: 3}
	node4 := &basicStruct.ListNode{Val: 4}
	node5 := &basicStruct.ListNode{Val: 5}
	node6 := &basicStruct.ListNode{Val: 6}
	node7 := &basicStruct.ListNode{Val: 7}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	newHead := reverseKGroupNew(node1, 3)
	newHead.Traverse()
}

func reverseKGroupNew(list *basicStruct.ListNode, k int) *basicStruct.ListNode {

	if list == nil {
		return list
	}
	var cur = list

	// 不足k个不翻转
	for i := 0; i < k; i++ {
		if cur == nil {
			return list
		}
		cur = cur.Next
	}

	// 返回翻转A到B的节点后的新的头节点，此时 list节点成了翻转后的尾节点
	newHead := reverseABNew(list, cur)
	// 递归对子链表执行K个一组翻转
	tempHead := reverseKGroupNew(cur, k)
	// 连接之前翻转的节点，和之后递归翻转的链表
	list.Next = tempHead
	return newHead
}

// 此方法执行左闭右开区间的翻转
func reverseABNew(list *basicStruct.ListNode, target *basicStruct.ListNode) *basicStruct.ListNode{

	var pre *basicStruct.ListNode
	cur := list
	for cur != target {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
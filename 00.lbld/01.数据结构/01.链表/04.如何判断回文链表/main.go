package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	// 01.后续遍历判断链表是否回文链表
	isPalindromeTest()
}

func isPalindromeTest() {
	node1 := &basicStruct.ListNode{Val: 1}
	node2 := &basicStruct.ListNode{Val: 2}
	node3 := &basicStruct.ListNode{Val: 3}
	node4 := &basicStruct.ListNode{Val: 4}
	node5 := &basicStruct.ListNode{Val: 3}
	node6 := &basicStruct.ListNode{Val: 2}
	node7 := &basicStruct.ListNode{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Println(isPalindrome01(node1))
	fmt.Println(isPalindrome02(node1))
}

var left *basicStruct.ListNode

func isPalindrome01(head *basicStruct.ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(head *basicStruct.ListNode) bool {
	if head == nil {
		return true
	}

	res := traverse(head.Next)
	res = res && head.Val == left.Val
	left = left.Next
	return res
}

func isPalindrome02(head *basicStruct.ListNode) bool {

	if head == nil {
		return true
	}

	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 将slow及其后的链表翻转
	newHead := reverse(slow)

	cur1 := head
	cur2 := newHead

	var isPalindrome bool
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			isPalindrome = false
			break
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	if cur1 == nil || cur2 == nil {
		isPalindrome = true
	}

	// 将slow及其后的链表再次翻转，修复到原来
	// 此时的链表结构 1(head)->2->3->4(slow)<-3<-2<-1(newHead)，其中slow.Next指向nil
	reverse(newHead)
	// 此时的链表结构 1(head)->2->3->4(slow)->3->2->1(newHead)，已经还原
	head.Traverse()
	return isPalindrome
}

func reverse(head *basicStruct.ListNode) *basicStruct.ListNode {

	if head.Next == nil {
		return head
	}

	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

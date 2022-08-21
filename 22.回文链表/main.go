package main

import "fmt"

/*

	给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

	 

	示例 1：


	输入：head = [1,2,2,1]
	输出：true
	示例 2：


	输入：head = [1,2]
	输出：false
	 

	提示：

	链表中节点数目在范围[1, 105] 内
	0 <= Node.val <= 9
	 

	进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	var slow, fast = head, head

	// 1234321
	for fast != nil {
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		}else{
			break
		}
	}

	reverseHead := reverseList(slow)

	var cur1 = head
	var cur2 = reverseHead
	var isPalindrome = true
	for cur2 != nil {
		if cur1.Val != cur2.Val {
			isPalindrome = false
		}
		cur1=cur1.Next
		cur2=cur2.Next
	}

	reverseList(reverseHead)

	return isPalindrome
}

// 翻转链表 1-2-3-4-1-2-3
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	pre := head
	cur := head.Next
	for cur != nil {

		next := cur.Next
		pre.Next = next
		cur.Next = head

		head = cur
		cur = next
	}
	return head
}

func main() {

	var head = &ListNode{Val: 1}
	var node2 = &ListNode{Val: 2}
	var node3 = &ListNode{Val: 3}
	var node4 = &ListNode{Val: 4}
	var node5 = &ListNode{Val: 3}
	var node6 = &ListNode{Val: 2}
	var node7 = &ListNode{Val: 1}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	TraverseList(head)
	fmt.Println(isPalindrome(head))
	TraverseList(head)
}

func TraverseList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d-", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

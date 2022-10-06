package main

import (
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	// 01. 翻转单链表
	//reverseTest()

	// 02. 翻转链表的前n个节点
	//reverseNTest01()
	reverseNTest02()

	// 03. 翻转链表的一部分
	//reverseBetweenTest()
}

func reverseTest() {
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
	newHead := reverse(node1)
	newHead.Traverse()
}

func reverse(list *basicStruct.ListNode) *basicStruct.ListNode {

	if list.Next == nil {
		return nil
	}

	last := reverse(list.Next)
	list.Next.Next = list
	list.Next = nil
	return last
}

func reverseNTest01() {
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

	newHead := reverseN01(node1, 3)
	newHead.Traverse()
}

func reverseN01(list *basicStruct.ListNode, n int) *basicStruct.ListNode {
	// 找到链表的第n+1个节点
	var nNode = list
	for i := 0; i < n; i++ {
		nNode = nNode.Next
	}
	//fmt.Println(nNode.Val)
	return reverseNBasic01(list, nNode)
}

func reverseNBasic01(list *basicStruct.ListNode, target *basicStruct.ListNode) *basicStruct.ListNode{

	if list.Next == target {
		return list
	}

	last := reverseNBasic01(list.Next, target)
	list.Next.Next = list
	list.Next = target

	return last
}

func reverseNTest02(){
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

	newHead := reverseN02(node1, 3)
	newHead.Traverse()
}

var nNode = &basicStruct.ListNode{}
func reverseN02(list *basicStruct.ListNode, n int) *basicStruct.ListNode {
	if n == 1 {
		nNode = list.Next
		return list
	}
	last := reverseN02(list.Next, n-1)
	list.Next.Next = list
	list.Next = nNode

	return last
}

func reverseBetweenTest(){
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

	newHead := reverseBetween(node1, 2,4)
	newHead.Traverse()
}

func reverseBetween(list *basicStruct.ListNode, m int, n int) *basicStruct.ListNode{

	if m == 1 {
		return reverseN01(list, n)
	}

	list.Next = reverseBetween(list.Next, m-1, n-1)
	return list
}
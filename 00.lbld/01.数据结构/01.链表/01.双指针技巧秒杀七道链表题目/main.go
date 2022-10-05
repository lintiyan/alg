package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	// 1. 合并两个有序链表
	//mergeTwoListsTest()

	// 2. 合并K个有序链表
	//mergeKListsTest()

	// 3. 单链表的倒数第K个节点
	//findFromEndTest()

	// 4. 删除链表的倒数第k个节点
	//removeKthFromEndTest()

	// 5. 找到单链表的中点
	//middleNodeTest()

	// 6. 判断链表是否包含环
	//hasCycleTest()

	// 7. 判断链表是否有环，如果有返回环的开始节点
	//detectCycleTest()

	// 8. 两个链表是否相交
	getIntersectionNodeTest()
}

func mergeTwoListsTest() {
	node1 := &basicStruct.ListNode{Val: 1}
	node2 := &basicStruct.ListNode{Val: 2}
	node3 := &basicStruct.ListNode{Val: 3}
	node4 := &basicStruct.ListNode{Val: 4}
	node5 := &basicStruct.ListNode{Val: 5}
	node6 := &basicStruct.ListNode{Val: 6}
	node7 := &basicStruct.ListNode{Val: 7}

	node1.Next = node4
	node4.Next = node7

	node2.Next = node3
	node3.Next = node5
	node5.Next = node6

	node1.Traverse()
	node2.Traverse()
	newHead := mergeTwoLists(node1, node2)
	newHead.Traverse()
}

func mergeTwoLists(list1, list2 *basicStruct.ListNode) *basicStruct.ListNode {

	var dummy = &basicStruct.ListNode{}
	var p, q *basicStruct.ListNode

	p = list1
	q = list2
	cur := dummy
	for p != nil && q != nil {

		if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}

		cur = cur.Next
	}

	if p != nil {
		cur.Next = p
	}

	if q != nil {
		cur.Next = q
	}

	return dummy.Next
}

func mergeKListsTest() {
	var lists []*basicStruct.ListNode
	mergeKLists(lists)
}

func mergeKLists(nodes []*basicStruct.ListNode) *basicStruct.ListNode {

	// 需要小顶堆的数据结构，golang貌似没有
	return nil
}

func findFromEndTest() {
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

	lastKNode := findFromEnd(node1, 7)
	fmt.Println(lastKNode.Val)
}

func findFromEnd(list *basicStruct.ListNode, k int) *basicStruct.ListNode {

	var dummy = &basicStruct.ListNode{}
	dummy.Next = list

	pre := dummy
	for i := 0; i < k; i++ {
		if pre != nil {
			pre = pre.Next
		} else {
			return nil
		}
	}

	last := dummy
	for pre.Next != nil {
		pre = pre.Next
		last = last.Next
	}

	return last.Next
}

func removeKthFromEndTest() {
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
	newHead := removeKthFromEnd(node1, 7)
	newHead.Traverse()
}

func removeKthFromEnd(list *basicStruct.ListNode, k int) *basicStruct.ListNode {

	var dummy = &basicStruct.ListNode{}
	dummy.Next = list

	var pre = dummy
	for i := 0; i < k; i++ {
		if pre != nil {
			pre = pre.Next
		} else {
			return nil
		}
	}

	last := dummy
	for pre.Next != nil {
		pre = pre.Next
		last = last.Next
	}

	last.Next = last.Next.Next
	return dummy.Next
}

func middleNodeTest() {
	node1 := &basicStruct.ListNode{Val: 1}
	node2 := &basicStruct.ListNode{Val: 2}
	node3 := &basicStruct.ListNode{Val: 3}
	node4 := &basicStruct.ListNode{Val: 4}
	//node5 := &basicStruct.ListNode{Val: 5}
	//node6 := &basicStruct.ListNode{Val: 6}
	//node7 := &basicStruct.ListNode{Val: 7}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	//node4.Next = node5
	//node5.Next = node6
	//node6.Next = node7

	fmt.Println(middleNode(node1).Val)
}

func middleNode(list *basicStruct.ListNode) *basicStruct.ListNode {

	var slow, fast = list, list

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func hasCycleTest() {
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

	node7.Next = node4
	fmt.Println(hasCycle(node1))
}

func hasCycle(list *basicStruct.ListNode) bool {
	var slow, fast = list, list

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}

func detectCycleTest() {
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

	node7.Next = node4

	fmt.Println(detectCycle(node1).Val)
}

func detectCycle(list *basicStruct.ListNode) *basicStruct.ListNode {
	// 快慢指针从链表起点开始走，两者相遇的时候，慢指针走了n步，则快指针走了2n步
	// 假设环的长度是m，链表起点到环的起点的距离为c，快慢指针相遇时距离起点的长度为d
	// n = c + p * m + d, 2n = c + q * m + d
	// 2n - 2*n = c + q * m + d - 2(c + p * m + d) = (q - 2p) * m - c - d = 0
	// c = (q - 2p) * m - d, 令 q-2p = 1, 则c= m-d ，即从链表起点到环起点的距离c 等于 环的长度m减去相遇点距离环起点的长度d
	// 如果在相遇点，如果有两个指针分别从链表的开始节点和相遇点 同时出发，两者相遇的点即为环的开始节点。

	var slow, fast = list, list
	var isCycle bool
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}

	fast = list
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func getIntersectionNodeTest(){
	node1 := &basicStruct.ListNode{Val: 1}
	node2 := &basicStruct.ListNode{Val: 2}
	node3 := &basicStruct.ListNode{Val: 3}
	node4 := &basicStruct.ListNode{Val: 4}
	node5 := &basicStruct.ListNode{Val: 5}
	node6 := &basicStruct.ListNode{Val: 6}
	node7 := &basicStruct.ListNode{Val: 7}

	node1.Next = node4
	node4.Next = node7

	node2.Next = node3
	node3.Next = node5
	node5.Next = node6

	node6.Next = node7
	fmt.Println(getIntersectionNode(node1,node2))
}

func getIntersectionNode(list1, list2 *basicStruct.ListNode) *basicStruct.ListNode {

	var p, q = list1, list2

	for p != q {

		if p == nil {
			p = list2
		}else {
			p = p.Next
		}

		if q == nil {
			q = list1
		}else {
			q = q.Next
		}

	}

	return p
}

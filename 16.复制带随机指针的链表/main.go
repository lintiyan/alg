package main

import "fmt"

/*
	给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

	构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。

	例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。

	返回复制链表的头节点。

	用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

	val：一个表示 Node.val 的整数。
	random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
	你的代码 只 接受原链表的头节点 head 作为传入参数。


	示例 1：

	输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
	输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

	示例 2：

	输入：head = [[1,1],[2,1]]
	输出：[[1,1],[2,1]]


	示例 3：

	输入：head = [[3,null],[3,0],[3,null]]
	输出：[[3,null],[3,0],[3,null]]


*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	var cur = head

	for cur != nil {

		newNode := &Node{
			Val: cur.Val,
		}
		newNode.Next = cur.Next
		cur.Next = newNode

		cur = newNode.Next
	}

	cur = head
	next := cur.Next
	for cur != nil {
		if cur.Random != nil {
			next.Random = cur.Random.Next
		}

		cur = cur.Next.Next
		if cur != nil {
			next = cur.Next
		}
	}

	var newHead = head.Next

	cur = head
	next = newHead

	for cur != nil{

		if cur.Next != nil {
			cur.Next = cur.Next.Next
			cur = cur.Next
		}

		if next != nil && next.Next != nil {
			next.Next = next.Next.Next
			next = next.Next
		}
	}

	return newHead
}

func main() {


	head0 := &Node{Val: 0}
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	head0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	head0.Random = node2
	node1.Random = node3
	node2.Random = node1
	node3.Random = head0
	node4.Random = node4

	newhead := copyRandomList(head0)
	PrintList(newhead)
}


func PrintList (node *Node){
	fmt.Println("========================")
	cur := node
	for cur != nil {
		//fmt.Println("val=",cur.Val)
		if cur.Random != nil {
			fmt.Println("val=",cur.Val, "rand=", cur.Random.Val)
		}else {
			fmt.Println("val=",cur.Val)
		}

		cur = cur.Next
	}

	fmt.Println("========================")
}
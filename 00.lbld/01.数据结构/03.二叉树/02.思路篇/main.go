package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

func main() {

	// 01. 原地翻转二叉树
	reverseTreeTest()
}

func reverseTreeTest(){
	node1 := &basicStruct.TreeNode{Val: 1}
	node2 := &basicStruct.TreeNode{Val: 2}
	node3 := &basicStruct.TreeNode{Val: 3}
	node4 := &basicStruct.TreeNode{Val: 4}
	node5 := &basicStruct.TreeNode{Val: 5}
	node6 := &basicStruct.TreeNode{Val: 6}
	node7 := &basicStruct.TreeNode{Val: 7}

	//	     1
	//	   2    3
	//	 4    5
	// 6       7

	//	     1
	//	   3    2
	//	    5     4
	//     7        6
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node3.Left = node5
	node4.Left = node6
	node5.Right = node7

	// 原来顺序
	node1.Traverse()
	fmt.Println()
	reverseTreeBasic(node1)

	// 翻转后的顺序
	node1.Traverse()
	fmt.Println()

	newHead := reverseTreeSplit(node1)
	// 再次翻转回原来的顺序
	newHead.Traverse()
}

// 以遍历的思维
func reverseTreeBasic(node *basicStruct.TreeNode){
	if node == nil {
		return
	}

	tmp := node.Left
	node.Left = node.Right
	node.Right = tmp
	reverseTreeBasic(node.Left)
	reverseTreeBasic(node.Right)
}

// 以分解的思维
func reverseTreeSplit(node *basicStruct.TreeNode) *basicStruct.TreeNode{
	if node == nil{
		return nil
	}

	left := reverseTreeSplit(node.Left)
	right := reverseTreeSplit(node.Right)

	node.Left = right
	node.Right = left

	return node
}
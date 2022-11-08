package main

import (
	"fmt"
	"github.com/alg/00.lbld/basicStruct"
)

// 二叉树相关题目的思想一般是两种：
// 1. 能否通过一次遍历得出答案 --- 回溯思想
// 2. 能否将问题拆分成子问题然后拼接得出答案 --- dp思想

// 二叉树的先中后遍历无非就是处理当前节点相关的问题的时机不同而已。
// 但是后序遍历有其特殊之处，后续遍历时能得到子树的处理结果，所以当题目所求与子树的遍历结果相关时，一般需要使用后续遍历了

func main() {

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
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node3.Left = node5
	node4.Left = node6
	node5.Right = node7

	// 01. 二叉树的最大深度
	// maxDepth(node1)

	// 02. 二叉树的直径
	diameterOfTree(node1)
	// 03. 二叉树的前序遍历

}

func maxDepth(node *basicStruct.TreeNode) {
	// 回溯算法
	maxDepthTraceback(node, 1)
	fmt.Println(maxDepthV)

	// 动态规划
	fmt.Println(maxDepthDP(node))
}

// 回溯算法
var maxDepthV int
func maxDepthTraceback(node *basicStruct.TreeNode, depth int) {

	// base case ，如果到达叶子节点，则统计深度
	if node.Left == nil && node.Right == nil {
		if maxDepthV < depth {
			maxDepthV = depth
		}
		return
	}

	if node.Left != nil {
		depth++
		maxDepthTraceback(node.Left, depth)
		depth--
	}

	if node.Right != nil {
		depth++
		maxDepthTraceback(node.Right, depth)
		depth--
	}
	return
}

// dp思想
func maxDepthDP(node *basicStruct.TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}

	left := maxDepthDP(node.Left)
	right := maxDepthDP(node.Right)

	return basicStruct.GetMax(left, right) + 1
}

func diameterOfTree(node *basicStruct.TreeNode){
	// 普通方法
	diameterOfTreeTraverse(node)
	fmt.Println(diameterOfT)

	// 优化方法
	diameterOfTreeDepth(node)
	fmt.Println(diameterOfTD)
}

var diameterOfT int
func diameterOfTreeTraverse(node *basicStruct.TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil{
		return
	}

	left := maxDepthDP(node.Left)
	right := maxDepthDP(node.Right)
	if diameterOfT < left + right  {
		diameterOfT = left + right
	}

	diameterOfTreeTraverse(node.Left)
	diameterOfTreeTraverse(node.Right)

	return
}

var diameterOfTD int
func diameterOfTreeDepth(node *basicStruct.TreeNode) int {

	if node == nil {
		return 0
	}

	left := diameterOfTreeDepth(node.Left)
	right := diameterOfTreeDepth(node.Right)
	diameterOfTD = basicStruct.GetMax(diameterOfTD, left + right)

	return basicStruct.GetMax(left, right) + 1
}
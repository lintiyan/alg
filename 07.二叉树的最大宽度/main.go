package main

import (
	"fmt"
	"math"
)

/**
	给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

	每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

	示例 1:

	输入:

			   1
			 /   \
			3     2
		   / \     \
		  5   3     9

	输出: 4
	解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。

	示例 2:

	输入:

			  1
			 /
			3
		   / \
		  5   3

	输出: 2
	解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
	示例 3:

	输入:

			  1
			 / \
			3   2
		   /
		  5

	输出: 2
	解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。
	示例 4:

	输入:

			  1
			 / \
			3   2
		   /     \
		  5       9
		 /         \
		6           7
	输出: 8
	解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。


	思路：

		给每个节点编号，如果当前节点的编号是 i， 那么它的左子节点的编号就是 2*i，右子节点的编号是 2*i+1
		层序遍历树，遍历每一层时，记录最左节点的编号，记录最右节点的编号，差值为当前层的宽度

		      1
		   /     \
		2          3
	   /          / \
	  4          5   9
	 /            \
	6              7
*/

func main() {

	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}


	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node3.Left = node5
	node4.Left = node6
	node5.Right = node7

	fmt.Println(widthOfBinaryTree(node1))
}

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	var index int
	root.Val = 1	// 将val值改为当前节点的编号
	queue = append(queue, root)	// 入队

	var maxWidth int	// 该树的最大宽度

	for len(queue) - index > 0 {	// 队列不为空则循环

		curLevelSize := len(queue) - index	// 每次走到这时，代表上一层的节点已经遍历结束，所以剩下的节点数表示当前层的节点数
		var isFirst = true		// 是否当前层的第一个节点
		var left, right int		// 记录最左及最右节点编号

		for curLevelSize > 0 {	// 遍历当前层，直到curLevelSize=0 表示当前层遍历结束

			cur := queue[index]	// 取出当前节点
			index++				// 索引值记得更新，代表遍历到了队列中的哪个节点

			if isFirst {		// 当前层的第一个，记录left
				left = cur.Val
				isFirst = false
			}
			right = cur.Val		// 修改right值

			maxWidth = int(math.Max(float64(maxWidth), float64(right - left + 1)))	// 针对当前层的每个节点，修改maxWidth值

			if cur.Left != nil {		// 左节点不为空，修改编号并入队
				cur.Left.Val = 2 * cur.Val
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {		// 右节点不为空，修改编号并入队
				cur.Right.Val = 2 * cur.Val + 1
				queue = append(queue, cur.Right)
			}

			curLevelSize--	// 当前层剩余节点数减1
		}
	}

	return maxWidth
}
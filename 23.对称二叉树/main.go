package main

/*
	给你一个二叉树的根节点 root ， 检查它是否轴对称。

	 

	示例 1：


	输入：root = [1,2,2,3,4,4,3]
	输出：true
	示例 2：


	输入：root = [1,2,2,null,3,null,3]
	输出：false
	 

	提示：

	树中节点数目在范围 [1, 1000] 内
	-100 <= Node.val <= 100
	 

	进阶：你可以运用递归和迭代两种方法解决这个问题吗？

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 校验两颗树是否对称
func check(p, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		// 如果根节点值相等，则校验p的左子树和q的右子树是否对称、 校验p的右子树和q的左子树是否对称
		return check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return false
}


// 对称树的特点是，根节点的左右子树是对称的
func isSymmetric(root *TreeNode) bool {

	return check(root, root)
}

func main() {

}

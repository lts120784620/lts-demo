package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.236. 二叉树的最近公共祖先
描述：
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
	百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
思路：
    1、递归，重点是确定递归的返回值是什么，左子树找到就返回那个节点，一直传递到root
时间：
    1、2021/10/20
*/

func TestLowestCommonAncestor(t *testing.T) {
	a := NewBinTreeByArrays([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	fmt.Println(a.String())
	p := a.Left.Right
	q := a.Right.Right
	fmt.Println(lowestCommonAncestor(a, p, q).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

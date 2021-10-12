package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.100 相同的树
描述：
	给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
	如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
思路：
    1、递归逐个对比即可，考虑一个先为nil的情况
时间：
    1、2021/9/29
*/

func TestIsSameTree(t *testing.T) {
	a := NewBinTreeByArrays([]interface{}{1, 2, 3, 4})
	b := NewBinTreeByArrays([]interface{}{1, 2, 3})
	fmt.Println(isSameTree(a, b))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil {
		return false
	}

	if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

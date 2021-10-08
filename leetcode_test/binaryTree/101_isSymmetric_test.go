package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.101 对称二叉树
描述：
	给定一个二叉树，检查它是否是镜像对称的。
	例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
思路：
    1、可转换为如何判断左子树等于右子树，声明两个指针，一个左右遍历，一个右到左遍历，判断是否相等.与101题思路类似
时间：
    1、2021/9/29
*/

func TestIsSymmtric(t *testing.T) {
	a := NewTreeByArrays([]interface{}{1, 2, 2, 3, 4, 4, 3, 0})
	fmt.Println(isSymmetric(a))
}

func isSymmetric(root *TreeNode) bool {
	return checkEqual(root, root)
}

func checkEqual(p *TreeNode, q *TreeNode) bool {
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

	return checkEqual(p.Left, q.Right) && checkEqual(p.Right, q.Left)
}

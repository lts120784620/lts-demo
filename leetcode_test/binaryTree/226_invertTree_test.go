package binaryTree

import "testing"

/**
No.226 翻转二叉树
描述：
	翻转一棵二叉树
思路：
    1、递归，左右节点交换即可
时间：
    1、2021/10/1
*/

func TestInvertTree(t *testing.T) {
	invertTree(NewBinTreeByArrays([]interface{}{4, 2, 7, 1, 3, 6, 9}))
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	t := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = t
	return root
}

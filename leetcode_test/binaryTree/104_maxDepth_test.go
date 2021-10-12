package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.104 二叉树的最大深度
描述：
	给定一个二叉树，找出其最大深度。
	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
	说明: 叶子节点是指没有子节点的节点。
思路：
    1、递归法，左边和右边比较，res+大的即可
时间：
    1、2021/9/29
*/

func TestMaxDepth(t *testing.T) {
	a := NewBinTreeByArrays([]interface{}{1, 2, 3, 4})
	fmt.Println(maxDepth(a))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}

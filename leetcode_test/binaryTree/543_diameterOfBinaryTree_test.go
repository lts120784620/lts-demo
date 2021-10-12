package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.543 二叉树的直径
描述：
	给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
思路：
    1、递归法，求两个最远的叶子节点到跟的距离之和，用一个全局变量保存最大值，当左子节点+右子节点之合大于最大值时进行替换。ps：两个节点的距离不一定通过根节点
时间：
    1、2021/10/1
*/

var max int = 1

func TestDiameterOfBinaryTree(t *testing.T) {
	r := NewBinTreeByArrays([]interface{}{-10, 9, 20, nil, nil, 15, 7})
	fmt.Println(diameterOfBinaryTree(r))
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDeep(root)
	return max - 1
}

func maxDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDeep(root.Left)
	maxRight := maxDeep(root.Right)

	if maxLeft+maxRight+1 > max {
		max = maxLeft + maxRight + 1
	}

	if maxLeft > maxRight {
		return maxLeft + 1
	} else {
		return maxRight + 1
	}
}

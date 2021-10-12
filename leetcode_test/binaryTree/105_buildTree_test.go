package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.105 从前序与中序遍历序列构造二叉树
描述：
	给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。
思路：
    1、前序第一个为根结点，中序的根结点两侧为左右子节点，中序即可用分治的思路分别在左右数组进行树节点拼装
时间：
    1、2021/10/6
	2、2021/10/11
*/

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	fmt.Println(buildTree1(preorder, inorder).String())
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	if inorder == nil || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	// 找到pre在inorder的位置
	idx := 0
	for i := range inorder {
		if preorder[0] == inorder[i] {
			idx = i
			break
		}
	}
	// 分治
	root.Left = buildTree(preorder[1:len(inorder[:idx])+1], inorder[:idx])
	root.Right = buildTree(preorder[len(inorder[:idx])+1:], inorder[idx+1:])
	return root
}

// 啥也不是，没有一遍就写出来
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 按前序的第一个位置，把中序分成两个子数组
	root := &TreeNode{Val: preorder[0]}
	idx := 0
	for i, v := range inorder {
		if v == preorder[0] {
			idx = i
			break
		}
	}
	// 如何切割出前序数组的索引位置，才是关键
	root.Left = buildTree1(preorder[1:len(inorder[:idx])+1], inorder[:idx])
	root.Right = buildTree1(preorder[len(inorder[:idx])+1:], inorder[idx+1:])
	return root
}

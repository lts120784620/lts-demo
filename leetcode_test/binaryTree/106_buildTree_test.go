package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.106. 从中序与后序遍历序列构造二叉树
描述：
	根据一棵树的中序遍历与后序遍历构造二叉树。
	注意:
		你可以假设树中没有重复的元素。
	例如，给出
		中序遍历 inorder = [9,3,15,20,7]
		后序遍历 postorder = [9,15,7,20,3]
		返回如下的二叉树：
			3
		   / \
		  9  20
			/  \
		   15   7
思路：
    1、和105题思路差不多，难点在于如何找到数组的索引位置，后序[左,右，根]，中序[左，根，右]
时间：
    1、2021/10/11
*/

func TestBuildTreeByPost(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	fmt.Println(buildTreeByPost(inorder, postorder).String())
}

func buildTreeByPost(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 找到中序的位置，用于拆分分治数组
	idx := len(postorder) - 1
	root := &TreeNode{Val: postorder[idx]}

	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			idx = i
			break
		}
	}

	root.Left = buildTreeByPost(inorder[:idx], postorder[:len(inorder[:idx])])
	root.Right = buildTreeByPost(inorder[idx+1:], postorder[len(inorder[:idx]):len(postorder)-1])

	return root
}

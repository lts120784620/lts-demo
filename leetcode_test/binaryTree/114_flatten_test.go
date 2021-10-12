package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.114. 二叉树展开为链表
描述：
	给你二叉树的根结点 root ，请你将它展开为一个单链表：
	展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
	展开后的单链表应该与二叉树 先序遍历 顺序相同。
思路：
    1、递归，后序遍历，对结果操作将左子树接到右子树上，删除左子树
时间：
    1、2021/10/11
*/

func TestFlatten(t *testing.T) {
	r := NewBinTreeByArrays([]interface{}{1, 2, 5, 3, 4, nil, 6})
	flatten(r)
	fmt.Println(r.String())
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	l := root.Left
	r := root.Right

	root.Right = l

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = r
	root.Left = nil
}

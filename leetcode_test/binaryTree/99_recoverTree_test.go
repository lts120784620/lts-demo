package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.99 恢复二叉搜索树
描述：
	给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。
	进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？
思路：
    1、一种思路是bst中序遍历为有序数列，中序遍历，用pre保存前一节点，当出现pre<root时，记下x位置，在此判断出现pre<root,记下y位置，xy做交换即可
	2、莫里斯遍历（Morris）是一种二叉树的遍历方式，主要是可以用O(1)的空间复杂度实现二叉树的遍历，其核心思想是借助节点额外指针指向需要栈弹出的节点，在遍历时去掉指针，使整个二叉树近似于图的结构
时间：
    1、2021/10/7
*/

func TestRecoverTree(t *testing.T) {
	r := NewBinTreeByArrays([]interface{}{3, 1, 4, nil, nil, 2})
	recoverTree(r)
	fmt.Println(r.String())
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	x, y, _ := inorder(root, nil, nil, nil)
	// 交换节点
	t := x.Val
	x.Val = y.Val
	y.Val = t
}

func inorder(root *TreeNode, x, y, pre *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if root == nil {
		return x, y, pre
	}
	x, y, pre = inorder(root.Left, x, y, pre)
	if pre != nil && pre.Val > root.Val {
		// 记录第一个节点
		if x == nil {
			x = pre
		}
		// 记录第二个节点
		y = root
	}
	pre = root
	x, y, pre = inorder(root.Right, x, y, pre)
	return x, y, pre
}

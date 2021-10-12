package binaryTree

import (
	"fmt"
	"lts-demo/leetcode_test/stack"
	"testing"
)

/**
No.94 二叉树的中序遍历
描述：
	给定一个二叉树的根节点 root ，返回它的 中序 遍历。
思路：
    1、递归遍历，左->根->右的顺序
	2、迭代遍历，利用栈结构，每次都要找到最左节点
时间：
    1、2021/9/28
*/

func TestInorderTraversal(t *testing.T) {
	n := NewBinTreeByArrays([]interface{}{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(inorderTraversal(n))
	fmt.Println(inorderTraversal2(n))
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	s := stack.NewStack()
	node := root
	for !s.IsEmpty() || node != nil {
		// 找到最左节点，注意不要重复寻找，这里借用一个毒丸right=nil的时候
		for node != nil {
			s.Push(node)
			node = node.Left
		}
		// 弹出节点
		node = s.Pop().(*TreeNode)
		// 写入左节点
		res = append(res, node.Val)
		// 如果有右节点，放到栈中，继续按照找最左节点方式继续
		node = node.Right

	}
	return res
}

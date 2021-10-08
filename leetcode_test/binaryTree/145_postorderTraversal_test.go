package binaryTree

import (
	"fmt"
	"lts-demo/leetcode_test/stack"
	"testing"
)

/**
No.145 二叉树的后序遍历
描述：
	给定一个二叉树，返回它的 后序 遍历。
思路：
    1、递归法：实现左->右->根即可
	2、迭代法：利用栈结构，前半部分与中序遍历一致，找到最左节点，若右节点不存在则暂存节点，读取节点，向上遍历ßß
时间：
    1、2021/9/29
*/

func TestPostorderTraversal(t *testing.T) {
	node := NewTreeByArrays([]interface{}{2, nil, 3, nil, nil, nil, 1})
	fmt.Println(postorderTraversal(node))
	fmt.Println(postorderTraversal2(node))
	fmt.Println(postorderTraversal3(node))
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 迭代法：栈实现
func postorderTraversal2(root *TreeNode) []int {
	res := []int{}
	s := stack.NewStack()
	var prev *TreeNode
	node := root
	for !s.IsEmpty() || node != nil {
		// 找到最左节点
		for node != nil {
			s.Push(node)
			node = node.Left
		}
		// 不弹出，判断右节点是否存在，在则先读右节点在弹出
		node = s.Pop().(*TreeNode)
		if node.Right != nil && node.Right != prev {
			// 不能立刻读根，把根塞回去，先读右节点
			s.Push(node)
			node = node.Right
		} else {
			res = append(res, node.Val)
			// 读右节点之前，保存一下根，防止重复向右读
			prev = node
			// 放入毒丸，防止继续向左读
			node = nil
		}

	}
	return res
}

// 迭代法：数组代替栈
func postorderTraversal3(root *TreeNode) []int {
	res := []int{}
	s := []*TreeNode{}
	var prev *TreeNode
	node := root
	for len(s) > 0 || node != nil {
		// 找到最左节点
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		// 不弹出，判断右节点是否存在，在则先读右节点在弹出
		node = s[len(s)-1]
		s = s[:len(s)-1]
		if node.Right != nil && node.Right != prev {
			// 不能立刻读根，把根塞回去，先读右节点
			s = append(s, node)
			node = node.Right
		} else {
			res = append(res, node.Val)
			// 读右节点之前，保存一下根，防止重复向右读
			prev = node
			// 放入毒丸，防止继续向左读
			node = nil
		}

	}
	return res
}

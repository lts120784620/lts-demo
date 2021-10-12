package binaryTree

import (
	"fmt"
	"lts-demo/leetcode_test/stack"
	"testing"
)

/**
No.144 二叉树的前序遍历
描述：
	给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
思路：
    1、前序就是根->左->右，递归就写出来了
时间：
    1、2021/9/28
*/

func TestPreorderTraversal(t *testing.T) {
	node := NewBinTreeByArrays([]interface{}{1, 2, 3, 4, 5})
	fmt.Println(preorderTraversal(node))
	fmt.Println(preorderTraversal2(node))
	fmt.Println(preorderTraversal3(node))
}

// 递归法
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

// 迭代法
func preorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	s := stack.NewStack()
	s.Push(root)
	for !s.IsEmpty() {
		node := s.Pop().(*TreeNode)
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
	return res
}

// 迭代法：用简易数组代替栈，因为golang没有提供原生栈，面试常用数组代替
func preorderTraversal3(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	s := []*TreeNode{root}
	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		if node.Right != nil {
			s = append(s, node.Right)
		}
		if node.Left != nil {
			s = append(s, node.Left)
		}
	}
	return res
}

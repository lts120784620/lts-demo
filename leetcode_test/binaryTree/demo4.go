package binaryTree

import (
	"lts-demo/leetcode_test/common"
)

// 中序遍历，非递归
func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	s := common.NewStack()
	node := root
	for !s.IsEmpty() || node != nil {
		if node != nil {
			s.Push(node)
			node = node.Left
		} else {
			node = s.Pop().(*TreeNode)
			res = append(res, node.Val)
			node = node.Right
		}
	}
	return res
}

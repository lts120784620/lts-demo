package binaryTree

import "github.com/golang-collections/collections/stack"
import "github.com/golang-collections/go-datastructures/queue"



/**
二叉树的层级遍历，从叶子节点到根节点（从左到右）
 */
func LevelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := queue.New(100)
	q.Put(root)
	for q.Empty() {
		//var t []int
		t:= []int{}
		node := s.Pop()
		if node == nil {
			continue
		}
		t = append(t, node.(*TreeNode).Val)
		s.Push(node.(*TreeNode).Right)
		s.Push(node.(*TreeNode).Left)
	}
	return res
}

package binaryTree

import (
	"lts-demo/test/common"
	"fmt"
)

/**
二叉树的层级遍历，从根节点到叶子节点（从左到右）
 */
func LevelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if (root == nil) {
		return res
	}
	q := common.NewQueue()
	q.Add(root)
	for (!q.IsEmpty()) {
		column := []int{}
		// 用这个list代表上层节点数量
		length := q.Size()
		for i:= 0 ;i<length;i++ {
			node := q.Offer()
			column = append(column, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil{
				q.Add(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil{
				q.Add(node.(*TreeNode).Right)
			}
		}
		res = append(res,column)
	}
	return res
}
// 层序输出，叶子节点开始
func LevelOrderBottom2(root *TreeNode) [][]int {
	res := [][]int{}
	if (root == nil) {
		return res
	}
	q := common.NewQueue()
	q.Add(root)
	for (!q.IsEmpty()) {
		column := []int{}
		// 用这个list代表上层节点数量
		length := q.Size()
		for i:= 0 ;i<length;i++ {
			node := q.Offer()
			column = append(column, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil{
				q.Add(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil{
				q.Add(node.(*TreeNode).Right)
			}
		}

		res = append(append([][]int{},column),res...)
		fmt.Println(res)
	}
	return res
}

// 中序遍历，非递归
func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil{
		return res
	}
	s := common.NewStack()
	node := root
	for !s.IsEmpty() || node != nil{
		if  node != nil{
			s.Push(node)
			node = node.Left
		}else{
			node = s.Pop().(*TreeNode)
			res = append(res,node.Val)
			node = node.Right
		}
	}
	return res
}

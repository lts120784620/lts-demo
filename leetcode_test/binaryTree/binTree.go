package binaryTree

import "lts-demo/leetcode_test/queue"

/**
No.
描述：
	二叉树基础
思路：
    1、
时间：
    1、2021/9/28
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree() *TreeNode {
	return &TreeNode{}
}

func NewTreeByArrays(arr []interface{}) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	return getChild(arr, 0, &TreeNode{Val: arr[0].(int)})
}

// 递归方式，数组转换二叉树结构
func getChild(arr []interface{}, i int, root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if 2*i+1 < len(arr) && arr[2*i+1] != nil {
		root.Left = getChild(arr, 2*i+1, &TreeNode{Val: arr[2*i+1].(int)})
	}
	if 2*i+2 < len(arr) && arr[2*i+2] != nil {
		root.Right = getChild(arr, 2*i+2, &TreeNode{Val: arr[2*i+2].(int)})
	}
	return root
}

// dfs 层序遍历二叉树
func (n *TreeNode) String() [][]int {
	res := [][]int{}
	if n == nil {
		return res
	}
	q := queue.New()
	q.Enqueue(n)
	for !q.IsEmpty() {
		col := []int{}
		width := q.Len()
		for i := 0; i < width; i++ {
			node := q.Dequeue().(*TreeNode)
			if node == nil {
				continue
			}
			col = append(col, node.Val)
			q.Enqueue(node.Left)
			q.Enqueue(node.Right)
		}
		if len(col) > 0 {
			res = append(res, col)
		}
	}
	return res
}

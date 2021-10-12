package binaryTree

import "lts-demo/leetcode_test/queue"

/**
No.
描述：
	n叉树结构
思路：
    1、
时间：
    1、2021/10/7
*/

type Node struct {
	Val      int
	Children []*Node

	// 完美二叉树
	Left  *Node
	Right *Node
	Next  *Node
}

func NewTreeByArrays(arr []interface{}) *Node {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	return getTreeChild(arr, 0, &Node{Val: arr[0].(int)})
}

// 递归方式，数组转换二叉树结构
func getTreeChild(arr []interface{}, i int, root *Node) *Node {
	if root == nil {
		return nil
	}

	if 2*i+1 < len(arr) && arr[2*i+1] != nil {
		root.Left = getTreeChild(arr, 2*i+1, &Node{Val: arr[2*i+1].(int)})
	}
	if 2*i+2 < len(arr) && arr[2*i+2] != nil {
		root.Right = getTreeChild(arr, 2*i+2, &Node{Val: arr[2*i+2].(int)})
	}
	return root
}

// dfs 层序遍历二叉树
func (n *Node) String() [][]int {
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
			node := q.Dequeue().(*Node)
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

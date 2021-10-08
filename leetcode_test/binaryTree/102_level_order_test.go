package binaryTree

import (
	"fmt"
	"lts-demo/leetcode_test/queue"
	"testing"
)

/**
No.102 二叉树的层序遍历
描述：
	给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
思路：
    1、用队列保存节点，保存现有数据长度代表当前行数，遍历队列内的行数
时间：
    1、2021/9/28
*/

func TestLevelOrder(t *testing.T) {
	n := NewTreeByArrays([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(levelOrder(n))
}

func levelOrder(root *TreeNode) [][]int {
	// 保存结果
	var res [][]int
	// 队列用来存放下一层的所有节点
	q := queue.New()
	q.Enqueue(root)
	// 外层循环遍历一行节点
	for !q.IsEmpty() {
		col := make([]int, 0)
		// 记录本行长度，很关键
		width := q.Len()
		for i := 0; i < width; i++ {
			node := q.Dequeue().(*TreeNode)
			if node == nil {
				continue
			}
			col = append(col, node.Val)
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		if len(col) > 0 {
			res = append(res, col)
		}
	}

	return res
}

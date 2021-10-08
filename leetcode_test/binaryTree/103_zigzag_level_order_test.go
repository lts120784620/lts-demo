package binaryTree

import (
	"fmt"
	"lts-demo/leetcode_test/queue"
	"testing"
)

/**
No.103 二叉树的锯齿形层序遍历
描述：
	给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）类似z序遍历
思路：
    1、
时间：
    1、2021/9/28
*/

func TestZigzagLevelOrder(t *testing.T) {
	node := NewTreeByArrays([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(zigzagLevelOrder(node))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	q := queue.New()
	q.Enqueue(root)
	flag := true
	for !q.IsEmpty() {
		col := make([]int, 0)
		width := q.Len()
		for i := 0; i < width; i++ {
			node := q.Dequeue().(*TreeNode)
			if node == nil {
				continue
			}
			if flag {
				col = append(col, node.Val)
			} else {
				// 切片头插，哈哈哈太牛逼了
				col = append([]int{node.Val}, col...)
			}
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		flag = !flag
		if len(col) > 0 {
			res = append(res, col)
		}
	}

	return res
}

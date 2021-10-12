package binaryTree

import (
	"fmt"
	"lts-demo/leetcode_test/queue"
	"testing"
)

/**
No.107 二叉树的层序遍历 II
描述：
	给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
思路：
    1、找到102题的解法，对结果进行反转，done，tnnd!
时间：
    1、2021/9/30
*/

func TestLevelOrderBottom(t *testing.T) {
	r := NewBinTreeByArrays([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(levelOrderBottom(r))
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	q := queue.New()
	q.Enqueue(root)
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
	for i := 0; i < len(res)/2; i++ {
		t := res[i]
		res[i] = res[len(res)-1-i]
		res[len(res)-1-i] = t
	}
	return res
}

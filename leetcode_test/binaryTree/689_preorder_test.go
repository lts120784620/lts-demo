package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.589 N 叉树的前序遍历
描述：
	给定一个 N 叉树，返回其节点值的 前序遍历 。
	N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
	进阶：
		递归法很简单，你可以使用迭代法完成此题吗?
思路：
    1、递归法，左右节点的遍历改成循环遍历即可，先塞值，在递归子节点
时间：
    1、2021/10/7
*/

func TestPreorder(t *testing.T) {
	r := &Node{Val: 1}
	r.Children = []*Node{
		{Val: 2, Children: []*Node{
			{Val: 5},
		}},
		{Val: 3, Children: []*Node{
			{Val: 6},
		}},
		{Val: 4, Children: []*Node{
			{Val: 7},
		}},
	}
	fmt.Println(preorder(r))
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	for _, c := range root.Children {
		res = append(res, preorder(c)...)
	}
	return res
}

package binaryTree

import (
	"fmt"
	"lts-demo/test/common"
	"strconv"
)

//二叉树的层级平均数
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}
	q := common.NewQueue()
	q.Add(root)
	for !q.IsEmpty() {
		tmp := []int{}
		length := q.Size()
		var avg float64
		var count float64
		for i := 0; i < length; i++ {
			node := q.Offer()
			tmp = append(tmp, node.(*TreeNode).Val)
			count = float64(node.(*TreeNode).Val) + count
			if node.(*TreeNode).Left != nil {
				q.Add(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				q.Add(node.(*TreeNode).Right)
			}
		}
		// 分子是float才能得到浮点型结果
		avg = float64(count / float64(length))
		res = append(res, avg)
	}

	return res
}

//二叉树的所有路径，思路：非递归的形式，用深度优先遍历，使用两个栈，一个存放节点，另一个存放生成的字符串
// 用时 50分钟
func BinaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	s := common.NewStack()
	t := common.NewStack()
	s.Push(root)
	t.Push(strconv.Itoa(root.Val))
	for !s.IsEmpty() {
		node := s.Pop().(*TreeNode)
		if node == nil {
			break
		}
		temp := t.Pop().(string)
		if node.Right != nil {
			s.Push(node.Right)
			t.Push(fmt.Sprintf("%s->%d", temp, node.Right.Val))
		}
		if node.Left != nil {
			s.Push(node.Left)
			t.Push(fmt.Sprintf("%s->%d", temp, node.Left.Val))
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, temp)
		}
	}

	return res
}

func GetDepth(root *TreeNode, s string, res []string) []string {
	//if root == nil {
	//	return res
	//}
	//s = fmt.Sprintf(s+"->%d", root.Val)
	//if root.Left == nil && root.Right == nil {
	//	//println(res)
	//	res = append(res, s)
	//	return res
	//}
	//if root.Left != nil {
	//	sLeft := s
	//	res = GetDepth(root.Left, sLeft, res)
	//}
	//if root.Right != nil {
	//	sRight := s
	//	res = GetDepth(root.Right, sRight, res)
	//}

	return res
}

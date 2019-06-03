package binaryTree

import (
	"lts-demo/test/common"
	"strconv"
)

// 无厘头的题目
func Tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Right == nil && t.Left == nil {
		return strconv.Itoa(t.Val)
	}
	if t.Left != nil && t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + Tree2str(t.Left) + ")"
	} else if t.Left == nil && t.Right != nil {
		return strconv.Itoa(t.Val) + "()" + "(" + Tree2str(t.Right) + ")"
	}
	return strconv.Itoa(t.Val) + "(" + Tree2str(t.Left) + ")" + "(" + Tree2str(t.Right) + ")"
}

func RangeSumBST(root *TreeNode, L int,R int ) int {
	res := 0
	if root == nil{
		return res
	}
	s := common.NewStack()
	node := root
	for !s.IsEmpty() || node != nil{
		if node != nil{
			s.Push(node)
			node = node.Left
		}else {
			node = s.Pop().(*TreeNode)
			if node.Val >= L && node.Val <= R{
				res += node.Val
			}
			node = node.Right
		}
	}
	return res
}
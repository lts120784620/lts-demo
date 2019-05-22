package binaryTree

import "lts-demo/test/common"

// 求二叉树的路径之和是否等于sum值
func HasPathSum(root *TreeNode , sum int) bool {
	res := false
	if root == nil{
		return res
	}
	s := common.NewStack()
	s.Push(root)
	t := common.NewStack()
	var total int
	t.Push(root.Val)
	for !s.IsEmpty(){
		node := s.Pop().(*TreeNode)
		total = t.Pop().(int)

		if node.Right != nil{
			s.Push(node.Right)
			t.Push(total + node.Right.Val)
		}
		if node.Left != nil{
			s.Push(node.Left)
			t.Push(total + node.Left.Val)
		}
		if node.Left == nil && node.Right == nil{
			if total == sum{
				return true
			}
		}
	}
	return res
}

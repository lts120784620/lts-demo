package binaryTree

import "lts-demo/test/common"

// 求二叉树的路径之和是否等于sum值
func HasPathSum(root *TreeNode, sum int) bool {
	res := false
	if root == nil {
		return res
	}
	s := common.NewStack()
	s.Push(root)
	t := common.NewStack()
	var total int
	t.Push(root.Val)
	for !s.IsEmpty() {
		node := s.Pop().(*TreeNode)
		total = t.Pop().(int)

		if node.Right != nil {
			s.Push(node.Right)
			t.Push(total + node.Right.Val)
		}
		if node.Left != nil {
			s.Push(node.Left)
			t.Push(total + node.Left.Val)
		}
		if node.Left == nil && node.Right == nil {
			if total == sum {
				return true
			}
		}
	}
	return res
}

// 判断两颗二叉树的叶子节点序列是否相同(非递归)
// 思路:两个栈，分别存放两个树节点，判断叶子节点
func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res := false
	if root1 == nil || root2 == nil{
		return false
	}
	s1 := common.NewStack()
	s2 := common.NewStack()
	s1.Push(root1)
	s2.Push(root2)
	for !s1.IsEmpty() && !s2.IsEmpty(){
		node1 := s1.Pop().(*TreeNode)
		node2 := s2.Pop().(*TreeNode)
		for node1.Left != nil || node1.Right != nil{
			if node1.Right != nil{
				s1.Push(node1.Right)
			}
			if node1.Left != nil{
				s1.Push(node1.Left)
			}
			node1 = s1.Pop().(*TreeNode)
		}
		for node2.Left != nil || node2.Right != nil{
			if node2.Right != nil{
				s2.Push(node2.Right)
			}
			if node2.Left != nil{
				s2.Push(node2.Left)
			}
			node2 = s2.Pop().(*TreeNode)
		}
		if node1.Val != node2.Val{
			return false
		}else{
			res = true
		}
	}
	return res
}

//func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
//	if root1 == nil || root2 == nil{
//		return false
//	}
//	for root1.Left == nil && root1.Right == nil {
//
//	}
//	if root2.Left == nil && root2.Right == nil {
//
//	}
//	if root1.Left != nil {
//		return false
//	}
//}

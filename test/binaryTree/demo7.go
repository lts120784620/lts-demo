package binaryTree

import (
	"fmt"
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

// 计算二叉搜索树在L R之间的的节点的和
// 遍历，判断，累加
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

// 寻找二叉树中所有左叶子节点的和
// 思路：判断左节点的左右节点是否为空，做累加
func SumOfLeftLeaves(root *TreeNode) int{
	if root == nil {
		return 0
	}
	if root.Left!= nil && root.Left.Left == nil && root.Right != nil{
		fmt.Println(root.Left.Val)
		return root.Left.Val + SumOfLeftLeaves(root.Right)
	}

	return SumOfLeftLeaves(root.Left) + SumOfLeftLeaves(root.Right)
}

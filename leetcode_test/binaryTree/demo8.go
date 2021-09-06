package binaryTree

import (
	"fmt"
	"lts-demo/leetcode_test/common"
)

//func IsSymmtric(root *TreeNode) bool {
//	res := true
//	IsSymmtric(root.Left)
//	if {
//
//	}
//	IsSymmtric(root.Right)
//	return res
//}

// 非递归方法
func IsSymmtric(root *TreeNode) bool {
	res := true
	s := common.NewStack()
	node := root
	sym := []*TreeNode{}
	for !s.IsEmpty() || node != nil {
		if node != nil {
			s.Push(node)
			node = node.Left
		} else {
			node = s.Pop().(*TreeNode)
			fmt.Println(node)
			sym = append(sym, node)
			node = node.Right
		}
	}
	if len(sym)%2 == 0 {
		return false
	}
	i := 0
	j := len(sym) - 1
	for i != j {
		if sym[i] != sym[j] {
			res = false
		}
		i++
		j--
	}
	return res
}

package binaryTree

import (
	"strconv"
)

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

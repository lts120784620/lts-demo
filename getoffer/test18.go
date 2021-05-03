package main

import (
	"fmt"
	"lts-demo/getoffer/binTree"
	_ "lts-demo/getoffer/binTree"
)

func main() {
	arr_a := []interface{}{8, 8, 7, 8, 2, nil, nil, nil, nil, 4, 7}
	a := binTree.Convert2Tree(arr_a, 0)

	arr_b := []interface{}{8, 9, 2}
	b := binTree.Convert2Tree(arr_b, 0)
	fmt.Println(hasSubTree(a, b))
}

func hasSubTree(A *binTree.BinaryTreeNode, B *binTree.BinaryTreeNode) bool {
	result := false

	if A == nil || B == nil{
		return result
	}
	if A.Value == B.Value {
		result = doesTreeAHasTreeB(A, B)
		if !result {
			result = hasSubTree(A.Left, B)
		}
		if !result {
			result = hasSubTree(A.Right, B)
		}
	}
	return result
}

func doesTreeAHasTreeB(A *binTree.BinaryTreeNode, B *binTree.BinaryTreeNode) bool {
	if A == nil {
		return false
	}
	if B == nil {
		return true
	}
	if A.Value != B.Value{
		return false
	}
	return doesTreeAHasTreeB(A.Left,B.Left) && doesTreeAHasTreeB(A.Right,B.Right)
}

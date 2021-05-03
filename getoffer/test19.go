package main

import (
	"fmt"
	"lts-demo/getoffer/binTree"
)

func main() {
	arr_a := []interface{}{8, 6, 10, 5, 7, 9, 11}
	a := binTree.Convert2Tree(arr_a, 0)
	binTree.PrintBinaryTree(a)
	mirrorRecursively(a)
	fmt.Println()
	binTree.PrintBinaryTree(a)
}

func mirrorRecursively(root *binTree.BinaryTreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	if root.Left != nil {
		mirrorRecursively(root.Left)
	}
	if root.Right != nil {
		mirrorRecursively(root.Right)
	}
}


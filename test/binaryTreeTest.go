package main

import (
	"lts-demo/test/binaryTree"
)

func initTree(arr []int) *binaryTree.TreeNode {
	list := []*binaryTree.TreeNode{}
	for _, v := range arr {
		node := &binaryTree.TreeNode{Val: v, Left: nil, Right: nil}
		if v == 0 {
			node = nil
		}
		list = append(list, node)
	}
	if len(list) <= 0 {
		return nil
	}
	for i := 0; i < len(arr) / 2 - 1; i++ {
		if list[2 * i + 1] != nil {
			list[i].Left = list[2 * i + 1]
		}
		if list[2 * i + 2] != nil {
			list[i].Right = list[2 * i + 2]
		}
	}
	// 最后一个节点
	last := len(arr) / 2 - 1
	list[last].Left = list[2 * last + 1]
	// 奇数，有右节点
	if len(arr) % 2 == 1 {
		list[last].Right = list[2 * last + 2]
	}
	return list[0]
}

func main() {
	initTree([]int{3, 9, 20, 0, 0, 15, 7})
}

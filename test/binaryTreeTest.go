package main

import (
	"lts-demo/test/binaryTree"
	"fmt"
)

func initTree(data []interface{}) *binaryTree.TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := &binaryTree.TreeNode{}
	switch t := data[0].(type) {
	case int:
		root.Val = t
	case nil:
		return nil
	default:
		panic("Unknown element type")
	}

	queue := make([]*binaryTree.TreeNode, 1)
	queue[0] = root

	data = data[1:]
	for len(data) > 0 && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 左侧节点
		node.Left = newNodeFromData(data[0])
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		data = data[1:]

		// 右侧节点
		if len(data) > 0 {
			node.Right = newNodeFromData(data[0])
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			data = data[1:]
		}
	}
	return root
}

// 根据数据来创建一个新的节点
func newNodeFromData(val interface{}) *binaryTree.TreeNode {
	switch t := val.(type) {
	case int:
		return &binaryTree.TreeNode{Val: t}
	case nil:
		return nil
	default:
		panic("Unknown element type")
	}
}
func travelse(root *binaryTree.TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, "->")
	travelse(root.Left)
	travelse(root.Right)
}

func main() {
	//root := initTree([]interface{}{3, 9, 20, 0, 0, 15, 7})
	//fmt.Println(binaryTree.MaxDepth(root))

	//travelse(initTree([]interface{}{4,2,7,1,3,6,9}))
	//fmt.Println()
	//travelse(binaryTree.InvertTree(initTree([]interface{}{4,2,7,1,3,6,9})))

	//fmt.Println(binaryTree.IsUnivalTree(initTree([]interface{}{2,2,2,5,2})))

	//fmt.Println(binaryTree.SearchBST(initTree([]interface{}{4,2,7,1,3}),2))

	//travelse(binaryTree.TrimBST(initTree([]interface{}{4,2,5,1,3}),2,3))

	//travelse(binaryTree.ConvertBST(initTree([]interface{}{2,1,3})))

	//fmt.Println(binaryTree.FindMode(initTree([]interface{}{1,nil,2,2})))
	//fmt.Println(binaryTree.FindMode(initTree([]interface{}{2,2})))

	//fmt.Println(binaryTree.LevelOrderBottom(initTree([]interface{}{1,nil,2,2})))

	//fmt.Println(binaryTree.LevelOrderBottom2(initTree([]interface{}{3,9,20,nil,nil,15,17})))

	//fmt.Println(binaryTree.InorderTraversal(initTree([]interface{}{3,9,20,nil,nil,15,17})))

	res := []string{}
	binaryTree.GetDepth(initTree([]interface{}{1, 2, 3, nil, 5}), "", res)
	for _,j:= range res {
		fmt.Println(j)
	}

	fmt.Println()
}

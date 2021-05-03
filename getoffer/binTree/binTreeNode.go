package binTree

import (
	"fmt"
	"lts-demo/test/common"
)

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Convert2Tree(arr []interface{}, idx int) *BinaryTreeNode {
	if arr == nil {
		return nil
	}
	length := len(arr)
	if idx >= length {
		return nil
	}

	if arr[idx] == nil {
		return nil
	}
	root := &BinaryTreeNode{Value: arr[idx].(int)}
	root.Left = Convert2Tree(arr, 2*idx+1)
	root.Right = Convert2Tree(arr, 2*idx+2)

	return root
}

func PrintBinaryTree(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ->", root.Value)
	PrintBinaryTree(root.Left)
	PrintBinaryTree(root.Right)
}

func PrintBinaryTree2(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	stack := common.NewStack()
	stack.Push(root)
	for !stack.IsEmpty(){
		node := stack.Pop().(*BinaryTreeNode)
		if node == nil{
			continue
		}
		fmt.Printf("%d ->", node.Value)
		stack.Push(node.Right)
		stack.Push(node.Left)
	}

}

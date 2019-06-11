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

	//fmt.Println(binaryTree.BinaryTreePaths(initTree([]interface{}{1,2,3,nil,5})))

	//fmt.Println(binaryTree.HasPathSum(initTree([]interface{}{1}),3))

	//root1 := initTree([]interface{}{44,79,25,nil,nil,112,7,74,49,2,122})
	//root2 := initTree([]interface{}{38,86,120,49,54,2,122,nil,nil,74,79})
	//fmt.Println(binaryTree.LeafSimilar(root1, root2))

	//fmt.Println(binaryTree.IncresingBST(initTree([]interface{}{5,3,6,2,4,nil,8,1,nil,nil,nil,7,9})))

	//fmt.Println(binaryTree.GetMinimumDifference(initTree([]interface{}{5,4,7})))

	//root1 := initTree([]interface{}{1, 2})
	//root2 := initTree([]interface{}{1, nil, 2})
	//fmt.Println(binaryTree.IsSameTree(root1, root2))

	//fmt.Println(binaryTree.MinDiffInBST(initTree([]interface{}{27,nil,34,nil,58,50,nil,44,nil,nil,nil})))

	//fmt.Println(binaryTree.Tree2str(initTree([]interface{}{1,2,3,nil,4})))

	//fmt.Println(binaryTree.RangeSumBST(initTree([]interface{}{10,5,15,3,7,nil,18}),7,15))

	//fmt.Println(binaryTree.SumOfLeftLeaves(initTree([]interface{}{3,9,20,nil,nil,15,7})))

	//fmt.Println(binaryTree.PathSum(initTree([]interface{}{10,5,-3,3,2,nil,11,3,-2,nil,1}),8))

	//fmt.Println(binaryTree.FindTilt(initTree([]interface{}{1,2,3,4,nil,5})))

	fmt.Println(binaryTree.IsBalanced(initTree([]interface{}{1,2,2,3,3,nil,nil,4,4})))
	fmt.Println()
}

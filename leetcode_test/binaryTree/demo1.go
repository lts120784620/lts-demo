package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
求给定二叉树的最大深度
*/
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

/**
合并两个二叉树，若相同位置节点都有值则相加，若只有一个有值则作为新节点
*/
func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var t *TreeNode
	if t1 == nil {
		t = t2
	} else if t2 == nil {
		t = t1
	} else {
		t = &TreeNode{t1.Val + t2.Val, nil, nil}
		t.Left = MergeTrees(t1.Left, t2.Left)
		t.Right = MergeTrees(t1.Right, t2.Right)
	}
	return t
}

/**
翻转二叉树
*/
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := InvertTree(root.Left)
	right := InvertTree(root.Right)
	// exchange
	root.Left = right
	root.Right = left
	return root
}

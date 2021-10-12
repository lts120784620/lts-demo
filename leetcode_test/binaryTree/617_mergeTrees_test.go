package binaryTree

import "testing"

/**
No.617
描述：
	给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
	你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
思路：
    1、递归的思路，两个指针分别在两棵树上遍历
时间：
    1、2021/10/1
*/

func TestMergeTrees(t *testing.T) {
	p := NewBinTreeByArrays([]interface{}{1, 3, 2, 5})
	q := NewBinTreeByArrays([]interface{}{2, 1, 3, nil, 4, nil, 7})
	mergeTrees(p, q)
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}

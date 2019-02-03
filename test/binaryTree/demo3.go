package binaryTree


/**
判断是否为单值二叉树（所有节点值都相同的为单值二叉树）
 */
func IsUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && (root.Val != root.Left.Val) {
		return false
	}
	if root.Right != nil && (root.Val != root.Right.Val) {
		return false
	}

	return IsUnivalTree(root.Left) && IsUnivalTree(root.Right)
}

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

/**
找到BST中等于给定值的节点
 */
func SearchBST(root *TreeNode, val int) *TreeNode {
	if (root == nil) {
		return nil
	}
	if (root.Val == val) {
		return root
	}
	if (val > root.Val) {
		return SearchBST(root.Right, val)
	} else {
		return SearchBST(root.Left, val)
	}
}

/**
有序数组转换二叉搜索树
 */
func sortedArrayToBST(num []int) *TreeNode {
	if (len(num) == 0) {
		return nil
	}
	return arrayToBST(num, 0, len(num))
}
func arrayToBST(num []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end ) / 2
	left := arrayToBST(num, start, mid - 1)
	right := arrayToBST(num, mid + 1, end)
	return &TreeNode{num[mid], left, right}
}

/**
根据左右边界裁剪二叉树
 */
func TrimBST(root *TreeNode, L, R int) *TreeNode {
	if root == nil || L > R {
		return nil
	}
	if root.Val < L {
		return TrimBST(root.Right, L, R)
	}
	if root.Val > R{
		return TrimBST(root.Left, L, R)
	}
	root.Left = TrimBST(root.Left, L, R)
	root.Right = TrimBST(root.Right, L, R)
	return root
}

func lowestCommonAncestor(root , p ,q *TreeNode) *TreeNode {
	if root == nil {

	}
}

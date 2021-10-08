package binaryTree

import "testing"

/**
No.108 将有序数组转换为二叉搜索树
描述：
	给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
	高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
思路：
    1、给定有序数组，利用二叉搜索树中序遍历即为递增序列的特点，直接从中点还原
	ps：要注意树的高度差为一
时间：
    1、2021/9/30
*/

func TestSortedArrayToBST(t *testing.T) {
	// 0 - 5 2
	// 0 - 1 2 3 - 5
	//
	sortedArrayToBST([]int{1, 2, 3, 4, 5, 6})
}

func sortedArrayToBST(nums []int) *TreeNode {
	return middleNode(nums, 0, len(nums)-1)
}

func middleNode(nums []int, left, right int) *TreeNode {
	if left >= 0 && right < len(nums) && left > right {
		return nil
	}
	// 中间位置
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = middleNode(nums, left, mid-1)
	root.Right = middleNode(nums, mid+1, right)
	return root
}

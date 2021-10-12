package binaryTree

import (
	"fmt"
	"testing"
)

/**
No.654. 最大二叉树
描述：
	给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
	二叉树的根是数组 nums 中的最大元素。
	左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
	右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
	返回有给定数组 nums 构建的 最大二叉树 。
思路：
    1、
时间：
    1、2021/10/11
*/

func TestConstructMaximumBinaryTree(t *testing.T) {
	r := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(r.String())
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return reverseMaxNum(nums, 0, len(nums)-1)
}

func reverseMaxNum(nums []int, left, right int) *TreeNode {
	if left > right || len(nums) == 0 {
		return nil
	}

	// 找到最大值的位置
	maxIdx := left
	max := nums[left]
	for i := left; i <= right; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}

	root := &TreeNode{Val: nums[maxIdx]}

	root.Left = reverseMaxNum(nums, left, maxIdx-1)
	root.Right = reverseMaxNum(nums, maxIdx+1, right)
	return root
}

package array

import (
	"fmt"
	"testing"
)

/**
No.34. 在排序数组中查找元素的第一个和最后一个位置
描述：
	给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
	如果数组中不存在目标值 target，返回[-1, -1]。
	进阶：
		你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？
思路：
    1、
时间：
    1、2021/10/12
*/

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func searchRange(nums []int, target int) []int {

	return []int{searchLeftSide(nums, target), searchRightSide(nums, target)}
}

func searchLeftSide(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			// 向左飘
			right = mid - 1
		} else if target > nums[mid] {
			// 在右边，left向右移
			left = mid + 1
		} else if target < nums[mid] {
			// 在左边，right向左移
			right = mid - 1
		}
	}
	if left > len(nums)-1 || nums[left] != target {
		return -1
	}
	return left
}

func searchRightSide(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			// 向左飘
			left = mid + 1
		} else if target > nums[mid] {
			// 在右边，left向右移
			left = mid + 1
		} else if target < nums[mid] {
			// 在左边，right向左移
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

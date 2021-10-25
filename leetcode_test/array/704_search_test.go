package array

import (
	"fmt"
	"testing"
)

/**
No.704. 二分查找
描述：
	给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
思路：
    1、二分查找
时间：
    1、2021/10/12
*/

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}

	return -1
}

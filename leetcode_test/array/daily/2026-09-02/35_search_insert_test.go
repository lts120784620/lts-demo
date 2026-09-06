package daily20260902

import "testing"

/*
LeetCode 35. 搜索插入位置（简单）
链接：https://leetcode.cn/problems/search-insert-position/

题目要求：
给定一个无重复元素、按升序排列的整数数组 nums 和目标值 target：
 1. 如果 target 存在，返回它的下标。
 2. 如果 target 不存在，返回将它按顺序插入数组时应处于的下标。

示例：

	nums = [1,3,5,6], target = 5，返回 2
	nums = [1,3,5,6], target = 2，返回 1
	nums = [1,3,5,6], target = 7，返回 4

约束：

	1 <= len(nums) <= 10^4
	-10^4 <= nums[i] <= 10^4
	nums 中的元素互不相同，并按升序排列
	-10^4 <= target <= 10^4

复杂度要求：时间 O(log n)。
今日训练点：统一二分搜索区间的定义，并正确处理 target 位于数组两端之外的情况。
建议限时：20 分钟。
提示触发点：独立思考 10 分钟后，仍无法解释循环结束时 left 或 right 的含义，
可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 搜索区间定义：
  - 循环结束时返回值为什么是插入位置：
*/
func searchInsert(nums []int, target int) int {
	// TODO: 在这里实现。
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "found", nums: []int{1, 3, 5, 6}, target: 5, want: 2},
		{name: "insert-middle", nums: []int{1, 3, 5, 6}, target: 2, want: 1},
		{name: "insert-after-end", nums: []int{1, 3, 5, 6}, target: 7, want: 4},
		{name: "insert-before-start", nums: []int{1, 3, 5, 6}, target: 0, want: 0},
		{name: "single-before", nums: []int{1}, target: 0, want: 0},
		{name: "single-after", nums: []int{1}, target: 2, want: 1},
		{name: "test1", nums: []int{1, 3}, target: 2, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.nums, tt.target); got != tt.want {
				t.Fatalf("searchInsert(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

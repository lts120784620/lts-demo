package daily20260912

import (
	"testing"
)

/*
LeetCode 704. 二分查找（简单）
链接：https://leetcode.cn/problems/binary-search/

题目要求：
给定一个按升序排列、元素互不重复的整数数组 nums 和目标值 target。
如果 target 存在，返回它的下标；否则返回 -1。

示例：

	nums = [-1,0,3,5,9,12], target = 9，返回 4
	nums = [-1,0,3,5,9,12], target = 2，返回 -1

约束：

	1 <= len(nums) <= 10^4
	-10^4 < nums[i], target < 10^4
	nums 中的元素互不重复
	nums 按升序排列

复杂度目标：时间 O(log n)，额外空间 O(1)。
今日训练点：固定一种二分区间定义，并让循环条件和左右边界更新保持一致。
建议限时：20 分钟。
提示触发点：独立思考 10 分钟后，仍然需要通过调试决定使用 left < right
还是 left <= right 时，可索取一级提示。

完成后请补充：
  - 搜索区间的定义：
  - 循环继续的条件：
  - left 和 right 的更新规则：
  - 时间复杂度：
  - 空间复杂度：
*/
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "middle", nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{name: "not-found", nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
		{name: "first", nums: []int{1, 3, 5, 7}, target: 1, want: 0},
		{name: "last", nums: []int{1, 3, 5, 7}, target: 7, want: 3},
		{name: "single-found", nums: []int{6}, target: 6, want: 0},
		{name: "single-not-found", nums: []int{6}, target: 5, want: -1},
		{name: "smaller-than-all", nums: []int{2, 4, 6}, target: 1, want: -1},
		{name: "greater-than-all", nums: []int{2, 4, 6}, target: 8, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.nums, tt.target); got != tt.want {
				t.Fatalf("binarySearch(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

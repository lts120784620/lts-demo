package daily20260906

import "testing"

/*
LeetCode 209. 长度最小的子数组（复习）
链接：https://leetcode.cn/problems/minimum-size-subarray-sum/

题目要求：
给定一个由正整数组成的数组 nums 和正整数 target，找出元素和大于或等于
target 的最短连续子数组，返回它的长度；如果不存在，返回 0。

示例：

	target = 7, nums = [2,3,1,2,4,3]，返回 2
	因为 [4,3] 是满足条件的最短连续子数组。

约束：

	1 <= target <= 10^9
	1 <= len(nums) <= 10^5
	1 <= nums[i] <= 10^4

复杂度目标：时间 O(n)，额外空间 O(1)。
今日训练方式：不打开旧代码，从记忆重写并解释窗口不变量。
建议限时：20 分钟。
提示触发点：独立思考 10 分钟后，仍无法说明左右边界分别在什么条件下移动时，
可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 窗口不变量：
  - 为什么 nums[i] 为正数很重要：
  - 本次掌握度（0～3）：
*/
func minSubArrayLenReview(target int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	min := len(nums) + 1
	for ; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < min {
				min = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if right == len(nums) && min == len(nums)+1 {
		return 0
	}
	return min
}

func TestMinSubArrayLenReview(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{name: "basic", target: 7, nums: []int{2, 3, 1, 2, 4, 3}, want: 2},
		{name: "single-answer", target: 4, nums: []int{1, 4, 4}, want: 1},
		{name: "no-answer", target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, want: 0},
		{name: "whole-array", target: 15, nums: []int{1, 2, 3, 4, 5}, want: 5},
		{name: "single-element", target: 1, nums: []int{1}, want: 1},
		{name: "shrink-more-than-once", target: 7, nums: []int{3, 1, 1, 1, 5}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenReview(tt.target, tt.nums); got != tt.want {
				t.Fatalf("minSubArrayLenReview(%d, %v) = %d, want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}

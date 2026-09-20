package daily20260920

import "testing"

/*
LeetCode 209. 长度最小的子数组（中等）
链接：https://leetcode.cn/problems/minimum-size-subarray-sum/

题目要求：
给定正整数数组 nums 和正整数 target，返回元素和大于等于 target 的最短连续子数组长度；
不存在时返回 0。

复杂度目标：时间 O(n)，额外空间 O(1)。
复习要求：使用可变长度滑动窗口，明确左右边界含义和 sum 对应的窗口。
建议限时：20 分钟。
提示触发点：独立思考 12 分钟后，仍然需要枚举每个起点和终点，或无法确定何时收缩
左边界，可索取一级提示。

完成评价（2026-09-20）：
  - 测试结果：通过（6 个测试用例）
  - left、right 分别表示什么：left、right 分别是当前窗口的左右端点
  - sum 对应哪个窗口：加入 nums[right] 后，sum 表示闭区间 [left, right] 的元素和
  - 窗口在什么条件下持续收缩：sum >= target 时持续移动 left
  - 为什么每个元素最多进出窗口一次：left、right 都只单向递增，不会回退
  - 时间复杂度及原因：O(n)，每个元素最多被 right 加入一次、被 left 移出一次
  - 空间复杂度：O(1)
  - 核心模式与不变量：扩张后先记录所有合法窗口，再收缩到 sum < target
  - 提示使用：无
  - 主要错误类型：无
  - 本次掌握度（0～3）：2
  - 下次复习日期：2026-09-23
  - 备注：本次独立完成到期复习并通过全部测试；耗时与复杂度口述未记录
*/
func minSubArrayLenDailyReview(target int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	min := len(nums) + 1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < min {
				min = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}

func TestMinSubArrayLenDailyReview(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{name: "basic", target: 7, nums: []int{2, 3, 1, 2, 4, 3}, want: 2},
		{name: "exact-single", target: 4, nums: []int{1, 4, 4}, want: 1},
		{name: "whole-array", target: 11, nums: []int{1, 2, 3, 4, 5}, want: 3},
		{name: "no-answer", target: 20, nums: []int{1, 2, 3, 4, 5}, want: 0},
		{name: "single-match", target: 5, nums: []int{5}, want: 1},
		{name: "repeated-shrink", target: 15, nums: []int{1, 2, 3, 4, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenDailyReview(tt.target, tt.nums); got != tt.want {
				t.Fatalf("minSubArrayLenDailyReview(%d, %v) = %d, want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}

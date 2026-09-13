package daily20260913

import "testing"

/*
LeetCode 209. 长度最小的子数组（数组毕业复习）
链接：https://leetcode.cn/problems/minimum-size-subarray-sum/

题目要求：
给定一个由正整数组成的数组 nums 和正整数 target，找出元素和大于或等于
target 的最短连续子数组，返回它的长度；如果不存在，返回 0。

示例：

	target = 7, nums = [2,3,1,2,4,3]，返回 2

约束：

	1 <= target <= 10^9
	1 <= len(nums) <= 10^5
	1 <= nums[i] <= 10^4

复杂度目标：时间 O(n)，额外空间 O(1)。
复习要求：不打开旧实现，从窗口不变量开始独立重写。
建议限时：25 分钟。
提示触发点：独立思考 15 分钟后仍无法确定左右边界的移动条件时，可索取一级提示。

完成评价（2026-09-13）：
  - 测试结果：通过（6 个测试用例）
  - 时间复杂度：O(n)，左右指针都只向右移动，每个元素最多进入和离开窗口一次
  - 空间复杂度：O(1)
  - 核心模式与不变量：滑动窗口；判断候选长度时，sum 等于闭区间 [left,right] 的元素和
  - 提示使用：三级
  - 主要错误类型：实现、边界
  - 掌握度：1/3
  - 下次复习日期：2026-09-15
  - 备注：能够识别滑动窗口主干；需要巩固“先记录有效窗口再收缩”以及使用合法范围外哨兵。耗时未记录
*/
func minSubArrayLenFinalReview(target int, nums []int) int {
	// 核心点是双指针构建的滑动窗口是吧，sum值，如果大于则
	// 0-i 表示遍历过的，减去的值，i-j表示sum的连续累加值，j-max 表示待移动的区间
	left := 0
	right := 0
	sum := 0
	res := len(nums) + 1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < res {
				res = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func TestMinSubArrayLenFinalReview(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{name: "basic", target: 7, nums: []int{2, 3, 1, 2, 4, 3}, want: 2},
		{name: "single-answer", target: 5, nums: []int{1, 5, 1}, want: 1},
		{name: "no-answer", target: 20, nums: []int{2, 3, 4}, want: 0},
		{name: "whole-array", target: 15, nums: []int{1, 2, 3, 4, 5}, want: 5},
		{name: "shrink-repeatedly", target: 8, nums: []int{2, 3, 1, 2, 4, 3}, want: 3},
		{name: "one-large-element", target: 6, nums: []int{10}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenFinalReview(tt.target, tt.nums); got != tt.want {
				t.Fatalf("minSubArrayLenFinalReview(%d, %v) = %d, want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}

package daily20260915

import "testing"

/*
LeetCode 209. 长度最小的子数组（中等）
链接：https://leetcode.cn/problems/minimum-size-subarray-sum/

题目要求：
给定一个含有 n 个正整数的数组 nums 和一个正整数 target，找出该数组中满足
元素和大于等于 target 的长度最小的连续子数组，并返回其长度；不存在时返回 0。

复杂度目标：时间 O(n)，额外空间 O(1)。
复习要求：不打开旧实现，从窗口不变量开始独立重写。
建议限时：20 分钟。
提示触发点：独立思考 12 分钟后，仍无法确定什么时候记录答案、什么时候移动 left，
可索取一级提示。

完成评价（2026-09-18）：
  - 测试结果：通过（5 个测试用例）
  - 窗口包含的下标范围：闭区间 [left,right]
  - 什么时候窗口有效：sum >= target
  - 记录答案与收缩窗口的先后顺序：先记录当前合法窗口长度，再移除 nums[left] 并移动 left
  - 时间复杂度及原因：O(n)，左右指针均只向右移动，每个元素最多进入和离开窗口一次
  - 空间复杂度：O(1)
  - 核心模式与不变量：加入 nums[right] 后，sum 等于闭区间 [left,right] 的元素和
  - 提示使用：三级
  - 主要错误类型：实现、边界
  - 本次掌握度（0～3）：1
  - 下次复习日期：2026-09-20
  - 备注：初版混淆了 right 作为边界和数组下标的含义，并遗漏 sum == target 的合法窗口；耗时未记录
*/
func minSubArrayLenRecheck(target int, nums []int) int {
	// 滑动数组的思路
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

func TestMinSubArrayLenRecheck(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{name: "basic", target: 7, nums: []int{2, 3, 1, 2, 4, 3}, want: 2},
		{name: "single-element-window", target: 4, nums: []int{1, 4, 4}, want: 1},
		{name: "no-valid-window", target: 10, nums: []int{1, 1, 1, 1}, want: 0},
		{name: "whole-array", target: 15, nums: []int{1, 2, 3, 4, 5}, want: 5},
		{name: "repeated-shrink", target: 11, nums: []int{1, 2, 3, 4, 5}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenRecheck(tt.target, tt.nums); got != tt.want {
				t.Fatalf("minSubArrayLenRecheck(%d, %v) = %d, want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}

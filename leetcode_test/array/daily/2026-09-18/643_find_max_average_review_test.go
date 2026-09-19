package daily20260918

import (
	"math"
	"testing"
)

/*
LeetCode 643. 子数组最大平均数 I（简单）
链接：https://leetcode.cn/problems/maximum-average-subarray-i/

题目要求：
给定整数数组 nums 和整数 k，找出长度恰好为 k 的连续子数组中的最大平均数。

复杂度目标：时间 O(n)，额外空间 O(1)。
复习要求：不打开旧实现，独立写出固定长度滑动窗口。
建议限时：15 分钟。
提示触发点：独立思考 10 分钟后，仍无法确定第一个窗口如何初始化，或窗口移动时
哪个元素进入、哪个元素离开，可索取一级提示。

完成评价（2026-09-19）：
  - 测试结果：通过（5 个测试用例）
  - 当前窗口包含的下标范围：闭区间 [left,right]
  - 窗口移动时进入与离开的元素：nums[right] 进入，窗口达到 k 后 nums[left] 离开
  - 为什么可以只比较窗口和：所有候选窗口长度都为 k，分母相同
  - 时间复杂度及原因：O(n)，每个元素最多进入和离开窗口一次
  - 空间复杂度：O(1)
  - 核心模式与不变量：记录答案时，sum 表示长度恰好为 k 的窗口元素和
  - 提示使用：无
  - 主要错误类型：无
  - 本次掌握度（0～3）：2
  - 下次复习日期：2026-09-22
  - 备注：独立完成到期复习并通过全部测试；耗时未记录
*/
func findMaxAverageReview(nums []int, k int) float64 {
	res := float64(math.MinInt64)
	sum := 0.0
	left := 0
	right := 0
	for right < len(nums) {
		sum += float64(nums[right])
		if right-left+1 == k {
			if sum > res {
				res = sum
			}
			sum -= float64(nums[left])
			left++
		}
		right++
	}
	avg := res / float64(right-left+1)
	return avg
}

func TestFindMaxAverageReview(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{name: "basic", nums: []int{1, 12, -5, -6, 50, 3}, k: 4, want: 12.75},
		{name: "all-negative", nums: []int{-1, -12, -5, -6, -50, -3}, k: 2, want: -5.5},
		{name: "whole-array", nums: []int{1, 2, 3}, k: 3, want: 2},
		{name: "window-size-one", nums: []int{-3, 4, -2}, k: 1, want: 4},
		{name: "single-negative", nums: []int{-7}, k: 1, want: -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxAverageReview(tt.nums, tt.k)
			if math.IsNaN(got) || math.Abs(got-tt.want) > 1e-5 {
				t.Fatalf("findMaxAverageReview(%v, %d) = %f, want %f", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

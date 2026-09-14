package daily20260914

import (
	"math"
	"testing"
)

/*
LeetCode 643. 子数组最大平均数 I（简单）
链接：https://leetcode.cn/problems/maximum-average-subarray-i/

题目要求：
给定一个长度为 n 的整数数组 nums 和整数 k，找出长度恰好为 k 的连续子数组中
最大的平均数，并返回该平均数。答案误差小于 10^-5 即视为正确。

示例：

	nums = [1,12,-5,-6,50,3], k = 4
	返回 12.75，因为 [12,-5,-6,50] 的和为 51，平均数为 51/4

约束：

	1 <= k <= len(nums) <= 10^5
	-10^4 <= nums[i] <= 10^4

复杂度目标：时间 O(n)，额外空间 O(1)。
今日训练点：固定长度滑动窗口；窗口移动前后，先明确有效窗口和更新答案的时机。
建议限时：20 分钟。
提示触发点：独立思考 10 分钟后，每移动一次窗口仍需要重新遍历 k 个元素求和时，
可索取一级提示。

完成评价（2026-09-15）：
  - 测试结果：本地完整测试通过
  - 窗口包含的下标范围：[right-k+1, right]
  - 窗口移动一步时进入和离开的元素：nums[right] 进入，nums[right-k] 离开
  - 时间复杂度：O(n)
  - 空间复杂度：O(1)
  - 核心模式与不变量：固定长度滑动窗口；sum 始终表示当前长度为 k 的窗口元素和
  - 提示使用：四级
  - 主要错误类型：实现、边界
  - 本次掌握度（0～3）：1
  - 下次复习日期：2026-09-17
  - 备注：初版因 0/0 产生 NaN，且在计算平均值前过早移除了窗口左端元素
*/
func findMaxAverage(nums []int, k int) float64 {
	// 先计算第一个长度为 k 的窗口
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	// 窗口向右滑动：加入一个，移除一个
	for right := k; right < len(nums); right++ {
		sum += nums[right]
		sum -= nums[right-k]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{name: "basic", nums: []int{1, 12, -5, -6, 50, 3}, k: 4, want: 12.75},
		{name: "single", nums: []int{5}, k: 1, want: 5},
		{name: "all-negative", nums: []int{-1, -12, -5, -6, -50, -3}, k: 2, want: -5.5},
		{name: "whole-array", nums: []int{1, 2, 3}, k: 3, want: 2},
		{name: "window-size-one", nums: []int{-3, 4, -2}, k: 1, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxAverage(tt.nums, tt.k)
			if math.IsNaN(got) || math.Abs(got-tt.want) > 1e-5 {
				t.Fatalf("findMaxAverage(%v, %d) = %f, want %f", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

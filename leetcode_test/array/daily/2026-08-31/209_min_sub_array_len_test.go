package daily20260831

import "testing"

/*
LeetCode 209. 长度最小的子数组（中等）
链接：https://leetcode.cn/problems/minimum-size-subarray-sum/

题目要求：
给定一个由 n 个正整数组成的数组 nums 和一个正整数 target，找出元素和大于
或等于 target 的最短连续子数组，并返回它的长度。如果不存在符合条件的连续
子数组，返回 0。

注意：“子数组”必须是原数组中连续的一段，不能跳过中间元素。

示例：

	输入：target = 7, nums = [2,3,1,2,4,3]
	输出：2
	解释：[4,3] 是满足条件的最短连续子数组。

约束：

	1 <= target <= 10^9
	1 <= len(nums) <= 10^5
	1 <= nums[i] <= 10^4

复杂度目标：先写出正确方案，再尝试将时间复杂度优化到 O(n)。

今日训练点：连续区间，以及如何避免为每个区间重复求和。
建议限时：30 分钟。
提示触发点：独立思考 15 分钟后只有枚举所有区间的方案，或无法继续优化时，
可索取一级提示。

完成后请在这里补充：
  - 时间复杂度：
  - 空间复杂度：
  - 关键不变量：
  - 为什么“所有 nums[i] 都是正数”很重要：
*/
func minSubArrayLen(target int, nums []int) int {
	i := 0 // 代表左侧的第一个值，减去前一个值
	j := i // 代表最右侧的值，加上后一个值
	min := len(nums) + 1
	if len(nums) == 0 {
		return 0
	}
	sum := nums[i]
	for j < len(nums) && i <= j {
		if sum < target {
			if j+1 >= len(nums) {
				break
			}
			j++
			sum += nums[j]
		} else {
			k := j - i + 1
			if k < min {
				// 如果比最大长度小，那就赋值
				min = k
			}

			sum -= nums[i]
			i++
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}

func TestMinSubArrayLen(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Fatalf("minSubArrayLen(%d, %v) = %d, want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}

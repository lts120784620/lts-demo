package daily20260901

import "testing"

/*
LeetCode 724. 寻找数组的中心下标（简单）
链接：https://leetcode.cn/problems/find-pivot-index/

题目要求：
给定整数数组 nums，寻找一个中心下标，使它左侧所有元素之和等于右侧所有元素
之和。中心下标自身不计入任何一侧。

如果中心下标位于最左端，左侧之和视为 0；位于最右端时同理。如果存在多个
中心下标，返回最靠左的一个；如果不存在，返回 -1。

示例：

	输入：nums = [1,7,3,6,5,6]
	输出：3
	解释：下标 3 左右两侧的元素和都为 11。

约束：

	1 <= len(nums) <= 10^4
	-1000 <= nums[i] <= 1000

复杂度目标：时间 O(n)，额外空间 O(1)。
今日训练点：维护已处理部分的和，并避免为每个下标重复计算整个区间。
建议限时：20 分钟。
提示触发点：独立思考 10 分钟后仍需要为每个下标重新遍历左右两侧时，
可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 关键不变量：
  - 为什么中心元素不属于左侧或右侧：
*/
func pivotIndex(nums []int) int {
	// TODO: 在这里实现。
	total := 0 // 计算一遍总和
	for i := range nums {
		total += nums[i]
	}
	preSum := 0 // 前n个数的总和
	i := 0
	has := false
	for ; i < len(nums); i++ {
		re := total - preSum - nums[i]
		if re != preSum {
			preSum += nums[i]
		} else {
			has = true
			break
		}
	}
	if !has {
		return -1
	}
	return i
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "basic", nums: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{name: "not-found", nums: []int{1, 2, 3}, want: -1},
		{name: "left-edge", nums: []int{2, 1, -1}, want: 0},
		{name: "right-edge", nums: []int{-1, 1, 2}, want: 2},
		{name: "first-of-many", nums: []int{0, 0, 0}, want: 0},
		{name: "single", nums: []int{5}, want: 0},
		{name: "test1", nums: []int{-1, -1, -1, -1, -1, 0}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.nums); got != tt.want {
				t.Fatalf("pivotIndex(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

package daily20260831

import (
	"lts-demo/leetcode_test/unit"
	"reflect"
	"testing"
)

/*
LeetCode 283. 移动零（复习）
链接：https://leetcode.cn/problems/move-zeroes/

题目要求：
给定整数数组 nums，将所有 0 原地移动到数组末尾，同时保持非零元素的相对顺序。
不能复制整个数组，函数直接修改 nums，不返回新数组。

示例：

	输入：nums = [0,1,0,3,12]
	修改后：nums = [1,3,12,0,0]

约束：

	1 <= len(nums) <= 10^4
	-2^31 <= nums[i] <= 2^31 - 1

今日训练方式：关闭 2026-08-27 的代码，从记忆重新实现并解释关键不变量。
建议限时：12 分钟。
提示触发点：8 分钟仍无法确定两个位置分别代表什么时，可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 关键不变量：
  - 与第一次实现相比，这次哪里更简洁：
*/
func moveZeroesReview(nums []int) {
	i := 0     // 代表0值的位置
	j := i + 1 // 代表i后面第一个非0
	for j < len(nums) && i < j {
		if nums[i] != 0 {
			i++
		} else if nums[j] == 0 {
			j++
		} else {
			unit.ArraysSwap(nums, i, j)
			i++
			j++
		}
	}
}

func TestMoveZeroesReview(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "basic", nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{name: "single-zero", nums: []int{0}, want: []int{0}},
		{name: "no-zero", nums: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "all-zero", nums: []int{0, 0, 0}, want: []int{0, 0, 0}},
		{name: "keep-order", nums: []int{0, -1, 0, 2, 0, -3}, want: []int{-1, 2, -3, 0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroesReview(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Fatalf("moveZeroesReview() changed nums to %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

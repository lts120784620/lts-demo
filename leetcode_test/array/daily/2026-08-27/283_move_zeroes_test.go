package daily20260827

import (
	"lts-demo/leetcode_test/unit"
	"reflect"
	"testing"
)

/*
LeetCode 283. 移动零（简单）
链接：https://leetcode.cn/problems/move-zeroes/

题目要求：
给定一个整数数组 nums，将数组中的所有 0 移动到末尾，同时保持所有非零元素
原有的相对顺序。

完成后必须满足：
 1. 直接修改传入的 nums，不返回新数组。
 2. 不复制整个数组，使用 O(1) 额外空间。
 3. 非零元素之间的相对顺序不能改变。

示例：

	输入：nums = [0,1,0,3,12]
	修改后：nums = [1,3,12,0,0]

约束：

	1 <= len(nums) <= 10^4
	-2^31 <= nums[i] <= 2^31 - 1

进阶：尽量减少对数组元素的写入或交换次数。

今日训练点：原地修改，并在移动元素时维持相对顺序。
建议限时：20 分钟。
提示触发点：独立思考 10 分钟后仍只能想到额外创建数组时，可索取一级提示。

完成后请在这里补充：
  - 时间复杂度：
  - 空间复杂度：
  - 关键不变量：
*/
func moveZeroes(nums []int) {
	// TODO: 在这里实现。函数需要直接修改 nums。
	i := 0
	j := i + 1
	for i < len(nums) && j < len(nums) {
		if nums[i] != 0 && i < j {
			i++
		} else if nums[j] == 0 {
			j++
		} else {
			// 交换
			unit.ArraysSwap(nums, i, j)
			i++
			j++
		}
	}
}

func TestMoveZeroes(t *testing.T) {
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
		{name: "", nums: []int{1, 0, 1}, want: []int{1, 1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Fatalf("moveZeroes() changed nums to %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

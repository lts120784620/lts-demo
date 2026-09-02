package daily20260901

import (
	"fmt"
	"testing"
)

/*
LeetCode 303. 区域和检索 - 数组不可变（简单）
链接：https://leetcode.cn/problems/range-sum-query-immutable/

题目要求：
给定一个不会再修改的整数数组 nums，实现：
 1. Constructor(nums)，使用 nums 初始化 NumArray。
 2. SumRange(left, right)，返回闭区间 [left, right] 内所有元素的和。

同一个 NumArray 会执行多次 SumRange 查询。

示例：

	nums = [-2,0,3,-5,2,-1]
	SumRange(0, 2) 返回 1
	SumRange(2, 5) 返回 -1
	SumRange(0, 5) 返回 -3

约束：

	1 <= len(nums) <= 10^4
	-10^5 <= nums[i] <= 10^5
	0 <= left <= right < len(nums)
	最多调用 10^4 次 SumRange

复杂度目标：初始化 O(n)，每次查询 O(1)。
今日训练点：用一次预处理减少大量重复的区间求和。
建议限时：30 分钟。
提示触发点：独立思考 15 分钟后每次查询仍需要遍历 [left, right] 时，
可索取一级提示。

完成后请补充：
  - 初始化时间复杂度：
  - 单次查询时间复杂度：
  - 空间复杂度：
  - 如何统一处理 left == 0：
*/
type NumArray struct {
	// TODO: 在这里设计需要保存的状态。
	pres []int
	// 构造前N sum的数组
}

func Constructor(nums []int) NumArray {
	// TODO: 在这里完成初始化。
	pres := make([]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		if i == 0 {
			pres[i] = 0
		} else {
			pres[i] = pres[i-1] + nums[i-1]
		}
	}
	fmt.Println("pres:", pres)
	return NumArray{pres: pres}
}

func (n *NumArray) SumRange(left int, right int) int {
	// TODO: 在这里完成区间查询。
	sum := n.pres[right+1] - n.pres[left]
	return sum
}

func TestNumArray(t *testing.T) {
	nums := Constructor([]int{-2, 0, 3, -5, 2, -1})
	// -2, 0, 3, -5, 2, -1
	// 0, -2, -2, 1, -4, -2, -3
	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{name: "prefix", left: 0, right: 2, want: 1},
		{name: "middle-to-end", left: 2, right: 5, want: -1},
		{name: "whole-array", left: 0, right: 5, want: -3},
		{name: "single-element", left: 3, right: 3, want: -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nums.SumRange(tt.left, tt.right); got != tt.want {
				t.Fatalf("SumRange(%d, %d) = %d, want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}

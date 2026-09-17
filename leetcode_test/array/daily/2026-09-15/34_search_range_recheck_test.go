package daily20260915

import (
	"reflect"
	"testing"
)

/*
LeetCode 34. 在排序数组中查找元素的第一个和最后一个位置（中等）
链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

题目要求：
给定一个按非递减顺序排列的整数数组 nums 和一个目标值 target，返回 target 在数组中的
开始位置和结束位置；如果不存在，返回 [-1,-1]。

复杂度目标：时间 O(log n)，额外空间 O(1)。不能命中 target 后线性向两侧扫描。
复习要求：不打开旧实现，分别写出左边界与右边界的搜索过程。
建议限时：25 分钟。
提示触发点：独立思考 12 分钟后，命中 target 时仍然只能立即返回，可索取一级提示。

完成评价（2026-09-18）：
  - 测试结果：通过（7 个测试用例）
  - 搜索左边界时命中 target 后如何移动：保存候选下标，令 right = mid - 1
  - 搜索右边界时命中 target 后如何移动：保存候选下标，令 left = mid + 1
  - 候选答案的初始值：-1
  - 时间复杂度及原因：O(log n)，分别进行两次二分查找
  - 空间复杂度：O(1)
  - 核心模式与不变量：命中目标后不立即返回，继续向对应方向寻找更靠边的候选下标
  - 提示使用：无
  - 主要错误类型：无
  - 本次掌握度（0～3）：2
  - 下次复习日期：2026-09-21
  - 备注：独立写出左右边界二分并通过全部测试；耗时和复杂度口述未记录
*/
func searchRangeRecheck(nums []int, target int) []int {
	// 用二分法先找到最左边界
	// 在用二分法找到最右边界
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	return []int{left, right}
}

func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := -1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			res = mid
			right = mid - 1
			continue
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func searchRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := -1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			res = mid
			left = mid + 1
			continue
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func TestSearchRangeRecheck(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "basic", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, want: []int{3, 4}},
		{name: "missing", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, want: []int{-1, -1}},
		{name: "empty", nums: []int{}, target: 0, want: []int{-1, -1}},
		{name: "single-match", nums: []int{1}, target: 1, want: []int{0, 0}},
		{name: "all-equal", nums: []int{2, 2, 2, 2}, target: 2, want: []int{0, 3}},
		{name: "left-edge", nums: []int{1, 1, 2, 3}, target: 1, want: []int{0, 1}},
		{name: "right-edge", nums: []int{1, 2, 3, 3}, target: 3, want: []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRangeRecheck(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("searchRangeRecheck(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

package daily20260913

import (
	"reflect"
	"testing"
)

/*
LeetCode 34. 在排序数组中查找元素的第一个和最后一个位置（数组毕业复习）
链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

题目要求：
给定一个按非递减顺序排列的整数数组 nums 和目标值 target，返回 target 第一次和
最后一次出现的下标。如果 target 不存在，返回 [-1,-1]。

示例：

	nums = [5,7,7,8,8,10], target = 8，返回 [3,4]
	nums = [5,7,7,8,8,10], target = 6，返回 [-1,-1]

约束：

	0 <= len(nums) <= 10^5
	-10^9 <= nums[i], target <= 10^9
	nums 按非递减顺序排列

复杂度目标：时间 O(log n)，额外空间 O(1)。不能找到 target 后线性向两侧扫描。
复习要求：不打开旧实现，分别写出左边界与右边界的搜索过程。
建议限时：30 分钟。
提示触发点：独立思考 15 分钟后，命中 target 时仍然只能立即返回，可索取一级提示。

完成评价（2026-09-13）：
  - 测试结果：通过（8 个测试用例）
  - 时间复杂度：O(log n)，分别执行两次二分查找
  - 空间复杂度：O(1)
  - 核心模式与不变量：命中 target 后先保存候选下标；找左边界继续收缩右侧，找右边界继续收缩左侧
  - 提示使用：三级
  - 主要错误类型：建模、边界
  - 掌握度：1/3
  - 下次复习日期：2026-09-15
  - 备注：已经理解候选答案的作用并写出正确主干；需要继续巩固闭区间模板及命中后的搜索方向。耗时未记录
*/
func searchRangeFinalReview(nums []int, target int) []int {
	// 二分法找到左边界值、在找右边界值
	left := findLeft(nums, target)
	right := findRight(nums, target)
	return []int{left, right}
}

func findLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := -1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			res = mid // 很关键，这里是暂存结果
			right = mid - 1
		}
	}
	return res
}

func findRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := -1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			res = mid // 很关键，这里是暂存结果
			left = mid + 1
		}
	}
	return res
}

func TestSearchRangeFinalReview(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "basic", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, want: []int{3, 4}},
		{name: "not-found-middle", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, want: []int{-1, -1}},
		{name: "empty", nums: []int{}, target: 0, want: []int{-1, -1}},
		{name: "all-target", nums: []int{4, 4, 4, 4, 4}, target: 4, want: []int{0, 4}},
		{name: "left-edge", nums: []int{1, 1, 1, 3, 5}, target: 1, want: []int{0, 2}},
		{name: "right-edge", nums: []int{1, 3, 5, 7, 7}, target: 7, want: []int{3, 4}},
		{name: "single-found", nums: []int{9}, target: 9, want: []int{0, 0}},
		{name: "outside-range", nums: []int{2, 4, 6}, target: 8, want: []int{-1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRangeFinalReview(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("searchRangeFinalReview(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

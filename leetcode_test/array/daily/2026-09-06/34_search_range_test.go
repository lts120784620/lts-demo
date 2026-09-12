package daily20260906

import (
	"reflect"
	"testing"
)

/*
LeetCode 34. 在排序数组中查找元素的第一个和最后一个位置（中等）
链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

题目要求：
给定一个按非递减顺序排列的整数数组 nums 和目标值 target，返回 target 在数组中
第一次和最后一次出现的下标。如果 target 不存在，返回 [-1,-1]。

示例：

	nums = [5,7,7,8,8,10], target = 8，返回 [3,4]
	nums = [5,7,7,8,8,10], target = 6，返回 [-1,-1]
	nums = [], target = 0，返回 [-1,-1]

约束：

	0 <= len(nums) <= 10^5
	-10^9 <= nums[i] <= 10^9
	nums 按非递减顺序排列
	-10^9 <= target <= 10^9

复杂度要求：时间 O(log n)。找到一个 target 后线性向两侧扫描不满足要求。
今日训练点：让二分查找在命中目标后继续寻找左边界或右边界。
建议限时：30 分钟。
提示触发点：独立思考 15 分钟后，仍然只能找到任意一个 target 的位置时，
可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 左边界搜索的不变量：
  - 右边界搜索的不变量：
*/
func searchRange(nums []int, target int) []int {
	// 二分法找到
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	return []int{left, right}
}

func searchLeft(nums []int, target int) int {
	// 二分法找到
	left := 0
	right := len(nums) - 1
	res := -1
	// 找左侧
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			// 等于的情况正常就推出了
			// 这种还需要通过二分法找到最左边的target
			res = mid // 临时记一下
			right = mid - 1
			continue
		}
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func searchRight(nums []int, target int) int {
	// 二分法找到
	left := 0
	right := len(nums) - 1
	res := -1
	// 找右侧
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			// 等于的情况正常就推出了
			// 这种还需要通过二分法找到最右边的target
			res = mid // 临时记一下
			left = mid + 1
			continue
		}
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "basic", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, want: []int{3, 4}},
		{name: "not-found", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, want: []int{-1, -1}},
		{name: "empty", nums: []int{}, target: 0, want: []int{-1, -1}},
		{name: "all-target", nums: []int{2, 2, 2, 2}, target: 2, want: []int{0, 3}},
		{name: "left-edge", nums: []int{1, 1, 2, 3}, target: 1, want: []int{0, 1}},
		{name: "right-edge", nums: []int{1, 2, 3, 3}, target: 3, want: []int{2, 3}},
		{name: "single-found", nums: []int{1}, target: 1, want: []int{0, 0}},
		{name: "single-not-found", nums: []int{1}, target: 0, want: []int{-1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("searchRange(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

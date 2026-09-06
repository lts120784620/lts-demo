package daily20260902

import "testing"

/*
LeetCode 33. 搜索旋转排序数组（中等）
链接：https://leetcode.cn/problems/search-in-rotated-sorted-array/

题目要求：
整数数组 nums 原本按升序排列，且所有值互不相同。数组可能在未知下标 k 处
旋转，例如 [0,1,2,4,5,6,7] 可能变成 [4,5,6,7,0,1,2]。

给定旋转后的 nums 和整数 target：如果 target 存在，返回它当前的下标；否则
返回 -1。

示例：

	nums = [4,5,6,7,0,1,2], target = 0，返回 4
	nums = [4,5,6,7,0,1,2], target = 3，返回 -1
	nums = [1], target = 0，返回 -1

约束：

	1 <= len(nums) <= 5000
	-10^4 <= nums[i] <= 10^4
	nums 中的元素互不相同
	nums 保证由某个升序数组旋转得到
	-10^4 <= target <= 10^4

复杂度要求：时间 O(log n)。
今日训练点：数组整体不再有序时，识别二分中点两侧的局部有序区间。
建议限时：35 分钟。
提示触发点：独立思考 15 分钟后，仍无法判断中点左右哪一侧一定有序时，
可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 每轮如何识别有序的一侧：
  - 如何判断 target 是否位于该有序区间：
*/
func search(nums []int, target int) int {
	// TODO: 在这里实现。
	// 用二分法才能logn的时间复杂度
	has := -1
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			has = mid
			break
		}
		// 判断那一半有序，在有序的判断边界，才能确定下一巡环的left&right
		// if target < mid 且左边有序·
		if nums[left] <= nums[mid] {
			// 左边有序 && 中值处于有序边界内，则向左边2分
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
				continue
			}
		} else if nums[right] >= nums[mid] {
			// 右边有序 && 中值处于有序边界内，则向右边2分
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
				continue
			}
		}

		if nums[left] > nums[mid] {
			// 左边无序,也要移动过去试试
			right = mid - 1
		} else {
			// 右边无序,也要移动过去试试
			left = mid + 1
		}
	}
	return has
}

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "found-after-pivot", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{name: "not-found", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{name: "found-before-pivot", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 6, want: 2},
		{name: "single-not-found", nums: []int{1}, target: 0, want: -1},
		{name: "single-found", nums: []int{1}, target: 1, want: 0},
		{name: "two-elements", nums: []int{3, 1}, target: 1, want: 1},
		{name: "pivot-near-start", nums: []int{5, 1, 3}, target: 5, want: 0},
		{name: "test1", nums: []int{5, 1, 3}, target: 3, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Fatalf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

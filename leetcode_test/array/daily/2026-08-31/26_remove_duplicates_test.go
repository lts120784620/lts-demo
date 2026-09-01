package daily20260831

import (
	"lts-demo/leetcode_test/unit"
	"testing"
)

/*
LeetCode 26. 删除有序数组中的重复项（简单）
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

题目要求：
给定一个按非递减顺序排列的整数数组 nums，原地删除重复元素，使每个元素只
出现一次，并返回唯一元素的数量 k。

完成后必须满足：
 1. nums 的前 k 个元素依次包含所有唯一元素。
 2. 唯一元素之间的相对顺序保持不变。
 3. 下标 k 之后的内容不参与评测。
 4. 使用 O(1) 额外空间，不创建保存结果的新数组。

示例：

	输入：nums = [0,0,1,1,1,2,2,3,3,4]
	输出：k = 5，nums 的前五个元素为 [0,1,2,3,4]

约束：

	1 <= len(nums) <= 3 * 10^4
	-100 <= nums[i] <= 100
	nums 已按非递减顺序排列

今日训练点：识别“有序”条件带来的信息，并原地维护有效区间。
建议限时：20 分钟。
提示触发点：独立思考 10 分钟后仍需要额外数组或 map 时，可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 关键不变量：
  - 如果 nums 没有排序，当前方法为什么会失效：
*/
func removeDuplicates(nums []int) int {
	// TODO: 在这里实现。
	i := 0     // [:i]代表排序好的有效区间
	j := i + 1 // [j:]代表交换后的无效区间
	total := len(nums)
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			// 如果不相同 && j!=i+1 就和 i+1交换
			if j > i+1 {
				unit.ArraysSwap(nums, j, i+1)
				total--
			}
			i++
			j++
		}
	}
	return i + 1
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "basic", nums: []int{1, 1, 2}, want: []int{1, 2}},
		{name: "multiple-groups", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: []int{0, 1, 2, 3, 4}},
		{name: "single", nums: []int{1}, want: []int{1}},
		{name: "all-same", nums: []int{2, 2, 2, 2}, want: []int{2}},
		{name: "already-unique", nums: []int{-3, -1, 0, 2}, want: []int{-3, -1, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK := removeDuplicates(tt.nums)
			if gotK != len(tt.want) {
				t.Fatalf("removeDuplicates() returned k = %d, want %d", gotK, len(tt.want))
			}
			for i := range tt.want {
				if tt.nums[i] != tt.want[i] {
					t.Fatalf("nums[:k] = %v, want %v", tt.nums[:gotK], tt.want)
				}
			}
		})
	}
}

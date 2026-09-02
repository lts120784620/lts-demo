package daily20260901

import (
	"lts-demo/leetcode_test/unit"
	"sort"
	"testing"
)

/*
LeetCode 27. 移除元素（复习）
链接：https://leetcode.cn/problems/remove-element/

题目要求：
给定整数数组 nums 和整数 val，原地移除所有等于 val 的元素，返回数组中不等于
val 的元素数量 k。处理后 nums 的前 k 个位置必须包含所有保留元素；元素顺序
可以改变，下标 k 之后的内容不参与评测。

约束：

	0 <= len(nums) <= 100
	0 <= nums[i] <= 50
	0 <= val <= 100

今日训练方式：不打开旧代码，从记忆重写，并说清楚有效区间的含义。
建议限时：12 分钟。
提示触发点：8 分钟仍无法确定返回值代表下标还是数量时，可索取一级提示。

完成后请补充：
  - 时间复杂度：
  - 空间复杂度：
  - 关键不变量：
  - 这次是否能在无提示情况下完成：
*/
func removeElementReview(nums []int, val int) int {
	// TODO: 从记忆重新实现，不要打开旧代码。
	i := 0 // 第1个val的
	j := i // 第1个非val的
	for j < len(nums) {
		if nums[i] != val {
			i++
			j++
		} else if nums[j] == val {
			j++
		} else {
			unit.ArraysSwap(nums, i, j)
		}
	}
	return i
}

func TestRemoveElementReview(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want []int
	}{
		{name: "basic", nums: []int{3, 2, 2, 3}, val: 3, want: []int{2, 2}},
		{name: "mixed", nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, want: []int{0, 0, 1, 3, 4}},
		{name: "empty", nums: []int{}, val: 1, want: []int{}},
		{name: "remove-all", nums: []int{1, 1}, val: 1, want: []int{}},
		{name: "remove-none", nums: []int{1, 2, 3}, val: 4, want: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK := removeElementReview(tt.nums, tt.val)
			if gotK != len(tt.want) {
				t.Fatalf("removeElementReview() returned k = %d, want %d", gotK, len(tt.want))
			}
			got := append([]int(nil), tt.nums[:gotK]...)
			sort.Ints(got)
			sort.Ints(tt.want)
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Fatalf("nums[:k] = %v, want elements %v", got, tt.want)
				}
			}
		})
	}
}

package daily20260827

import (
	"fmt"
	"sort"
	"testing"
)

/*
LeetCode 27. 移除元素（简单）
链接：https://leetcode.cn/problems/remove-element/

题目要求：
给定一个整数数组 nums 和一个整数 val，原地移除 nums 中所有等于 val 的元素，
并返回数组中不等于 val 的元素数量 k。

完成后必须满足：
 1. nums 的前 k 个位置包含所有不等于 val 的元素。
 2. 前 k 个元素的顺序可以改变。
 3. 下标 k 之后的内容不参与评测。
 4. 不要创建另一个数组来保存最终结果。

示例：

	输入：nums = [3,2,2,3], val = 3
	输出：k = 2，nums 的前两个元素为 [2,2]

约束：

	0 <= len(nums) <= 100
	0 <= nums[i] <= 50
	0 <= val <= 100

今日训练点：数组原地修改。
建议限时：15 分钟。
提示触发点：独立思考 8 分钟仍无法描述“哪些元素需要保留”时，可索取一级提示。

完成后请在这里补充：
  - 时间复杂度：
  - 空间复杂度：
  - 关键不变量：
*/
func removeElement(nums []int, val int) int {
	j := len(nums) - 1
	i := 0
	for i <= j {
		if nums[i] != val {
			// [:i] 是前k个位置
			i++
		} else if nums[j] == val {
			// [j:] 是len-k个
			j--
		} else {
			// 交换
			t := nums[i]
			nums[i] = nums[j]
			nums[j] = t
			i++
			j--
		}

	}
	fmt.Println(nums)
	return i
}

func TestRemoveElement(t *testing.T) {
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
			gotK := removeElement(tt.nums, tt.val)
			if gotK != len(tt.want) {
				t.Fatalf("removeElement() returned k = %d, want %d", gotK, len(tt.want))
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

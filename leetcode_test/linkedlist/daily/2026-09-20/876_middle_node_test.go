package daily20260920

import (
	"reflect"
	"testing"
)

/*
LeetCode 876. 链表的中间结点（简单）
链接：https://leetcode.cn/problems/middle-of-the-linked-list/

题目要求：
给定单链表头结点 head，返回链表的中间结点；如果有两个中间结点，返回第二个。

复杂度目标：时间 O(n)，额外空间 O(1)。
今日训练点：快慢指针的速度关系，以及奇偶长度链表的停止条件。
建议限时：15 分钟。
提示触发点：独立思考 10 分钟后，仍然需要先遍历一次统计长度，或无法判断偶数长度时
慢指针应停在哪里，可索取一级提示。

完成评价（2026-09-20）：
  - 测试结果：通过（5 个测试用例）
  - 快慢指针每轮分别移动几步：p1 移动 1 步，p2 移动 2 步
  - 循环继续的条件：p2 != nil && p2.Next != nil
  - 为什么偶数长度会返回第二个中间结点：快指针走到 nil 时，慢指针恰好前进链表长度的一半
  - 时间复杂度及原因：O(n)，慢指针最多遍历一半链表，快指针最多遍历整条链表
  - 空间复杂度：O(1)
  - 核心模式与不变量：快指针走过的距离始终是慢指针的两倍
  - 提示使用：三级
  - 主要错误类型：边界
  - 本次掌握度（0～3）：1
  - 下次复习日期：2026-09-22
  - 备注：初版仅判断快指针非空，奇数长度会使慢指针多走一步；明确停止条件后通过全部测试，耗时未记录
*/
func middleNodeDaily(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next // 慢指针一次移动1个
		p2 = p2.Next // 快指针一次移动2个
		if p2 != nil {
			p2 = p2.Next
		}
	}
	return p1
}

func TestMiddleNodeDaily(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{name: "odd", values: []int{1, 2, 3, 4, 5}, want: []int{3, 4, 5}},
		{name: "even", values: []int{1, 2, 3, 4, 5, 6}, want: []int{4, 5, 6}},
		{name: "single", values: []int{1}, want: []int{1}},
		{name: "two-nodes", values: []int{1, 2}, want: []int{2}},
		{name: "empty", values: []int{}, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.values)
			got := listToSlice(middleNodeDaily(head))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("middleNodeDaily(%v) = %v, want suffix %v", tt.values, got, tt.want)
			}
		})
	}
}

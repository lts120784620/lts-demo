package daily20260919

import (
	"reflect"
	"testing"
)

/*
LeetCode 83. 删除排序链表中的重复元素（简单）
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

题目要求：
给定一个升序链表，删除所有重复元素，使每个元素只出现一次，并返回链表头结点。

复杂度目标：时间 O(n)，额外空间 O(1)。
今日训练点：利用链表有序性比较相邻结点；删除后是否移动当前指针。
建议限时：20 分钟。
提示触发点：独立思考 12 分钟后，遇到三个连续相同值时仍会漏删，可索取一级提示。

完成评价（2026-09-20）：
  - 测试结果：通过（5 个测试用例）
  - 为什么只需比较相邻结点：链表已有序，相同值必然连续出现
  - 删除重复结点后是否移动当前指针：保留结点 pre 不移动，cur 更新为 pre.Next
  - 保留当前结点时如何移动：pre、cur 各向后移动一个结点
  - 时间复杂度及原因：O(n)，每个结点最多访问一次
  - 空间复杂度：O(1)
  - 核心模式与不变量：pre 是最后一个保留结点，cur 是当前待判断结点
  - 提示使用：无（过程未记录）
  - 主要错误类型：无
  - 本次掌握度（0～3）：2
  - 下次复习日期：2026-09-23
  - 备注：连续重复值处理正确并通过全部测试；耗时与独立完成情况未记录
*/
func deleteDuplicatesDaily(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := head
	cur := head.Next
	for cur != nil {
		if pre.Val == cur.Val {
			// 删除cur
			pre.Next = cur.Next
			cur = pre.Next
			continue
		}
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}

func TestDeleteDuplicatesDaily(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{name: "basic", values: []int{1, 1, 2}, want: []int{1, 2}},
		{name: "multiple-groups", values: []int{1, 1, 2, 3, 3}, want: []int{1, 2, 3}},
		{name: "empty", values: []int{}, want: []int{}},
		{name: "single", values: []int{1}, want: []int{1}},
		{name: "all-equal", values: []int{2, 2, 2, 2}, want: []int{2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.values)
			got := listToSlice(deleteDuplicatesDaily(head))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("deleteDuplicatesDaily(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}

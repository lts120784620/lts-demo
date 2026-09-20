package daily20260919

import (
	"reflect"
	"testing"
)

/*
LeetCode 21. 合并两个有序链表（简单）
链接：https://leetcode.cn/problems/merge-two-sorted-lists/

题目要求：
将两个升序链表合并为一个新的升序链表，并返回合并后链表的头结点。

复杂度目标：时间 O(m+n)，额外空间 O(1)。
今日训练点：虚拟头结点、双指针与结果链表尾指针。
建议限时：20 分钟。
提示触发点：独立思考 12 分钟后，仍无法确定每次应移动哪个输入指针，或忘记处理
其中一条链表的剩余部分时，可索取一级提示。

完成评价（2026-09-20）：
  - 测试结果：通过（5 个测试用例）
  - 两个输入指针分别表示什么：cur1、cur2 分别指向两条链表当前待比较的结点
  - 结果链表尾指针表示什么：res 始终指向已合并链表的最后一个结点
  - 主循环结束后如何处理剩余结点：将尚未遍历完的链表直接接到 res.Next
  - 时间复杂度及原因：O(m+n)，两条链表的每个结点最多访问一次
  - 空间复杂度：O(1)
  - 核心模式与不变量：res 之前是已按序合并的部分，cur1、cur2 指向各自未处理部分
  - 提示使用：无（过程未记录）
  - 主要错误类型：无
  - 本次掌握度（0～3）：2
  - 下次复习日期：2026-09-23
  - 备注：实现主干正确并通过全部测试；耗时与独立完成情况未记录
*/
func mergeTwoListsDaily(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2
	dummy := &ListNode{}
	res := dummy
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			res.Next = cur1
			res = res.Next
			cur1 = cur1.Next
		} else {
			res.Next = cur2
			res = res.Next
			cur2 = cur2.Next
		}
	}
	if cur1 != nil {
		res.Next = cur1
	} else {
		res.Next = cur2
	}
	return dummy.Next
}

func TestMergeTwoListsDaily(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{name: "basic", list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, want: []int{1, 1, 2, 3, 4, 4}},
		{name: "both-empty", list1: []int{}, list2: []int{}, want: []int{}},
		{name: "first-empty", list1: []int{}, list2: []int{0}, want: []int{0}},
		{name: "second-empty", list1: []int{1, 2}, list2: []int{}, want: []int{1, 2}},
		{name: "negative-values", list1: []int{-3, -1, 2}, list2: []int{-2, 0, 3}, want: []int{-3, -2, -1, 0, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := listFromSlice(tt.list1)
			list2 := listFromSlice(tt.list2)
			got := listToSlice(mergeTwoListsDaily(list1, list2))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("mergeTwoListsDaily(%v, %v) = %v, want %v", tt.list1, tt.list2, got, tt.want)
			}
		})
	}
}

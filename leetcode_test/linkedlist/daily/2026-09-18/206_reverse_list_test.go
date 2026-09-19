package daily20260918

import (
	"reflect"
	"testing"
)

/*
LeetCode 206. 反转链表（简单）
链接：https://leetcode.cn/problems/reverse-linked-list/

题目要求：
给定单链表头结点 head，反转链表并返回新的头结点。

复杂度目标：时间 O(n)，额外空间 O(1)。
今日训练点：指针重连；修改 Next 前保存后继结点。
建议限时：25 分钟。
提示触发点：独立思考 12 分钟后，修改 cur.Next 后丢失剩余链表，或无法说明三个指针
分别代表哪一部分时，可索取一级提示。

完成评价（2026-09-19）：
  - 测试结果：通过（5 个测试用例）
  - 已反转部分由哪个指针表示：pre，它始终指向已反转部分的头结点
  - 当前待处理结点由哪个指针表示：cur
  - 为什么修改 Next 前要保存后继结点：避免覆盖唯一通往未处理部分的指针
  - 时间复杂度及原因：O(n)，每个结点只处理一次
  - 空间复杂度：O(1)
  - 核心模式与不变量：pre 指向已反转部分，cur 指向未处理部分，二者覆盖原链表全部结点
  - 提示使用：三级
  - 主要错误类型：建模、实现
  - 本次掌握度（0～3）：1
  - 下次复习日期：2026-09-21
  - 备注：初版重连方向会形成自环并跳过 cur；在关键指针轨迹讲解后完成，耗时未记录
*/
func reverseListDaily(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		// 反转cur
		tmp := cur.Next
		cur.Next = pre
		// 向后移动
		pre = cur
		cur = tmp
	}
	return pre
}

func TestReverseListDaily(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{name: "basic", values: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "empty", values: []int{}, want: []int{}},
		{name: "single", values: []int{1}, want: []int{1}},
		{name: "two-nodes", values: []int{1, 2}, want: []int{2, 1}},
		{name: "duplicates", values: []int{1, 1, 2}, want: []int{2, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.values)
			got := listToSlice(reverseListDaily(head))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("reverseListDaily(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}

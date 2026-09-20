package daily20260920

import (
	"reflect"
	"testing"
)

/*
LeetCode 19. 删除链表的倒数第 N 个结点（中等）
链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

题目要求：
给定单链表头结点 head 和整数 n，删除链表的倒数第 n 个结点，并返回头结点。

复杂度目标：只遍历链表一次，时间 O(n)，额外空间 O(1)。
今日训练点：虚拟头结点、固定间距双指针，以及删除头结点的统一处理。
建议限时：25 分钟。
提示触发点：独立思考 12 分钟后，仍然只能先统计长度再删除，或在删除头结点时需要
额外分支，可索取一级提示。

完成评价（2026-09-20）：
  - 测试结果：通过（5 个测试用例）
  - 为什么使用虚拟头结点：统一删除头结点和普通结点的处理
  - 快慢指针之间需要保持怎样的间距：p2 始终领先 p1 n-1 个结点
  - 同步移动在什么条件下停止：p2 到达链表最后一个结点时执行删除，随后遍历结束
  - 停止时慢指针位于待删除结点的什么位置：p1 指向待删除结点，p1Pre 指向其前驱
  - 时间复杂度及原因：O(n)，快慢指针都只向后移动且不回退
  - 空间复杂度：O(1)
  - 核心模式与不变量：保持 p2 与 p1 的固定间距，用 p1Pre 完成结点重连
  - 提示使用：无
  - 主要错误类型：无
  - 本次掌握度（0～3）：2
  - 下次复习日期：2026-09-23
  - 备注：独立写出固定间距双指针并通过全部测试；耗时与复杂度口述未记录
*/
func removeNthFromEndDaily(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1Pre := dummy
	p1 := head
	p2 := head
	// 先移动p2
	for i := 1; i < n; i++ {
		p2 = p2.Next
	}
	// 开始遍历直到 p2.next = nil
	for p1 != nil && p2 != nil {
		if p2.Next == nil {
			// 删除p1
			p1Pre.Next = p1.Next
		}
		p1Pre = p1
		p1 = p1.Next
		p2 = p2.Next
	}
	return dummy.Next
}

func TestRemoveNthFromEndDaily(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		n      int
		want   []int
	}{
		{name: "middle", values: []int{1, 2, 3, 4, 5}, n: 2, want: []int{1, 2, 3, 5}},
		{name: "remove-head", values: []int{1, 2, 3}, n: 3, want: []int{2, 3}},
		{name: "remove-tail", values: []int{1, 2, 3}, n: 1, want: []int{1, 2}},
		{name: "single", values: []int{1}, n: 1, want: []int{}},
		{name: "two-remove-head", values: []int{1, 2}, n: 2, want: []int{2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.values)
			got := listToSlice(removeNthFromEndDaily(head, tt.n))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("removeNthFromEndDaily(%v, %d) = %v, want %v", tt.values, tt.n, got, tt.want)
			}
		})
	}
}

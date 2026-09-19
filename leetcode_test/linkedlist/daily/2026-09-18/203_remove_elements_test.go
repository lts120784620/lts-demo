package daily20260918

import (
	"reflect"
	"testing"
)

/*
LeetCode 203. 移除链表元素（简单）
链接：https://leetcode.cn/problems/remove-linked-list-elements/

题目要求：
给定链表头结点 head 和整数 val，删除链表中所有满足 Node.Val == val 的结点，
并返回新的头结点。

复杂度目标：时间 O(n)，额外空间 O(1)。
今日训练点：虚拟头结点；让删除头结点和删除中间结点使用同一套逻辑。
建议限时：20 分钟。
提示触发点：独立思考 12 分钟后，仍然需要为连续删除多个头结点编写单独循环时，
可索取一级提示。

完成评价（2026-09-19）：
  - 测试结果：通过（5 个测试用例）
  - 为什么需要虚拟头结点：统一删除原头结点和普通结点的逻辑
  - 检查的是当前结点还是当前结点的 Next：当前结点 cur
  - 删除结点后指针是否立即前进：cur 前进，pre 留在最后一个保留结点
  - 时间复杂度及原因：O(n)，每个结点只访问一次
  - 空间复杂度：O(1)
  - 核心模式与不变量：pre 指向最后一个保留结点，cur 指向当前待判断结点
  - 提示使用：无（完成后复盘）
  - 主要错误类型：无
  - 本次掌握度（0～3）：2
  - 下次复习日期：2026-09-22
  - 备注：独立写出正确主干并通过全部测试；耗时未记录
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := dummy.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		val    int
		want   []int
	}{
		{name: "basic", values: []int{1, 2, 6, 3, 4, 5, 6}, val: 6, want: []int{1, 2, 3, 4, 5}},
		{name: "empty", values: []int{}, val: 1, want: []int{}},
		{name: "remove-all", values: []int{7, 7, 7}, val: 7, want: []int{}},
		{name: "remove-head", values: []int{1, 2, 3}, val: 1, want: []int{2, 3}},
		{name: "consecutive-middle", values: []int{1, 2, 2, 1}, val: 2, want: []int{1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.values)
			got := listToSlice(removeElements(head, tt.val))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("removeElements(%v, %d) = %v, want %v", tt.values, tt.val, got, tt.want)
			}
		})
	}
}

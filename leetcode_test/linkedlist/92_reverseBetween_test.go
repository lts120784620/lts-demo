package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.92. 反转链表 II
描述：
	给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
思路：
    1、递归法：保存right节点，返回left节点倒置
时间：
    1、2021/10/10
*/

func TestReverseBetween(t *testing.T) {
	l := NewForArrays([]interface{}{1})
	fmt.Println(reverseBetween(l, 1, 4).String())
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p := head
	for i := 0; i < left-1; i++ {
		p = p.Next
		dummy = dummy.Next
	}
	last, _ := reverse(p, right-left+1, nil)
	// 原头指向倒置后头
	dummy.Next = last

	if left <= 1 {
		return last
	}
	return head
}

func reverse(head *ListNode, n int, final *ListNode) (*ListNode, *ListNode) {
	// 右边终止判断,倒置n-1个
	if head.Next == nil || n == 1 {
		// 细节二：保存最右的后置节点
		return head, head.Next
	}
	// 向右偏移
	n--
	last, final := reverse(head.Next, n, final)
	head.Next.Next = head
	// 细节一：指向最右的后置节点
	head.Next = final
	return last, final
}

package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.19 删除链表的倒数第 N 个结点
描述：
	给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
	进阶：你能尝试使用一趟扫描实现吗？
思路：
    1、经典双指针法：p1、p2，p2先走n后p1出发，当p2走到结尾，屁p1位置即为n-k
时间：
    1、2021/10/9
*/

func TestRemoveNthFromEnd(t *testing.T) {
	l := NewForArrays([]interface{}{1})
	fmt.Println(removeNthFromEnd(l, 1).String())
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 细节一：哑节点避免特殊判断
	dummy := &ListNode{Next: head}
	p1, p2 := head, dummy
	// 先走n-1步，因为链表删除要找到要删除的前一个节点
	for i := 0; i < n && p1 != nil; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	// 删除p2
	if p2.Next != nil {
		p2.Next = p2.Next.Next
	}

	return dummy.Next
}

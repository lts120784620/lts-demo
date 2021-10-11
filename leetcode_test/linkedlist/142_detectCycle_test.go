package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.142. 环形链表 II
描述：
	给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
	说明：不允许修改给定的链表。
	进阶：
		你是否可以使用 O(1) 空间解决此题？
思路：
    1、双指针法：先找到相交位置，在把一个节点放到head，两个指针同时走，在相遇时的位置到head的距离。
时间：
    1、2021/10/9
*/

func TestDetectCycle(t *testing.T) {
	//n4 := &ListNode{Val: -4}
	//n3 := &ListNode{Val: 0, Next: n4}
	//n2 := &ListNode{Val: 2, Next: n3}
	//n4.Next = n2 // 环
	//head := &ListNode{Val: 3, Next: n2}

	n2 := &ListNode{Val: 2}
	head := &ListNode{Val: 1, Next: n2}
	n2.Next = head

	fmt.Println(detectCycle(head).Val)
}

func detectCycle(head *ListNode) *ListNode {
	// p2快，p1慢
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		// 走一步
		p1 = p1.Next
		// 走两步
		p2 = p2.Next.Next
		if p1 == p2 {
			// p2清零，从头开始
			p2 = head
			break
		}
	}

	if p2 == nil || p2.Next == nil {
		return nil
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

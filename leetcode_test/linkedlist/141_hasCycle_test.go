package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.141. 环形链表
描述：
	给定一个链表，判断链表中是否有环。
	如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
	如果链表中存在环，则返回 true 。 否则，返回 false 。
思路：
    1、双指针法：p1比p2速度快一倍，p1一次走一步，p2一次走两步。当p1追上p2即有环存在
时间：
    1、2021/10/9
*/

func TestHasCycle(t *testing.T) {
	n4 := &ListNode{Val: -4}
	n3 := &ListNode{Val: 0, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n4.Next = n2 // 环
	n1 := &ListNode{Val: 3, Next: n2}

	fmt.Println(hasCycle(n1))
}

func hasCycle(head *ListNode) bool {
	// p2快，p1慢
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		// 走一步
		p1 = p1.Next
		// 走两步
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}

	return false
}

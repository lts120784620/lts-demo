package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.234. 回文链表
描述：
	给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false
思路：
    1、先用双指针法找到中间节点，用递归方式遍历后一半与前一半比较
时间：
    1、2021/10/10
*/

func TestIsPalindrome(t *testing.T) {
	l := NewForArrays([]interface{}{1, 3, 1})
	fmt.Println(isPalindrome(l))
}

// a-b-b-a
func isPalindrome(head *ListNode) bool {
	// p1：慢指针，一次走一步
	// p2：快指针，没病走两步
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	// p1位置为中点,从p1开始倒置,p2指向开头
	p2 = head
	last := postOrder(p1)
	for last != nil && p2 != nil {
		if last.Val != p2.Val {
			return false
		}
		last = last.Next
		p2 = p2.Next
	}

	return true
}

func postOrder(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := postOrder(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

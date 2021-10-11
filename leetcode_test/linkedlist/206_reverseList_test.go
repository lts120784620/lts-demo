package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.206. 反转链表
描述：
	给你单链表的头节点 head ，请你反转链表，并返回反转后的链表
思路：
    1、迭代法：定义pre、cur双指针，向后移动即可，注意最后head的后指针
	2、递归法:
时间：
    1、2021/10/9
*/

func TestReverseList(t *testing.T) {
	l := NewForArrays([]interface{}{1, 2, 3, 4, 5})
	fmt.Println(reverseList(l).String())
	l = NewForArrays([]interface{}{1, 2, 3, 4, 5})
	fmt.Println(reverseList1(l).String())
}

// 迭代法：重点是dummy节点，临时节点保存后面指针
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	// 细节一：最后要把原头节点指向nil，否则会出现dummy和head循环首尾相接的问题
	head.Next = nil
	return pre
}

// 递归法：
func reverseList1(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	// 这种情况可以递归，但是拿不到最后一个节点
	last := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

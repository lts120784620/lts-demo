package linkedlist

import "testing"

/**
No.206 链表翻转
思路：
	1、非递归
	两个指针，一个代表当前指针cur，一个前指针prev，从head开始用cur循环遍历链表，
用临时指针保存当前后指针t=cur.next，目的是防止后指针前置后丢失，再将cur.next指向前指针prev,
此时prev向后移赋给cur，cur用临时指针t向后移动。待整个链表移动结束后，prev作为头指针退出
	2、递归
	递归的本质是栈结构，把函数想想成栈，顺序入栈，则尾节点优先出栈，如果当前出栈的节点为nil,则作为新链表的头结点，
否则将当前节点的next的next指针指向当前节点
*/

func init() {
}

func TestReverseList1(t *testing.T) {
	if l.Head == nil || l.Head.Next == nil{
		return
	}
	cur := l.Head
	var prev *ListNode
	for cur != nil{
		// 三元赋值，其实就是用了一个中间变量，与下面的写法一样
		cur.Next,prev,cur = prev,cur,cur.Next
	}
	l.Head = prev

	l.Travelse()
}

func TestReverseList2(t *testing.T) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}
	cur := l.Head
	var prev *ListNode
	for cur != nil {
		t := cur.Next
		cur.Next = prev
		prev = cur
		cur = t
	}
	l.Head = prev

	l.Travelse()
}

// 翻转链表，递归
func TestReverseListByRecursive(t *testing.T) {
	l := NewInitLinkedList(5)
	l.Head = recursion(l.Head)
	l.Travelse()
}

func recursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := recursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return n
}
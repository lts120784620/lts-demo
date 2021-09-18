package linkedlist

import "fmt"

/**
合并两个有序链表
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1.Next.Next
	p2 := l2.Next.Next
	dummyHead := &ListNode{0, p1}
	for p2 != nil {
		if p1.Next != nil && p1.Val <= p2.Val {
			p1 = p1.Next
		} else {
			fmt.Println("before:", p1.Val, p1.Next.Val, p2.Val, p2.Next.Val)
			p1.Next, p2.Next, p1, p2 = p2, p1.Next, p2.Next, p2
			fmt.Println("after:", p1.Val, p1.Next.Val, p2.Val, p2.Next.Val)
		}
	}
	return dummyHead.Next
}

/**
判断链表是否为回文链表
*/
func IsPalindrome(l *ListNode) bool {
	head := l.Next.Next
	if head == nil || head.Next == nil {
		return true
	}
	// 取中点
	p1 := head
	p2 := head.Next
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	//p1中点
	// 翻转后一半链表
	//fmt.Println(p1.Val,p2.Val)
	var pre *ListNode
	cur := p1.Next
	var next *ListNode
	//fmt.Println(pre.Val,cur.Val,next.Val)
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	p1.Next = pre
	// 验证相等
	h := head
	q := p1.Next
	for h != nil && q != nil && h.Val == q.Val {
		h = h.Next
		q = q.Next
	}
	return q == nil
}

package linkedList

import "fmt"

/**
合并两个有序链表
 */
func MergeTwoLists(l1 *MyLinkedList, l2 *MyLinkedList) *Node{
	p1 := l1.Head.Next
	p2 := l2.Head.Next
	dummyHead := &Node{"",p1}
	for p2 != nil{
		if p1.Next != nil && p1.Data <= p2.Data{
			p1 = p1.Next
		}else{
			fmt.Println("before:",p1.Data,p1.Next.Data,p2.Data,p2.Next.Data)
			p1.Next,p2.Next,p1,p2 = p2,p1.Next,p2.Next,p2
			fmt.Println("after:",p1.Data,p1.Next.Data,p2.Data,p2.Next.Data)
		}
	}
	return dummyHead.Next
}

/**
判断链表是否为回文链表
 */
func IsPalindrome(l *MyLinkedList) bool {
	head := l.Head.Next
	if head == nil || head.Next == nil{
		return true
	}
	// 取中点
	p1 := head
	p2 := head.Next
	for p2 != nil && p2.Next != nil{
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	//p1中点
	// 翻转后一半链表
	//fmt.Println(p1.Data,p2.Data)
	var pre *Node
	cur := p1.Next
	var next *Node
	//fmt.Println(pre.Data,cur.Data,next.Data)
	for cur != nil{
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	p1.Next = pre
	// 验证相等
	h := head
	q := p1.Next
	for h!= nil && q != nil && h.Data == q.Data{
		h = h.Next
		q = q.Next
	}
	return q == nil
}

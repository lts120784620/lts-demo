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
	p1 := l.Head
	p2 := l.Head
	for p1 != nil && p2.Next != nil{
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	//p1中点
	t := p1
	pre := p1
	cur := p1
	for cur != nil{

	}
}

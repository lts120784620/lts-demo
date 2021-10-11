package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.160. 相交链表
描述：
	给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
思路：
    1、双指针法：重点是想办法可以让两个指针可以遍历到共同的节点，逻辑上将AB两个链表链接起来，判断p1p2相等或都为空时，即是相交
时间：
    1、2021/10/9
*/

func TestGetIntersectionNode(t *testing.T) {
	s := NewForArrays([]interface{}{8, 4, 5})
	a := NewForArrays([]interface{}{4, 1})
	a.Add(s)
	b := NewForArrays([]interface{}{5, 0, 1})
	b.Add(s)

	fmt.Println(getIntersectionNode(a, b).Val)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	// 每次比较两个指针，若到链表结尾则指向另一个链表的头结点，这样一次结束会抹除长度差，第二次则会在相同位置开始遍历
	for p1 != p2 {
		p1 = p1.Next
		if p1 == nil {
			p1 = headB
		}
		p2 = p2.Next
		if p2 == nil {
			p2 = headA
		}
	}
	return p1
}

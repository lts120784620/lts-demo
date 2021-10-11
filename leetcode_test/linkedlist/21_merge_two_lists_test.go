package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.21 合并两个有序链表
描述：
	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
思路：
	双指针法
    1、简单的方法，两个指针p1,p1分别指向l1,l2，比较大小，大的指向新链表，向后移动
	2、递归解法，终止条件为l1,l2为nil
时间：
	1、2021-09-12
*/

func TestMergeTwoLists(t *testing.T) {
	l1 := NewForArrays([]interface{}{1, 2, 4})
	l2 := NewForArrays([]interface{}{1, 3, 4})
	fmt.Println(mergeTwoLists(l1, l2))
	fmt.Println(mergeTwoLists2(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p0, p1, p2 := res, l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p0.Next = p1
			p1 = p1.Next
		} else {
			p0.Next = p2
			p2 = p2.Next
		}
		p0 = p0.Next
	}
	if p1 != nil {
		p0.Next = p1
	}

	if p2 != nil {
		p0.Next = p2
	}

	return res.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val >= l2.Val {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	}
}

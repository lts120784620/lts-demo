package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.25. K 个一组翻转链表
描述：
	给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
	k是一个正整数，它的值小于或等于链表的长度。
	如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
	进阶：
		你可以设计一个只使用常数额外空间的算法来解决此问题吗？
		你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
思路：
    1、
时间：
    1、2021/10/10
*/
func TestReverseKGroup(t *testing.T) {
	l := NewForArrays([]interface{}{1, 2, 3, 4, 5, 6})
	fmt.Println(reverseKGroup(l, 2).String())
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	for i := 0; i < k; i++ {
		// base case :不足k个节点退出
		if p == nil {
			return head
		}
		p = p.Next
	}
	lasthead := reverseKGroup(p, k)
	// 1-2- 3-4- 5-6
	// 6-5
	fmt.Println(head.String())
	last, _ := reverse(head, k, nil)
	head.Next = lasthead
	fmt.Println(head.String())
	//fmt.Println(last.String())
	//fmt.Println(head.String())
	return last
}

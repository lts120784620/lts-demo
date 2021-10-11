package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.24. 两两交换链表中的节点
描述：
	给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
	你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
思路：
    1、实现翻转链表的递归写法->实现链表前N个节点翻转->实现N个一组的翻转的递归写法
时间：
    1、2021/10/11
*/

func TestSwapPairs(t *testing.T) {
	l := NewForArrays([]interface{}{1, 2, 3, 4, 5})
	fmt.Println(swapPairs(l).String())
}

// 递归翻转前两个节点
func swapPairs(head *ListNode) *ListNode {
	p := head
	for i := 0; i < 2; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	// 1，2，3，4，5
	lastHead := swapPairs(p)
	fmt.Println(head)
	// head ->3-4-5
	// 开始倒置
	last, _ := reverseN(head, 1)
	// last ->4-3-5
	//fmt.Println(head)
	//fmt.Println(last)
	//fmt.Println(">>>>>>")
	head.Next = lastHead
	return last
}

// 翻转前N-1个节点
func reverseN(head *ListNode, n int) (*ListNode, *ListNode) {
	if head.Next == nil || n <= 0 {
		// 细节
		return head, head.Next
	}
	n--
	last, final := reverseN(head.Next, n)
	head.Next.Next = head
	// 细
	head.Next = final
	return last, final
}

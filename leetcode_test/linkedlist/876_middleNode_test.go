package linkedlist

import (
	"fmt"
	"testing"
)

/**
No.876 链表的中间结点
描述：
	给定一个头结点为 head 的非空单链表，返回链表的中间结点。
	如果有两个中间结点，则返回第二个中间结点。
思路：
    1、双指针法：p1比p2速度快一倍，p1一次走一步，p2一次走两步。当p2结束，p1就是中点位置
时间：
    1、2021/10/9
*/

func TestMiddleNode(t *testing.T) {
	l := NewForArrays([]interface{}{1, 2, 3, 4, 5})
	fmt.Println(middleNode(l).Val)
}

func middleNode(head *ListNode) *ListNode {
	// p2快，p1慢
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		// 走一步
		p1 = p1.Next
		// 走两步
		p2 = p2.Next.Next
	}

	return p1
}

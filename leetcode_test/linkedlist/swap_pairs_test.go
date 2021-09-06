package linkedlist

import "testing"

/**
No.24 两两交换链表中的节点
思路：
	使用三个指针遍历，prev作为移动遍历指针，a=prev.next作为左边交换的节点，b=a.next作为右边交换的节点a,b交换的顺序为，从前到后防止丢失点
(1)先用prev.next=b挂住b节点，完成prev指向b；
(2)用a后继指向c，完成a的交换；
(3)用b后继指向a，完成b的交换；
(4)prev移动到a的位置，a向后移动，b向后移动
整体做法不难，关键是要能流畅的写出代码来
*/

func TestSwapPairs(t *testing.T) {
	prev := l.Head
	for prev != nil && prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := a.Next
		// 开始交换
		prev.Next = b
		a.Next = b.Next
		b.Next = a
		// 交换结束，向后移动
		prev = a
	}
	l.Travelse()
}

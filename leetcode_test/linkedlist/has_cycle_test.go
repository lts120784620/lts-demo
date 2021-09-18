package linkedlist

import (
	"testing"
)

/**
No.141 判断链表是否有环
思路：
	1、硬遍历：无限循环判断是否存在nil节点，存在则无环，反之有环，可以设置最大超时时间，不推荐
	2、判重法：利用set/map等在遍历过程中，判断是否存在重复元素，存在则有环，反之无环，空间复杂度o(n)
	3、快慢指针法：使用两个指针，fast指针每次走一步，slow指针每次走两步，重合及有环，反之无环。因为跑道是环形的，无限跑快的一定会追上慢的
*/

func init() {
	// 创造一个有环链表
	//tail := l.Get(l.Length())
	//tail.Next = l.Get(l.Length() - l.Length()/2)
}

// 判重法
func TestHasCycle1(t *testing.T) {
	//m := map[string]struct{}{}
	//cur := l.Head
	//for cur != nil {
	//	if _, ok := m[cur.Data]; ok {
	//		println("有环")
	//		return
	//	}
	//	m[cur.Data] = struct{}{}
	//	cur = cur.Next
	//}
	println("无环")
}

// 快慢指针法
func TestHasCycle2(t *testing.T) {
	//slow, fast := l.Head, l.Head
	//for slow != nil && fast != nil && fast.Next != nil {
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//	if slow == fast {
	//		println("有环")
	//		return
	//	}
	//}
	println("无环")
}

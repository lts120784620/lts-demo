package linkedlist

/**
 删除倒数第k个节点
 */
func GetNthFromEnd(l *MyLinkedList, k int) *ListNode {
	first := l.Head
	second := l.Head
	i := 0
	// 第二个指针比第一个快k个节点
	for i< k && second != nil{
		i++
		second = second.Next
	}
	if second == nil{
		return l.Head.Next
	}
	// 找到倒数第k-1个
	for second.Next != nil {
		first = first.Next
		second = second.Next
	}
	// 删除
	first.Next = first.Next.Next
	return l.Head
}

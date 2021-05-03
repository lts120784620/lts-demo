package linkedlist

/**
 求链表的中间节点
 */
func MiddleNode(l *MyLinkedList) *ListNode {
	slow := l.Head
	fast := l.Head
	if l.Head.Next == nil {
		return l.Head
	}
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

/**
判断链表是否有环
 */
func HasCycle(l *MyLinkedList) bool {
	slow := l.Head
	fast := l.Head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			return true
		}
	}
	return false
}

/**
 获取环形链表的入口节点
 */
func DetectCycle(l *MyLinkedList) *ListNode {
	// 用双指针法，先找到两个指针相遇的点
	fast := l.Head
	slow := l.Head
	var counter *ListNode
	for slow != nil && fast != nil && fast.Next !=nil{
		// fast走两步
		fast = fast.Next.Next
		// slow走一步
		slow = slow.Next
		if slow == fast{
			counter = slow
			break
		}
	}
	if counter == nil{
		return nil
	}
	// 相遇之后，计算得出，s两指针分别从链表头和相遇点开始，每次走一步，相遇的点即为环入口
	p1 := l.Head
	p2 := counter
	for p1 != p2{
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

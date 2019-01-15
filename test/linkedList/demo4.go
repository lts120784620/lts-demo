package linkedList

/**
 求链表的中间节点
 */
func MiddleNode(l *MyLinkedList) *Node{
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

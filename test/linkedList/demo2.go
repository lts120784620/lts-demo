package linkedList

// 翻转链表,非递归
func ReverseList(l *MyLinkedList) *Node {
	if l.Head == nil || l.Head.Next == nil{
		return l.Head
	}
	prev := l.Head
	cur := l.Head.Next
	temp := cur.Next
	for cur != nil{
		//fmt.Println("p:",prev.Data,"c:",cur.Data,"t:",temp.Data)
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	l.Head.Next = nil
	return prev
}

// 翻转链表，递归
func ReverseList2(head *Node) *Node {
	if head == nil || head.Next == nil{
		return head
	}
	n := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next  = nil
	return n
}



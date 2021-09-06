package linkedlist

/**
删除给定排序链表中的所有重复元素
*/
func DeleteDuplicates(l *MyLinkedList) *ListNode {
	head := l.Head.Next
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head.Next
	for p1 != nil && p2 != nil {
		if p2.Data == p1.Data {
			p1.Next = p2.Next
			p2 = p1.Next
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return head
}

/**
删除链表中相同的所有节点
*/
func RemoveElements(l *MyLinkedList, val string) *ListNode {
	head := l.Head.Next
	for head != nil && head.Data == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	pre := head
	cur := pre.Next
	for cur != nil {
		if cur.Data == val {
			pre.Next = cur.Next
			cur = pre.Next
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}
	return head
}

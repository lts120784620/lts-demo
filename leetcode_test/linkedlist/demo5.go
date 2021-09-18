package linkedlist

/**
判断两个链表是否相交（无环）
*/
func IsIntersect(l *ListNode, r *ListNode) *ListNode {
	n1 := l.Next
	for n1.Next != nil {
		n1 = n1.Next
	}
	n2 := r.Next
	for n2.Next != nil {
		n2 = n2.Next
	}
	if n1 == n2 {
		return n2
	}
	return nil
}

/**
获取两个链表第一次相交的节点
*/
func GetIntersectionNode(l *ListNode, r *ListNode) *ListNode {
	// 计算两条链表长度
	var len1, len2 int
	p1 := l.Next
	p2 := r.Next
	for p1 != nil {
		p1 = p1.Next
		len1++
	}
	for p2 != nil {
		p2 = p2.Next
		len2++
	}
	// 长链表先走len1-len2个，同时遍历，相同的节点就是相交的点
	p1 = l.Next
	p2 = r.Next
	count := 0
	if len2 > len1 {
		for len2-len1 > count {
			count++
			p2 = p2.Next
		}
	} else {
		for len1-len2 > count {
			count++
			p1 = p1.Next
		}
	}
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}

/**
获取两个链表第一次相交的节点(更优)
*/
func GetIntersectionNode2(l *ListNode, r *ListNode) *ListNode {
	p1 := l.Next
	p2 := r.Next
	if p1 == nil || p2 == nil {
		return nil
	}
	// 每次比较两个指针，若到链表结尾则指向另一个链表的头结点，这样一次结束会抹除长度差，第二次则会在相同位置开始遍历
	for p1 != p2 {
		if p1 == nil {
			p1 = r.Next
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = l.Next
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

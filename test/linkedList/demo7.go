package linkedList

//todo
/**
 删除给定排序链表中的所有重复元素
 */
func DeleteDuplicates(l *MyLinkedList) *Node {
	head := l.Head.Next
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	p1 := head
	p2 := head.Next
	for p2 != nil{
		if p2.Data == p1.Data{
			p1.Next = p2.Next
			if p1.Next == nil{
				p2= nil
				break
			}else{
				p2 = p1.Next.Next
			}
		}else{
			p2 = p2.Next
		}
		p1 = p1.Next
	}
	return head
}

/**
 删除链表中相同的所有节点
 */
func RemoveElements(l *MyLinkedList,val string) *Node{
	return nil
}

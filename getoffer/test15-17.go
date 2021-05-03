package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func main() {
	//f := &ListNode{6, nil}
	//e := &ListNode{5, f}
	//d := &ListNode{4, e}
	//c := &ListNode{3, d}
	//b := &ListNode{2, c}
	//a := &ListNode{1, b}
	//head := &ListNode{0, a}

	// 15
	//fmt.Println(findKthToTail(head,1).value)

	// 16
	//fmt.Println(recursiveReverseList(head))

	d1 := &ListNode{7, nil}
	c1 := &ListNode{5, d1}
	b1 := &ListNode{3, c1}
	a1 := &ListNode{1, b1}
	p1 := &ListNode{0, a1}

	d2 := &ListNode{8, nil}
	c2 := &ListNode{6, d2}
	b2 := &ListNode{4, c2}
	a2 := &ListNode{2, b2}
	p2 := &ListNode{0, a2}

	// 17
	printLinkedList(merge(p1, p2))
}

func merge(p1 *ListNode, p2 *ListNode) *ListNode {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	var head *ListNode
	if p1.value < p2.value {
		head = p1
		p1.next = merge(p1.next, p2)
	} else {
		head = p2
		p2.next = merge(p1, p2.next)
	}

	return head
}

func findKthToTail(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return nil
	}
	p1 := head
	for i := 0; i < k-1; i++ {
		if p1.next == nil {
			return nil
		} else {
			p1 = p1.next
		}
	}

	p2 := head
	for p1.next != nil {
		p1 = p1.next
		p2 = p2.next
	}

	return p2
}

// 反转链表（非递归）
func reverseList(head *ListNode) *ListNode {
	var reverseHead *ListNode
	var pre *ListNode
	p := head
	for p != nil {
		next := p.next
		if next == nil {
			reverseHead = p
		}
		p.next = pre
		pre = p
		p = next
	}

	return reverseHead
}

// 反转链表（递归）
func recursiveReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return head
	}
	n := recursiveReverseList(head.next)
	head.next.next = head
	head.next = nil
	return n
}

func printLinkedList(head *ListNode) {
	for p := head; p != nil; p = p.next {
		if p.next != nil{
			fmt.Printf("%d -> ",p.value)
		}else {
			fmt.Printf("%d ",p.value)
		}
	}
}

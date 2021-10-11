package linkedlist

import "fmt"

/**
链表基础
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList interface {
	Length() int
	Insert(idx int, n *ListNode)
	Delete(idx int)
	Get(idx int) *ListNode
	String() string
}

func New() *ListNode {
	// 返回头节点
	return &ListNode{0, nil}
}

func NewForArrays(arr []interface{}) *ListNode {
	head := &ListNode{0, nil}
	for _, v := range arr {
		n := &ListNode{v.(int), nil}
		head.Add(n)
	}
	head = head.Next
	fmt.Printf("链表初始化完成:%s\n", head.String())
	return head
}

func (l *ListNode) Length() int {
	if l == nil {
		return 0
	}
	n := l
	i := 0
	for n != nil {
		i++
		n = n.Next
	}
	return i
}

func (l *ListNode) Add(node *ListNode) {
	if node == nil {
		return
	}

	n := l
	i := 0
	for n != nil && n.Next != nil {
		i++
		n = n.Next
	}
	n.Next = node
}

func (l *ListNode) Insert(idx int, node *ListNode) {
	if idx < 0 && node == nil {
		return
	}
	n := l.Next
	i := 0
	for i < idx-1 && n != nil {
		i++
		n = n.Next
	}
	node.Next = n.Next
	n.Next = node
}

func (l *ListNode) Get(idx int) *ListNode {
	if idx < 0 {
		return nil
	}
	n := l
	i := 0
	for i < idx && n != nil {
		i++
		n = n.Next
	}
	return n
}

func (l *ListNode) Delete(idx int) {
	if idx < 0 || idx > l.Length() {
		return
	}
	n := l
	i := 0
	for i < idx-1 && n != nil {
		i++
		n = n.Next
	}
	n.Next = n.Next.Next
}

func (l *ListNode) String() string {
	n := l
	s := fmt.Sprintf("长度:%d, head->", l.Length())
	for n != nil {
		s = fmt.Sprintf("%s%d", s, n.Val)
		if n.Next != nil {
			s = fmt.Sprintf("%s->", s)
		}
		n = n.Next
	}
	return s
}

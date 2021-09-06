package linkedlist

import "fmt"

/**
链表基础
*/

type ListNode struct {
	Data string
	Next *ListNode
}

type MyLinkedList struct {
	Head *ListNode
	size int
}

type linkedList interface {
	Length() int
	Insert(idx int, n *ListNode)
	Delete(idx int)
	Get(idx int) *ListNode
	Traverse()
}

var l = NewInitLinkedList(5)

func NewLinkedList() *MyLinkedList {
	return &MyLinkedList{&ListNode{"head", nil}, 0}
}

func NewInitLinkedList(size int) *MyLinkedList {
	t := []string{"A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N"}
	l := &MyLinkedList{&ListNode{"head", nil}, 0}
	for i := 0; i < size; i++ {
		n := &ListNode{t[i], nil}
		l.Add(n)
	}
	fmt.Print("链表初始化为：")
	TravelseHead(l.Head)
	return l
}

func (l *MyLinkedList) Length() int {
	n := l.Head.Next
	i := 0
	for n != nil {
		i++
		n = n.Next
	}
	return i
}

func (l *MyLinkedList) Add(node *ListNode) {
	if node == nil {
		return
	}
	n := l.Head
	i := 0
	for i < l.size && n != nil {
		i++
		n = n.Next
	}
	if n == nil {
		return
	}
	n.Next = node
	l.size++
}

func (l *MyLinkedList) Insert(idx int, node *ListNode) {
	if idx < 0 && node == nil {
		return
	}
	n := l.Head
	i := 0
	for i < idx-1 && n != nil {
		i++
		n = n.Next
	}
	node.Next = n.Next
	n.Next = node
	l.size++
}

func (l *MyLinkedList) Get(idx int) *ListNode {
	if idx < 0 {
		return nil
	}
	n := l.Head
	i := 0
	for i < idx && n != nil {
		i++
		n = n.Next
	}
	return n
}

func (l *MyLinkedList) Delete(idx int) {
	if idx < 0 || idx > l.Length() {
		return
	}
	n := l.Head
	i := 0
	for i < idx-1 && n != nil {
		i++
		n = n.Next
	}
	n.Next = n.Next.Next
}

func (l *MyLinkedList) Travelse() {
	n := l.Head
	i := 0
	for n != nil {
		fmt.Print(n.Data)
		if n.Next != nil {
			fmt.Print("-> ")
		}
		n = n.Next
		i++
	}
	fmt.Println()
}

func TravelseHead(head *ListNode) {
	n := head
	i := 0
	for n != nil {
		fmt.Print(n.Data)
		if n.Next != nil {
			fmt.Print("-> ")
		}
		n = n.Next
		i++
	}
	fmt.Println()
}

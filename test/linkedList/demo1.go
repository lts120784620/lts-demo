package linkedList

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	size int
}

type linkedList interface {
	length() int
	Insert(idx int, n *Node)
	Delete(idx int)
	get(idx int) *Node
	Traverse()
}

func New() *MyLinkedList {
	return &MyLinkedList{&Node{"head", nil}, 0}
}

func (l *MyLinkedList) length() int {
	n := l.Head.Next
	i := 0
	for n != nil {
		i++
		n = n.Next
	}
	return i
}

func (l *MyLinkedList) Insert(idx int, node *Node) {
	if idx < 0 && node == nil {
		return
	}
	n := l.Head
	i := 0
	for i < idx - 1 && n != nil {
		i++
		n = n.Next
	}
	node.Next = n.Next
	n.Next = node
}
func (l *MyLinkedList) get(idx int) *Node {
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
	if idx < 0 || idx > l.length() {
		return
	}
	n := l.Head
	i := 0
	for i < idx - 1 && n != nil {
		i++
		n = n.Next
	}
	n.Next = n.Next.Next
}

func (l *MyLinkedList) Travelse() {
	n := l.Head
	i := 0
	for n != nil {
		fmt.Print(i, ":", n.Data)
		fmt.Println()
		n = n.Next
		i++
	}
}

func TravelseHead(head *Node) {
	n := head
	i := 0
	for n != nil {
		fmt.Print(i, ":", n.Data)
		fmt.Println()
		n = n.Next
		i++
	}
}

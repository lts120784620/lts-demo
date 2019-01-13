package main

import "fmt"

type node struct {
	data string
	next *node
}

type myLinkedList struct {
	head *node
	size int
}

type linkedList interface {
	length() int
	insert(idx int,n *node)
	delete(idx int)
	get(idx int) *node
	Traverse()
}

func New() *myLinkedList {
	return &myLinkedList{&node{"head",nil},0}
}

func (l *myLinkedList) length() int {
	n := l.head.next
	i := 0
	for n != nil{
		i++
		n = n.next
	}
	return i
}

func (l *myLinkedList) insert(idx int,node *node) {
	if idx <0 && node == nil{
		return
	}
	n := l.head
	i := 0
	for i < idx-1 && n != nil{
		i++
		n = n.next
	}
	node.next = n.next
	n.next = node
}
func (l *myLinkedList) get(idx int) *node {
	if idx <0{
		return nil
	}
	n := l.head
	i := 0
	for i<idx && n != nil{
		i++
		n = n.next
	}
	return n
}

func (l *myLinkedList) delete(idx int) {
	if idx <0 || idx > l.length(){
		return
	}
	n := l.head
	i := 0
	for i < idx -1 && n != nil{
		i++
		n = n.next
	}
	n.next = n.next.next
}

func (l *myLinkedList) Travelse() {
	n := l.head
	i := 0
	for n != nil{
		fmt.Print(i,":",n.data)
		fmt.Println()
		n = n.next
		i++
	}
}

func main()  {
	l := New()
	l.insert(1,&node{"a",nil})
	l.insert(2,&node{"b",nil})
	l.Travelse()
	fmt.Println(l.get(0).data)
	fmt.Println(l.get(2).data)
	fmt.Println(l.length())
	l.delete(1)
	l.Travelse()
}
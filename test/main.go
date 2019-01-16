package main

import (
	"lts-demo/test/linkedList"
	"fmt"
)

func main() {
	l := linkedList.New()
	a := &linkedList.Node{"a", nil}
	b := &linkedList.Node{"b", nil}
	c := &linkedList.Node{"c", nil}
	d := &linkedList.Node{"d", nil}
	g := &linkedList.Node{"g", nil}
	l.Insert(1, a)
	l.Insert(2, b)
	l.Insert(3, c)
	l.Insert(4, d)
	l.Insert(5, g)
	l.Travelse()

	r := linkedList.New()
	e := &linkedList.Node{"e", nil}
	f := &linkedList.Node{"f", nil}
	r.Insert(1, e)
	r.Insert(2, f)
	r.Insert(3, d)
	r.Insert(4, g)
	r.Travelse()

	//翻转链表
	//head := linkedList.ReverseList(l)
	//head2 := linkedList.ReverseList2(l.Head)
	//linkedList.TravelseHead(head)
	//linkedList.TravelseHead(head2)

	// 删除倒数k个节点
	//fmt.Println(linkedList.GetNthFromEnd(l,1).Data)

	//fmt.Println(linkedList.MiddleNode(l).Data)

	//fmt.Println(linkedList.HasCycle(l))

	//fmt.Println(linkedList.IsIntersect(l,r).Data)

	fmt.Println(linkedList.GetIntersectionNode(l,r).Data)
	fmt.Println()
}

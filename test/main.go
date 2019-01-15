package main

import (
	"lts-demo/test/linkedList"
	"fmt"
)

func main() {
	l := linkedList.New()
	a := &linkedList.Node{"a", nil}
	l.Insert(1, a)
	l.Insert(2, &linkedList.Node{"b", nil})
	l.Insert(3, &linkedList.Node{"c", nil})
	d := &linkedList.Node{"d", nil}
	l.Insert(4, d)
	d.Next = a
	//l.Travelse()

	//翻转链表
	//head := linkedList.ReverseList(l)
	//head2 := linkedList.ReverseList2(l.Head)
	//linkedList.TravelseHead(head)
	//linkedList.TravelseHead(head2)

	// 删除倒数k个节点
	//fmt.Println(linkedList.GetNthFromEnd(l,1).Data)

	//fmt.Println(linkedList.MiddleNode(l).Data)

	fmt.Println(linkedList.HasCycle(l))
	fmt.Println()
}

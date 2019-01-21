package main

import (
	"lts-demo/test/linkedList"
	"fmt"
)

func main() {
	//l := linkedList.New()
	//a := &linkedList.Node{"a", nil}
	//b := &linkedList.Node{"b", nil}
	//c := &linkedList.Node{"c", nil}
	//d := &linkedList.Node{"d", nil}
	//g := &linkedList.Node{"g", nil}
	//l.Insert(1, a)
	//l.Insert(2, b)
	//l.Insert(3, c)
	//l.Insert(4, d)
	//l.Insert(5, g)
	//l.Travelse()
	//
	//r := linkedList.New()
	//e := &linkedList.Node{"e", nil}
	//f := &linkedList.Node{"f", nil}
	//r.Insert(1, e)
	//r.Insert(2, f)
	//r.Insert(3, d)
	//r.Insert(4, g)
	//r.Travelse()


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

	//fmt.Println(linkedList.GetIntersectionNode(l,r).Data)

	//求环形入口节点
	//g.Next = r.Head
	//fmt.Println(linkedList.DetectCycle(l).Data)

	// 合并两个有序链表
	//l1 := linkedList.New()
	//a := &linkedList.Node{"1", nil}
	//b := &linkedList.Node{"3", nil}
	//c := &linkedList.Node{"5", nil}
	//l1.Insert(1, a)
	//l1.Insert(2, b)
	//l1.Insert(3, c)
	//l2 := linkedList.New()
	//a1 := &linkedList.Node{"2", nil}
	//b1 := &linkedList.Node{"4", nil}
	//c1 := &linkedList.Node{"6", nil}
	//l2.Insert(1, a1)
	//l2.Insert(2, b1)
	//l2.Insert(3, c1)
	//linkedList.MergeTwoLists(l1,l2)

	// 判断回文链表
	l := linkedList.New()
	a := &linkedList.Node{"0", nil}
	b := &linkedList.Node{"0", nil}
	l.Insert(1, a)
	l.Insert(2, b)

	fmt.Println(linkedList.IsPalindrome(l))
	fmt.Println()
}

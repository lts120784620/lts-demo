package main

import (
	"lts-demo/test/linkedlist"
	"fmt"
	"github.com/toolkits/time"
	"strconv"
)

func main() {

	fmt.Println(len(strconv.Itoa(1551425031780)))
	fmt.Println(time.FormatTs(1551425031780))
	l := linkedlist.New()
	//a := &linkedlist.ListNode{"a", nil}
	//b := &linkedlist.ListNode{"b", nil}
	//c := &linkedlist.ListNode{"c", nil}
	//d := &linkedlist.ListNode{"d", nil}
	//g := &linkedlist.ListNode{"g", nil}
	//l.Insert(1, a)
	//l.Insert(2, b)
	//l.Insert(3, c)
	//l.Insert(4, d)
	//l.Insert(5, g)
	l.Travelse()
	//
	//r := linkedlist.New()
	//e := &linkedlist.ListNode{"e", nil}
	//f := &linkedlist.ListNode{"f", nil}
	//r.Insert(1, e)
	//r.Insert(2, f)
	//r.Insert(3, d)
	//r.Insert(4, g)
	//r.Travelse()


	//翻转链表
	//head := linkedlist.ReverseList(l)
	//head2 := linkedlist.ReverseList2(l.Head)
	//linkedlist.TravelseHead(head)
	//linkedlist.TravelseHead(head2)

	// 删除倒数k个节点
	//fmt.Println(linkedlist.GetNthFromEnd(l,1).Data)

	//fmt.Println(linkedlist.MiddleNode(l).Data)

	//fmt.Println(linkedlist.HasCycle(l))

	//fmt.Println(linkedlist.IsIntersect(l,r).Data)

	//fmt.Println(linkedlist.GetIntersectionNode(l,r).Data)

	//求环形入口节点
	//g.Next = r.Head
	//fmt.Println(linkedlist.DetectCycle(l).Data)

	// 合并两个有序链表
	//l1 := linkedlist.New()
	//a := &linkedlist.ListNode{"1", nil}
	//b := &linkedlist.ListNode{"3", nil}
	//c := &linkedlist.ListNode{"5", nil}
	//l1.Insert(1, a)
	//l1.Insert(2, b)
	//l1.Insert(3, c)
	//l2 := linkedlist.New()
	//a1 := &linkedlist.ListNode{"2", nil}
	//b1 := &linkedlist.ListNode{"4", nil}
	//c1 := &linkedlist.ListNode{"6", nil}
	//l2.Insert(1, a1)
	//l2.Insert(2, b1)
	//l2.Insert(3, c1)
	//linkedlist.MergeTwoLists(l1,l2)

	// 判断回文链表
	//l := linkedlist.New()
	//a := &linkedlist.ListNode{"0", nil}
	//b := &linkedlist.ListNode{"0", nil}
	//l.Insert(1, a)
	//l.Insert(2, b)
	//
	//fmt.Println(linkedlist.IsPalindrome(l))

	//l := linkedlist.New()
	//a := &linkedlist.ListNode{"a", nil}
	//b := &linkedlist.ListNode{"b", nil}
	//c := &linkedlist.ListNode{"b", nil}
	//d := &linkedlist.ListNode{"b", nil}
	//d := &linkedlist.ListNode{"1", nil}
	//l.Insert(1, a)
	//l.Insert(2, b)
	//l.Insert(3, c)
	//l.Insert(4, d)
	//linkedlist.DeleteDuplicates(l)

	//linkedlist.RemoveElements(l,"a")
	//l.Travelse()

	fmt.Println()
}

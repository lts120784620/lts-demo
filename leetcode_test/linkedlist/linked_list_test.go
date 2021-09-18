package linkedlist

import (
	"fmt"
	"testing"
)

/**
基础链表测试
*/

func TestInit(t *testing.T) {
	fmt.Println(NewLinkedList())
}

func TestNewInitLinkedListByArray(t *testing.T) {
	l := NewInitLinkedListByArray([]interface{}{1, 2, 3, 4, 5, 6})
	fmt.Println(l)
}

func TestListNode_Length(t *testing.T) {
	l := NewInitLinkedListByArray([]interface{}{1, 2, 3, 4, 5, 6})
	fmt.Println(l.Length())
	l.Add(&ListNode{Val: 7})
	fmt.Println(l.Length())
}

func TestListNode_Add(t *testing.T) {
	l := NewLinkedList()
	l.Add(&ListNode{Val: 1})
	l.Add(&ListNode{Val: 2})
	l.Add(&ListNode{Val: 3})
	fmt.Println(l)
}

func TestListNode_Insert(t *testing.T) {
	l := NewInitLinkedListByArray([]interface{}{1, 2, 3, 5})
	l.Insert(3, &ListNode{Val: 4})
	fmt.Println(l)
}

func TestListNode_Delete(t *testing.T) {
	l := NewInitLinkedListByArray([]interface{}{1, 2, 3, 3, 4})
	l.Delete(4)
	fmt.Println(l)
}

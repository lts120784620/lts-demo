package linkedlist

import (
	"container/heap"
	"fmt"
	_heap "lts-demo/leetcode_test/heap"
	"testing"
)

/**
No.23. 合并K个升序链表
描述：
	给你一个链表数组，每个链表都已经按升序排列。
	请你将所有链表合并到一个升序链表中，返回合并后的链表。
思路：
    1、利用优先级队列（小顶堆），思路和合并两个链表类似，寻找k个节点头是通过堆输出
时间：
    1、2021/10/9
*/

func TestMergeKLists(t *testing.T) {
	list1 := NewForArrays([]interface{}{1, 1, 2})
	list2 := NewForArrays([]interface{}{1, 2, 2})

	lists := []*ListNode{list2, list1}
	fmt.Println(mergeKLists(lists).String())
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	// 优先级队列
	pq := _heap.New()
	for _, list := range lists {
		if list != nil {
			item := &_heap.Item{
				Value:    list,
				Priority: list.Val,
			}
			heap.Push(&pq, item)
		}
	}

	// 遍历队列
	for pq.Len() != 0 {
		// 弹出当前最小的
		node := heap.Pop(&pq).(*_heap.Item).Value.(*ListNode)
		if node == nil {
			continue
		}
		// 合并链表
		p.Next = node
		if node.Next != nil {
			item := &_heap.Item{
				Value:    node.Next,
				Priority: node.Next.Val,
			}
			heap.Push(&pq, item)
		}
		p = p.Next
	}
	return dummy.Next
}

package heap

import (
	"container/heap"
	"fmt"
	"lts-demo/test/queue"
)

/**
思路：优先级队列（堆实现）
*/

type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func New(items map[interface{}]int) PriorityQueue {
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	// 初始化堆
	heap.Init(&pq)
	return pq
}

func (pq PriorityQueue) Len() int { return len(pq) }

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Peek() interface{} {
	return (*pq)[0]
}

func (pq *PriorityQueue) PrintHeap() {
	size := (*pq).Len()
	// 找头结点，打印
	children := queue.New()
	children.Enqueue(0)
	for !children.IsEmpty() {
		var column []int
		cLen := children.Len()
		for i := 0; i < cLen; i++ {
			n := children.Dequeue().(int)
			// 打印根节点
			//fmt.Println((*pq)[n].value)
			column = append(column, (*pq)[n].value.(int))
			if 2*n+1 > size-1 {
				continue
			}
			// 放左孩子
			children.Enqueue(2*n + 1)
			if 2*n+2 > size-1 {
				continue
			}
			// 放右孩子
			children.Enqueue(2*n + 2)
		}
		fmt.Println(column)
	}
	fmt.Println("*********************")
}

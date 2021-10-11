package heap

import (
	"container/heap"
	"fmt"
	"lts-demo/leetcode_test/queue"
)

/**
思路：优先级队列（堆实现）
*/

type Item struct {
	Value    interface{} // The Value of the item; arbitrary.
	Priority int         // The Priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
	Index int // The Index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func New() PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return pq
}

func NewForMap(items map[interface{}]int) PriorityQueue {
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	// 初始化堆
	heap.Init(&pq)
	return pq
}

func (pq PriorityQueue) Len() int { return len(pq) }

// update modifies the Priority and Value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
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
			//fmt.Println((*pq)[n].Value)
			column = append(column, (*pq)[n].Value.(int))
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

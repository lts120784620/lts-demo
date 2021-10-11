package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

/**
No.703 数据流中的第 K 大元素
思路：
    1、采用heap的结构，创建一个小顶堆，根为最小值，每次增加数据后和堆顶值比较，如果大于则把堆顶值弹出，写入新值并调堆，弹出最小的，保证堆中剩最大的
时间复杂度:数据长度N * 调堆logK
*/

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

type KthLargest struct {
	Pq PriorityQueue
	K  int
}

func Constructor(k int, nums []int) KthLargest {
	m := map[interface{}]int{}
	for _, n := range nums {
		m[n] = n
	}
	kHeap := KthLargest{NewForMap(m), k}
	return kHeap
}

func (this *KthLargest) Add(val int) int {
	if this.Pq.Len() < this.K {
		item := &Item{
			Value:    val,
			Priority: val,
		}
		heap.Push(&this.Pq, item)
	} else if this.Pq.Peek().(*Item).Value.(int) < val {
		heap.Pop(&this.Pq)
		item := &Item{
			Value:    val,
			Priority: val,
		}
		heap.Push(&this.Pq, item)
	}
	return this.Pq.Peek().(*Item).Value.(int)
}

func TestKthLargest(t *testing.T) {
	nums := []int{4, 5, 8, 2, 1}
	k := 3
	kth := Constructor(k, nums)
	kth.Pq.PrintHeap()
	fmt.Println("第k大的值为：", kth.Add(3))
	kth.Pq.PrintHeap()
	fmt.Println("第k大的值为：", kth.Add(5))
	kth.Pq.PrintHeap()
	fmt.Println("第k大的值为：", kth.Add(10))
	kth.Pq.PrintHeap()
	fmt.Println("第k大的值为：", kth.Add(9))
	kth.Pq.PrintHeap()
	fmt.Println("第k大的值为：", kth.Add(4))
	kth.Pq.PrintHeap()

}

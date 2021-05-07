package heap

import (
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
	nItems := map[int]interface{}{}
	for _, j := range nums {
		nItems[j] = j
	}
	return KthLargest{
		New(nItems), k,
	}
}

func (this *KthLargest) Add(val int) int {
	this.Pq.Push(val)
	if this.Pq.Len() > this.K{
		this.Pq.Pop()
	}
	// todo list
	// peek方法
	return this.Pq.
}

func TestKthLargest(t *testing.T) {
	nums := []int{}
	k := 3
	kth := Constructor(k, nums)
	fmt.Println(kth.Add(5))
	fmt.Println(kth.Add(3))
	fmt.Println(kth.Add(1))
	fmt.Println(kth.Add(6))
}

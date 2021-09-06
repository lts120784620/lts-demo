package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

/**
No.239 滑动窗口最大值
描述：
	输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
	输出：[3,3,5,5,6,7]
	解释：
	滑动窗口的位置                最大值
	---------------               -----
	[1  3  -1] -3  5  3  6  7       3
	 1 [3  -1  -3] 5  3  6  7       3
	 1  3 [-1  -3  5] 3  6  7       5
	 1  3  -1 [-3  5  3] 6  7       5
	 1  3  -1  -3 [5  3  6] 7       6
	 1  3  -1  -3  5 [3  6  7]      7

思路：
    1、使用小顶堆的方式
	2、使用双端队列的方式

todo 未做完

*/

func TestName(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow1(nums, k))
}

// 创建优先级队列
func maxSlidingWindow1(nums []int, k int) []int {
	res := []int{}
	pq := New(map[interface{}]int{})
	for _, n := range nums {
		var top *Item
		if pq.Len() < k {
			item := Item{
				value:    n,
				priority: -n,
			}
			heap.Push(&pq, item)
			top = pq.Peek().(*Item)
			continue
		}
		// 注意这里要构造大顶堆，即当前大于堆顶最大的元素，做入栈出栈操作
		heap.Push(&pq, n)

		top = pq.Peek().(*Item)
		res = append(res, top.value.(int))
	}
	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	res := []int{}

	return res
}

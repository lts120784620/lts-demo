package sort

import (
	"fmt"
	"testing"

	"lts-demo/leetcode_test/unit"
)

/**
No.
描述：
	堆排序
思路：
    1、通常用来解决topK问题，时间复杂度O(n^2)
时间：
    1、2021/9/17
*/

func TestHeapSort(t *testing.T) {
	arr := unit.RandomArrays(10)
	fmt.Println(arr)
	fmt.Println(">>>>>>>>>>>>>>>>")
	heapSort(arr)
	fmt.Println(arr)
}

func heapSort(arr []int) {
	// 建堆
	// 从最后一个非叶子节点开始
	for i := len(arr)/2 - 1; i >= 0; i-- {
		fixHeap(arr, i, len(arr))
	}
	fmt.Println(arr)
	// 调堆
	for i := len(arr) - 1; i > 0; i-- {
		unit.ArraysSwap(arr, 0, i)
		fixHeap(arr, 0, i)
	}
}

func fixHeap(arr []int, i, n int) {
	t := arr[i]
	j := 2*i + 1 // i的左子节点
	for j < n {
		for j+1 < n && arr[j+1] < arr[j] {
			j++
		}
		if arr[j] >= t {
			break
		}
		arr[i] = arr[j]
		i = j
		j = 2*i + 1
	}
	arr[i] = t
}

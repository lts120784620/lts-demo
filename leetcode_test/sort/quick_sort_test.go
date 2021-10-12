package sort

import (
	"fmt"
	"testing"

	"lts-demo/leetcode_test/unit"
)

/**
No.
描述：
	快速排序
思路：二叉树的前序遍历
    1、快速排序的核心也是递归+分治的思路，和归并排序类似,不需要归并使用额外的空间
	2、第一部分，取左边0号为作为pivot基准值，两个指针left+1、right从前后分别遍历，如果不符合左边小于pivot，右边大雨pivot，则两边交换
	3、重点是结束时交换pivot标志位，返回中间位置给递归使用
	4、递归left~mid-1,mid+1~right，结束条件为left >= right
	3、时间复杂度O(n*logn)
时间：
	1、--
    2、2021/9/17
*/

func TestQuickSort(t *testing.T) {
	arr := unit.RandomArrays(10)
	fmt.Println(arr)
	fmt.Println(">>>>>>>>>>>>>>>>")
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// second
func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	// 中间划分
	mid := partition(arr, start, end)
	quickSort(arr, start, mid-1)
	quickSort(arr, mid+1, end)
}

func partition(arr []int, left, right int) int {
	p := left
	pivot := arr[p]
	start, end := left+1, right
	// 开始交换排列大小
	for start < end {
		for arr[start] < pivot && start < len(arr) {
			start++
		}

		for arr[end] > pivot && end > 0 {
			end--
		}

		if start < end && end > 0 && start < len(arr) {
			unit.ArraysSwap(arr, start, end)
		}
	}
	// 核心，返回中间位置
	unit.ArraysSwap(arr, end, p)
	return end
}

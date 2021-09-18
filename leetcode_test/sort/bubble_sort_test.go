package sort

import (
	"fmt"
	"lts-demo/leetcode_test/unit"
	"testing"
)

/**
No.
描述：
	冒泡排序
思路：
    1、两遍循环，每次比较相邻两个元素，交换位置，每次遍历最大的元素会排在最后，时间复杂度O(n^2)
时间：
    1、2021/9/16
*/

func TestBubbleSort(t *testing.T) {
	arr := unit.RandomArrays(10)
	fmt.Println(arr)
	fmt.Println(">>>>>>>>>>>>>>>>")
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				unit.ArraysSwap(arr, j, j+1)
			}
		}
	}
}

func bubbleSort1(arr []int) {
	// 如果是倒序排好的，会不会浪费
}

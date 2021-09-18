package sort

import (
	"fmt"
	"lts-demo/leetcode_test/unit"
	"testing"
)

/**
No.
描述：
	选择排序
思路：
    1、两次遍历，外层0~n-1,内层找到最小值进行交换，每次都在剩余未排序的值中确定一个最小值，时间复杂度O(n^2)
时间：
    1、2021/9/16
*/

func TestSelectSort(t *testing.T) {
	arr := unit.RandomArrays(10)
	fmt.Println(arr)
	fmt.Println(">>>>>>>>>>>>>>>>")
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// idx保存最小值的下标
		idx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[idx] {
				idx = j
			}
		}
		unit.ArraysSwap(arr, i, idx)
	}
}

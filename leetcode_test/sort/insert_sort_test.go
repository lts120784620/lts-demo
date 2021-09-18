package sort

import (
	"fmt"
	"lts-demo/leetcode_test/unit"
	"testing"
)

/**
No.
描述：
	插入排序
思路：
    1、两层循环，第一层从0～n-1,第二层判断如果当前值小于前半部分的有序数组，则插入指定位置，其余有序向后移动，时间复杂度O(n^2)
时间：
    1、2021/9/16
*/

func TestInsertSort(t *testing.T) {
	arr := unit.RandomArrays(5)
	fmt.Println(arr)
	fmt.Println(">>>>>>>>>>>>>>>>")
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		t := arr[i]
		j := i - 1
		// 判断当前值应该插入到前面有序部分的那个位置,直到0或者当前值大于等于
		for j >= 0 && t < arr[j] {
			// i 小于前面的值，前面的往后挪位置
			arr[j+1] = arr[j]
			j--
		}
		// 插入位置，这里+1是因为开始j=i-1，考虑只有一个值或者i最大的情况
		arr[j+1] = t
		fmt.Println(arr)
	}
}

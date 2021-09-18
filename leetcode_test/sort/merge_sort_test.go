package sort

import (
	"fmt"
	"lts-demo/leetcode_test/unit"
	"testing"
)

/**
No.
描述：
	归并排序
思路：
    1、归并是典型的分治思路，归并的核心思路是如何将两个有序的数组合并成一个有序的大数组
	2、首先通过递归的方式，将数组分成前后两个小部分，直到两个小部分为单值，
	3、在递归划分内部实现将两个小数组合并成一个大数组
	4、需要利用额外空间，空间复杂度O(n)，时间复杂度O(n*logn)
时间：
    1、2021/9/17
*/

func TestMergeSort(t *testing.T) {
	arr := unit.RandomArrays(10)
	fmt.Println(arr)
	fmt.Println(">>>>>>>>>>>>>>>>")
	a := mergeSort(arr)
	fmt.Println(a)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[0:mid]), mergeSort(arr[mid:]))
}

func merge(l1, l2 []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(l1) && j < len(l2) {
		if l1[i] <= l2[j] {
			res = append(res, l1[i])
			i++
		} else {
			res = append(res, l2[j])
			j++
		}
	}

	for i < len(l1) {
		res = append(res, l1[i])
		i++
	}

	for j < len(l2) {
		res = append(res, l2[j])
		j++
	}
	return res
}

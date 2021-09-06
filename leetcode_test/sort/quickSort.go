package main

import "fmt"

func main() {
	arr := []int{3, 2, 5, 6, 4, 1, 9, 7, 8}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, start, end int) {
	if arr == nil {
		return
	}
	if start < end {
		mid := partition(arr, start, end)
		quickSort(arr, start, mid)
		quickSort(arr, mid+1, end)
	}
}

func partition(arr []int, start, end int) int {
	if arr == nil {
		return 0
	}
	pivot := arr[start]
	i, j := start, end
	for i < j {
		// 因为pivot是左边，所以要从大的右边开始判断，这样最后相遇时，i会在小的位置和start交换
		for i < j && arr[j] > pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		Swap(arr, i, j)
	}
	// 将标志位置和相遇位置，交换
	Swap(arr, start, i)
	return i
}

//交换
func Swap(arr []int, x, y int) {
	t := arr[x]
	arr[x] = arr[y]
	arr[y] = t
}

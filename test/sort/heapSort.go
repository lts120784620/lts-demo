package main

import "fmt"

func main() {
	arr := []int{3, 2, 5, 6, 4, 1, 9, 7, 8}
	heapSort(arr)
	fmt.Println(arr)
}

func heapSort(arr []int) {
	// 建堆
	// 从最后一个非叶子节点开始
	for i := len(arr) / 2 - 1; i >= 0; i-- {
		fmt.Println(i)
		fixHeap(arr, i, len(arr))
	}
	// 调堆
	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)
		fixHeap(arr, 0, i)
	}
}

func fixHeap(arr []int, i, n int) {
	t := arr[i]
	j := 2 * i + 1 // i的左子节点
	for j < n {
		for j+1 < n && arr[j + 1] < arr[j] {
			j++
		}
		if arr[j] >= t{
			break
		}
		arr[i] = arr[j]
		i = j
		j = 2*i+1
	}
	arr[i] = t
}

//交换
func swap(arr []int, x, y int) {
	t := arr[x]
	arr[x] = arr[y]
	arr[y] = t
}

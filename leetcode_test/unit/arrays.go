package unit

/**
数组通用操作
*/

//交换
func ArraysSwap(arr []int, x, y int) {
	t := arr[x]
	arr[x] = arr[y]
	arr[y] = t
}

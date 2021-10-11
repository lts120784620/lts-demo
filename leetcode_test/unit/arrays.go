package unit

import "math"

/**
数组通用操作
*/

//交换
func ArraysSwap(arr []int, x, y int) {
	t := arr[x]
	arr[x] = arr[y]
	arr[y] = t
}

func ArraysMax(arr []int) int {
	max := math.MinInt64
	for _, v := range arr {
		max = int(math.Max(float64(v), float64(max)))
	}
	return max
}

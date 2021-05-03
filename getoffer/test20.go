package main

import "fmt"

func main() {
	arr := [][]int{}
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	c := []int{9, 10, 11, 12}
	d := []int{13, 14, 15, 16}
	arr = append(arr, a, b, c, d)

	printArr(arr)
}

func printArr(arr [][]int) {
	for _, i := range arr {
		for _, j := range i {
			fmt.Printf("%d  ", j)
		}
		fmt.Println()
	}
}

func printMatrixClockwisely(arr [][]int, rows, clomuns int) {
	if arr == nil || rows <= 0 || clomuns <= 0 {
		return
	}

	start := 0
	for rows > start*2 && clomuns > start*2 {
		printMatrixInCircle(arr, rows, clomuns, start)
		start += 1
	}

}

func printMatrixInCircle(arr [][]int, rows, clomns, start int) {
	endX := clomns - 1 - start
	endY := rows - 1 - start

	for i := start; i < endX; i++ {
		fmt.Printf("%d  ", arr[start][i])
	}

	if start < endX {
		for i := start + 1; i < endY; i++ {
			fmt.Printf("%d  ", arr[i][endX])
		}
	}

	if start < endY && start < endX {
		for i := endX - 1; i >= start; i-- {
			fmt.Printf("%d  ", arr[endY][i])
		}
	}

	if start < endX && start < endY-1 {
		for i := endY - 1; i >= start; i-- {
			fmt.Printf("%d  ", arr[i][endX])
		}
	}

}

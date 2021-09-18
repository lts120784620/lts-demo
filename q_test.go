package main

import (
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {
	nums := []int{5, 1, 3, 2}
	k := 3
	fmt.Println(topK(nums, k))
}

//		3
//	5   	2
//4    1
func topK(nums []int, k int) int {
	h := initHeap()
	for _, n := range nums {
		if len(h) < k {
			h = append(h, n)
		} else {
			if h[0] > n {
				h[0] = n
			}
		}

		adjustHeap(h)
	}
	return h[0]
}

func adjustHeap(nums []int) {
}

func initHeap() []int {
	return []int{}
}

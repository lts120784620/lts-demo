package array

import (
	"fmt"
	"testing"
)

/**
No.485 最大连续 1 的个数
思路：
    1、count变量记录每次遇到的1，记录++
	2、result统计当第一次遇到1或者数组结束时最大的result
*/
func TestName(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	var result int
	var count int
	for n := range nums {
		if n == 1 {
			count++
		} else {
			if count > result {
				result = count
			}
			count = 0
		}
		if count > result {
			result = count
		}
	}
	return result
}

package array

import (
	"fmt"
	"testing"
)

/**
No.53 最大子序和
描述：
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
思路：
    1、暴力迭代，从1～len(nums),每个迭代从数组开始遍历到结束，时间复杂读为O(n*m)
时间：
    1、2021/9/13
*/

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	for {

	}
}

package array

import (
	"fmt"
	"testing"
)

/**
No.1 两数之和
描述：
	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

思路：
    1、循环遍历两次数组，类似冒泡法
	2、存到map中，map[value]下标，遍历取第一个值，再直接获取m[target-value]是否存在
*/

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 22))
	fmt.Println(twoSum2(nums, 22))
}

func twoSum(nums []int, target int) []int {
	res := []int{}
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	res := []int{}
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return res
}

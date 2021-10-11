package DynamicProgramming

import (
	"fmt"
	"math"
	"testing"
)

/**
No.53 最大子序和
描述：
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
思路：
    1、动态规划：dp数组下标是数组索引，值是当前的最大和值。状态转移方程dp[i] = max(cur + dp[i-1])
时间：
    1、2021/10/8
*/

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-1}))
}

func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	// dp数组/初始化
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	// base case
	maxRes := dp[0]
	for i := 1; i < len(dp); i++ {
		// 极值
		dp[i] = int(math.Max(float64(nums[i]), float64(nums[i]+dp[i-1])))
		maxRes = int(math.Max(float64(dp[i]), float64(maxRes)))
	}
	return maxRes
}

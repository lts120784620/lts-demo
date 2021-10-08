package DynamicProgramming

import (
	"fmt"
	"math"
	"testing"
)

/**
No.300 最长递增子序列
描述：
	给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
	子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
思路：
    1、动态规划思路，注意状态转移方程，dp数组定义即可
	2、还有一种是二分法，较复杂，先不算了
时间：
    1、2021/10/8
*/
func TestLengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

func lengthOfLIS(nums []int) int {
	// 下标：nums索引
	// 值：以当前数字结尾的最长子序列长度
	dp := make([]int, len(nums))
	// 初始化
	for i := range dp {
		dp[i] = 1
	}

	maxRes := dp[0]
	for i := range dp {
		// 状态转移，向前找到小于当前值的，dp[i]的最大值
		for j := i - 1; j >= 0; j-- {
			// 找到前面小于当前的值
			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		maxRes = int(math.Max(float64(maxRes), float64(dp[i])))
	}

	return maxRes
}

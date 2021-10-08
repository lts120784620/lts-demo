package DynamicProgramming

import (
	"fmt"
	"math"
	"testing"
)

/**
No.322 零钱兑换
描述：
	给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
	计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
	你可以认为每种硬币的数量是无限的。
思路：
	典型的动态规划题目，找到最优解，有三种思路
    1、自上而下递归：
	2、自上而下递归，增加备忘录：
	3、迭代法：
时间：
    1、2021/10/8
*/

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange2([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange3([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := math.MaxInt
	for _, coin := range coins {
		subRes := coinChange(coins, amount-coin)
		if subRes == -1 {
			continue
		}
		res = int(math.Min(float64(res), float64(subRes+1)))
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

// 带备忘录，避免重复计算
func coinChange2(coins []int, amount int) int {
	var mem = make([]int, amount+1)
	for i := range mem {
		mem[i] = -999
	}
	_, res := dpCoin(coins, amount, mem)
	return res
}

func dpCoin(coins []int, amount int, mem []int) ([]int, int) {
	// base case
	if amount < 0 {
		return mem, -1
	}
	if amount == 0 {
		return mem, 0
	}
	// 先读备忘录，有则放弃整个子树遍历
	if n := mem[amount]; n != -999 {
		return mem, n
	}
	res := math.MaxInt
	for _, coin := range coins {
		_, subRes := dpCoin(coins, amount-coin, mem)
		if subRes == -1 {
			continue
		}
		res = int(math.Min(float64(res), float64(subRes+1)))
	}

	if res == math.MaxInt {
		res = -1
	}
	// 写入备忘录
	mem[amount] = res
	return mem, res
}

// 动态规划，数组迭代法
func coinChange3(coins []int, amount int) int {
	// dp下标为amount，值为子结果
	dp := make([]int, amount+1)
	// 细节一：初始化值一定要大于amount的最大值，很重要
	for i := range dp {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	// 外层取结果
	for subAmount := range dp {
		// 内层分子树
		for _, coin := range coins {
			// 终止判断
			if subAmount-coin < 0 {
				continue
			}
			dp[subAmount] = int(math.Min(float64(dp[subAmount]), float64(1+dp[subAmount-coin])))
		}
	}
	// 细节二
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

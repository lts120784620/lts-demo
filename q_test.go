package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	fmt.Println(maxStr("babad"))
}

func maxStr(s string) string {
	// dp数组
	dp := make([][]bool, len(s))
	// 初始化dp数组
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	left := 0
	// 定义距离left的最大长度
	maxRes := 1
	// 最大长度
	for i := 1; i <= len(s); i++ {
		// 状态转移方程
		for j := 0; j < len(s); j++ {
			right := j + i - 1
			if right > len(s)-1 {
				continue
			}
			// 判断左下标和右下标是否相等
			if s[j] != s[right] {
				dp[j][right] = false
			} else {
				//
				if j+1 < len(dp) && right-1 > 0 {
					dp[j][right] = dp[j+1][right-1]
				}
			}

			if !dp[j][right] {
				continue
			}
			// 最大长度判断
			if maxRes > right-j+1 {
				maxRes = right - j + 1
				left = j
			}

		}
	}

	return s[left : left+maxRes]
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	startIndex := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for l := 2; l <= len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			j := l + i - 1
			if j >= len(s) {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if l <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if maxLen < l && dp[i][j] {
				maxLen = l
				startIndex = i
			}
		}
	}
	return s[startIndex : startIndex+maxLen]
}

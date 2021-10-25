package main

import (
	"fmt"
	"math"
	"testing"
)

/**
No.1143. 最长公共子序列
描述：
	给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
	一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
	例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
	两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
思路：
    1、动态规划+暴力递归+备忘录，双指针在两个字符串遍历，判断在字符不相等情况下的操作次数，重点是
时间：
    1、2021/10/25
*/

func TestLongestCommonSubsequence(t *testing.T) {
	text1, text2 := "abcde", "acec"
	fmt.Println(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dpMom := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dpMom[i] = []int{}
		for j := 0; j < len(text2); j++ {
			dpMom[i] = append(dpMom[i], -1)
		}
	}
	dpReverse(text1, 0, text2, 0, dpMom)

	// 找最后一个，最大的[0][0]
	return dpMom[0][0]
}

// dp函数
func dpReverse(text1 string, i int, text2 string, j int, dpMon [][]int) int {
	// base case
	if i >= len(text1) || j >= len(text2) {
		return 0
	}
	if dpMon[i][j] != -1 {
		return dpMon[i][j]
	}
	if text1[i] == text2[j] {
		dpMon[i][j] = dpReverse(text1, i+1, text2, j+1, dpMon) + 1
	} else {
		// text1[i]不在text2[j]中，j向下寻找
		a := dpReverse(text1, i, text2, j+1, dpMon)
		// text1[j]不在text2[i]中，i向下寻找
		b := dpReverse(text1, i+1, text2, j, dpMon)
		dpMon[i][j] = int(math.Max(float64(a), float64(b)))
	}
	return dpMon[i][j]
}

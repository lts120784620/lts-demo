package main

import (
	"fmt"
	"testing"
)

/**
No.5. 最长回文子串
描述：
	给你一个字符串 s，找到 s 中最长的回文子串。
思路：
    1、中心扩散迭代法：两层遍历，一层遍历s的每个字符，二层用双指针法,遍历以s为中心的，从1-len(s)的所有
	2、动态规划：状态转移还是依据中心扩撒，i = i-1
时间：
    1、2021/10/22
*/

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("a"))
}

// 迭代中心扩散
func longestPalindrome(s string) string {
	res := ""
	// 一层遍历
	for i := range s {
		// 分奇偶两种情况
		// 奇数
		left, right := i, i
		subA := ""
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				if right-left >= len(subA) {
					subA = s[left : right+1]
				}
			} else {
				break
			}
			left--
			right++
		}

		// 偶数情况，一次遍历两个
		left, right = i-1, i
		subB := ""
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				if right-left >= len(subB) {
					subB = s[left : right+1]
				}

			} else {
				break
			}
			left--
			right++
		}

		if len(res) < len(subA) {
			res = subA
		}
		if len(res) < len(subB) {
			res = subB
		}

	}
	return res
}

// 动态规划
func longestPalindrome2(s string) string {
	// dp

	// base case
	//
	return ""
}

package main

import (
	"fmt"
	"testing"
)

/**
No.409. 最长回文串
描述：
	给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
	在构造过程中，请注意区分大小写。比如"Aa"不能当做一个回文字符串。
	注意:
	假设字符串的长度不会超过 1010。
思路：
    1、map保存每个字符出现的次数，找到偶数次+中心1即可
时间：
    1、2021/10/24
*/

func Test409LongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome409("abccccdd"))
}

func longestPalindrome409(s string) int {
	res := 0
	m := make(map[string]int, len(s))

	for _, char := range s {
		m[string(char)] = m[string(char)] + 1
		if cnt, ok := m[string(char)]; ok {
			// 成队出现，减去数量，增加结果
			n := cnt / 2
			if n > 0 {
				res += n * 2
				m[string(char)] = m[string(char)] - n*2
			}
		}
	}

	for _, v := range m {
		if v > 0 && v%2 > 0 {
			res++
			break
		}
	}

	return res
}

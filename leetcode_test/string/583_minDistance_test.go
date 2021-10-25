package main

import (
	"fmt"
	"testing"
)

/**
No.583. 两个字符串的删除操作
描述：
	给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
思路：
    1、动态规划+暴力递归+备忘录，
时间：
    1、2021/10/25
*/

func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("sea", "eat"))
}

func minDistance(word1 string, word2 string) int {

}

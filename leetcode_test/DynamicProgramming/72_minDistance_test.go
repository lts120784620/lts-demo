package DynamicProgramming

import (
	"fmt"
	"math"
	"testing"
)

/**
No.72 编辑距离
描述：
	给你两个单词word1 和word2，请你计算出将word1转换成word2 所使用的最少操作数。
	你可以对一个单词进行如下三种操作：
		插入一个字符
		删除一个字符
		替换一个字符
思路：
    1、递归法，暴力求解，用备忘录（二维数组）保存递归记录，避免重复递归
时间：
    1、2021/10/8
*/

func TestMinDistance(t *testing.T) {
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistance2(word1, word2))
}

// 递归解法 + 备忘录
func minDistance(word1 string, word2 string) int {
	var dpMom = make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		dpMom[i] = make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			dpMom[i][j] = -1
		}
	}
	return dpMin(dpMom, word1, word2, len(word1)-1, len(word2)-1)
}

// 定义递归函数，返回0-i，0-j的最小距离
func dpMin(dpMom [][]int, word1, word2 string, i, j int) int {
	if i >= 0 && j >= 0 && dpMom[i][j] > -1 {
		return dpMom[i][j]
	}
	// base case
	if i < 0 {
		// i结束，把j全部新增上
		return j + 1
	}
	if j < 0 {
		// j结束，把i全部新增上
		return i + 1
	}

	if word1[i] == word2[j] {
		// 当两值相等，同时向前移动
		dpMom[i][j] = dpMin(dpMom, word1, word2, i-1, j-1)
	} else {
		// 暴力递归出所有操作的组合，找出最小的移动方案
		res := dpMin(dpMom, word1, word2, i-1, j) + 1 // i删除
		a := dpMin(dpMom, word1, word2, i, j-1) + 1   // i插入
		b := dpMin(dpMom, word1, word2, i-1, j-1) + 1 // i,j 修改
		res = int(math.Min(float64(res), float64(a)))
		res = int(math.Min(float64(res), float64(b)))
		dpMom[i][j] = res
	}
	return dpMom[i][j]
}

// 动态规划，
func minDistance2(word1 string, word2 string) int {
	// dp数组初始化
	var dpMom = make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dpMom[i] = make([]int, len(word2)+1)
		for j := 0; j <= len(word2); j++ {
			dpMom[i][j] = 0
		}
	}

	// base case
	// 当其中一个完成了，另一个直接追加即可
	for i := 1; i <= len(word1); i++ {
		dpMom[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		dpMom[0][j] = j
	}

	// 状态转移方程
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dpMom[i][j] = dpMom[i-1][j-1]
			} else {
				min := dpMom[i-1][j] + 1 //
				a := dpMom[i][j-1] + 1   //
				b := dpMom[i-1][j-1] + 1 //
				// 最值
				min = int(math.Min(float64(min), float64(a)))
				min = int(math.Min(float64(min), float64(b)))
				dpMom[i][j] = min
			}
		}
	}

	return dpMom[len(word1)][len(word2)]
}

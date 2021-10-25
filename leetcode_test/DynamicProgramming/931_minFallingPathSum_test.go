package DynamicProgramming

import (
	"fmt"
	"math"
	"testing"
)

/**
No.931. 下降路径最小和
描述：
	给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。

思路：
    1、暴力递归，n层矩阵，每次遍历最小值即可, 但是要限制下一层的范围

时间：
    1、2021/10/24
*/

func TestMinFallingPathSum(t *testing.T) {
	//[[100,-42,-46,-41],[31,97,10,-10],[-58,-51,82,89],[51,81,69,-51]]
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(matrix))
}

// 暴力递归+备忘录
func minFallingPathSum(matrix [][]int) int {
	var dpMom = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dpMom[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dpMom[i][j] = math.MaxInt8
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		reverse931(matrix, 0, i, dpMom)
	}
	min := math.MaxInt8
	for i := 0; i < len(dpMom[0]); i++ {
		min = int(math.Min(float64(min), float64(dpMom[0][i])))
	}
	return min
}

func reverse931(matrix [][]int, col, i int, dpMom [][]int) int {
	if col >= len(matrix) {
		return 0
	}
	if v := dpMom[col][i]; v != math.MaxInt8 {
		return v
	}
	min := math.MaxInt8
	if i-1 >= 0 {
		min = int(math.Min(float64(min), float64(reverse931(matrix, col+1, i-1, dpMom))))
	}

	if i+1 <= len(matrix[col])-1 {
		min = int(math.Min(float64(min), float64(reverse931(matrix, col+1, i+1, dpMom))))
	}

	min = int(math.Min(float64(min), float64(reverse931(matrix, col+1, i, dpMom))))
	dpMom[col][i] = matrix[col][i] + min

	return dpMom[col][i]
}

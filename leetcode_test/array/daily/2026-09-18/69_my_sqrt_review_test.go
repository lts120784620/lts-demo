package daily20260918

import (
	"testing"
)

/*
LeetCode 69. x 的平方根（简单）
链接：https://leetcode.cn/problems/sqrtx/

题目要求：
给定非负整数 x，返回其算术平方根的整数部分，不能使用内置求幂能力。

复杂度目标：时间 O(log x)，额外空间 O(1)。
复习要求：不打开旧实现，在单调条件中寻找最后一个合法候选值。
建议限时：15 分钟。
提示触发点：独立思考 10 分钟后，仍然只能判断精确平方根，或不知道命中合法候选后
向哪个方向继续搜索，可索取一级提示。

完成评价（2026-09-19）：
  - 测试结果：通过（8 个测试用例）
  - 候选值满足的条件：mid*mid <= x
  - 命中合法候选后如何移动边界：保存 mid，并令 left = mid + 1 继续向右搜索
  - 如何避免 mid*mid 的边界问题：当前约束在 64 位 int 下可安全相乘；通用写法可比较 mid <= x/mid
  - 时间复杂度及原因：O(log x)，每轮将搜索区间缩小约一半
  - 空间复杂度：O(1)
  - 核心模式与不变量：在单调条件中寻找最后一个平方不超过 x 的整数
  - 提示使用：无
  - 主要错误类型：无
  - 本次掌握度（0～3）：2
  - 下次复习日期：2026-09-22
  - 备注：独立完成到期复习并通过全部测试；耗时未记录
*/
func mySqrtReview(x int) int {
	left := 1
	right := x
	max := 0
	if x <= 1 {
		return x
	}
	for left <= right {
		mid := (right-left)/2 + left
		if mid*mid <= x {
			if mid > max {
				max = mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return max
}

func TestMySqrtReview(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{name: "zero", x: 0, want: 0},
		{name: "one", x: 1, want: 1},
		{name: "two", x: 2, want: 1},
		{name: "perfect-square", x: 4, want: 2},
		{name: "truncate", x: 8, want: 2},
		{name: "before-square", x: 15, want: 3},
		{name: "large-boundary", x: 2147395599, want: 46339},
		{name: "max-int32", x: 2147483647, want: 46340},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrtReview(tt.x); got != tt.want {
				t.Fatalf("mySqrtReview(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}

package daily20260914

import "testing"

/*
LeetCode 69. x 的平方根（简单）
链接：https://leetcode.cn/problems/sqrtx/

题目要求：
给定一个非负整数 x，返回它的算术平方根的整数部分，小数部分直接舍去。
不能使用 pow、指数运算符等内置求幂能力。

示例：

	x = 4，返回 2
	x = 8，返回 2，因为 sqrt(8) 约为 2.828，小数部分舍去

约束：

	0 <= x <= 2^31 - 1

复杂度目标：时间 O(log x)，额外空间 O(1)。
今日训练点：在单调条件上做二分，并保存最后一个满足条件的候选答案。
建议限时：25 分钟。
提示触发点：独立思考 10 分钟后，仍然只能判断某个 mid 是否为精确平方根，
却无法处理 x = 8 这类答案不存在于等式中的情况时，可索取一级提示。

完成评价（2026-09-15）：
  - 测试结果：本地完整测试通过
  - 候选答案需要满足的条件：mid <= x/mid
  - 命中候选条件后向哪个方向继续搜索：记录 mid，并继续向右寻找更大的合法值
  - 如何避免乘法比较的边界问题：使用 mid <= x/mid，避免计算 mid*mid
  - 时间复杂度：O(log x)
  - 空间复杂度：O(1)
  - 核心模式与不变量：在单调条件中寻找最后一个平方不超过 x 的整数
  - 提示使用：二级
  - 主要错误类型：复杂度
  - 本次掌握度（0～3）：1
  - 下次复习日期：2026-09-17
  - 备注：能够识别从 O(x) 线性枚举优化为二分查找
*/
func mySqrtLinear(x int) int {
	if x <= 1 {
		return x
	}
	for i := 0; i <= x; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return 0
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left := 1
	right := x
	answer := 0
	for left <= right {
		mid := (right-left)/2 + left
		if mid <= x/mid {
			answer = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return answer
}

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{name: "zero", x: 0, want: 0},
		{name: "one", x: 1, want: 1},
		{name: "2", x: 2, want: 1},
		{name: "perfect-square", x: 4, want: 2},
		{name: "truncate", x: 8, want: 2},
		{name: "before-square", x: 15, want: 3},
		{name: "another-square", x: 16, want: 4},
		{name: "large-boundary", x: 2147395599, want: 46339},
		{name: "max-int32", x: 2147483647, want: 46340},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Fatalf("mySqrt(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}

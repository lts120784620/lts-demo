package math

import (
	"fmt"
	"testing"
)

/**
No.509 斐波那契数
描述：
	斐波那契数，通常用F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
	F(0) = 0，F(1)= 1
	F(n) = F(n - 1) + F(n - 2)，其中 n > 1
	给你 n ，请计算 F(n)
思路：
    1、递归，简单
时间：
    1、2021/10/7
*/

func TestFib(t *testing.T) {
	fmt.Println(fib(4))
}

func fib(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

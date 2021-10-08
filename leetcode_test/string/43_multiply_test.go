package main

import (
	"fmt"
	"testing"
)

/**
No.43 字符串相乘
描述：
	给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
思路：
    1、模拟乘法计算，转换为A*B1+A*B2+...,转换为加法操作，A*Bn模仿乘法进位即可
时间：
    1、2021/9/27
*/

func TestMultiply(t *testing.T) {
	fmt.Println(multiply("401716832807512840963", "167141802233061013023557397451289113296441069"))
}

func multiply(num1 string, num2 string) string {
	// 遍历num2,拆成n*1000+n*100+1的形式
	x := ""
	sum1 := ""
	for i := len(num2) - 1; i >= 0; i-- {
		b := int(num2[i] - '0')
		// 求num1 * b
		y := x
		for j := len(num1) - 1; j >= 0; j-- {
			a := int(num1[j] - '0')
			// 注意：补零不能直接乘，要字符串追加，否则会越界
			s := fmt.Sprintf("%d", int64(a*b))
			if int(a*b) > 0 {
				s += y
			}
			fmt.Println(a, "*", b, "*", y, "*", x)
			sum1 = add(s, sum1)
			// 移位补零
			y += "0"
		}
		// 移位补零
		x += "0"
	}

	return sum1
}

func add(num1, num2 string) string {
	res := ""
	p1, p2 := len(num1)-1, len(num2)-1
	// 表示进位
	k := 0
	for p1 >= 0 && p2 >= 0 {
		a := int(num1[p1] - '0')
		b := int(num2[p2] - '0')
		sum := a + b + k
		res = fmt.Sprintf("%d%s", sum%10, res)
		k = sum / 10
		p1--
		p2--
	}

	for p1 >= 0 {
		a := int(num1[p1] - '0')
		sum := a + k
		res = fmt.Sprintf("%d%s", sum%10, res)
		k = sum / 10
		p1--
	}

	for p2 >= 0 {
		b := int(num2[p2] - '0')
		sum := b + k
		res = fmt.Sprintf("%d%s", sum%10, res)
		k = sum / 10
		p2--
	}

	if k > 0 {
		res = fmt.Sprintf("%d%s", k, res)
	}

	return res
}

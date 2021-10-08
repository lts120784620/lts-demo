package main

import (
	"fmt"
	"testing"
)

/**
No.415 字符串相加
描述：
	给定两个字符串形式的非负整数num1 和num2，计算它们的和并同样以字符串形式返回。你不能使用任何內建的用于处理大整数的库（比如 BigInteger），也不能直接将输入的字符串转换为整数形式。
思路：
    1、模拟按位加，结果进位到下一位
时间：
    1、2021/9/26
*/

func TestDddStrings(t *testing.T) {
	fmt.Println(addStrings("999711", "292"))
}

func addStrings(num1 string, num2 string) string {
	res := ""
	p1, p2 := len(num1)-1, len(num2)-1
	k := 0
	for p1 >= 0 && p2 >= 0 {
		sum := int(num1[p1]-'0') + int(num2[p2]-'0') + k
		res = fmt.Sprintf("%d%s", sum%10, res)
		k = sum / 10
		p1--
		p2--
	}
	// 注意p1超长处理
	for p1 >= 0 {
		sum := int(num1[p1]-'0') + k
		k = sum / 10
		res = fmt.Sprintf("%d%s", sum%10, res)
		p1--
	}
	// 注意p2超长处理
	for p2 >= 0 {
		sum := int(num2[p2]-'0') + k
		k = sum / 10
		res = fmt.Sprintf("%d%s", sum%10, res)
		p2--
	}
	// 注意最终
	if k != 0 {
		res = fmt.Sprintf("%d%s", k, res)
	}
	return res
}

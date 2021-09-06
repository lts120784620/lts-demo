package int

import (
	"fmt"
	"testing"
)

/**
No.9 回文数
描述：
	给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

思路：
    1、将数字转换成字符串，判断
	2、或者将数字的后半段取出来与前半段比较，取后半段数字的方式可以通过求余的方法
*/

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(12213))
	fmt.Println(isPalindrome(21120))
	fmt.Println(isPalindrome(1))
}

func isPalindrome(x int) bool {
	// 如果最后一位是0，不可能是回文
	if x < 0 || (x != 0 && x%10 == 0) {
		// 负数一定不是回文
		return false
	}
	// 1221 0
	// 122 1
	// 12 21
	after := 0
	for x > after {
		after = after*10 + x%10
		x = x / 10
		fmt.Println(x, " = ", after)
	}
	// 注意奇数判断
	if x < after {
		return x == after/10
	}
	return x == after
}

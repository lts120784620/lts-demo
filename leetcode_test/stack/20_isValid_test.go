package stack

import (
	"fmt"
	"testing"
)

/**
No.20 有效的括号
思路：
	1、使用栈结构达到先入后出匹配的目的，用map确定括号相对的规则，
关键点是左放到栈中，遍历到右时，弹栈匹配

时间：
	2021-09-11
*/

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	m := map[string]string{"]": "[", "}": "{", ")": "("}
	stack := NewStack()

	for _, v := range s {
		if _, ok := m[string(v)]; !ok {
			stack.Push(string(v))
		} else {
			left := stack.Pop()
			// "]"
			if left == nil {
				return false
			}
			if m[string(v)] != left.(string) {
				return false
			}
		}
	}

	// "[)"
	if !stack.IsEmpty() {
		return false
	}
	return true
}

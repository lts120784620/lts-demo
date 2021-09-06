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
*/

func TestIsValid(t *testing.T) {
	str := "()([]{}"
	s := NewStack()
	m := map[string]string{"]": "[", "}": "{", ")": "("}
	for _, v := range str {
		// 如果不在map中，说明是左括号，放到栈中
		fmt.Println(string(v))
		if _, ok := m[string(v)]; !ok {
			s.Push(string(v))
		} else if s.IsEmpty() || s.Pop().(string) != m[string(v)] {
			fmt.Println("错误，", string(v), "不符合校验规则")
			return
		}
	}
	if s.IsEmpty() {
		fmt.Println("正确，符合校验规则")
	} else {
		fmt.Println("错误，不符合校验规则")
	}
}

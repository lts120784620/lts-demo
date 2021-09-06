package main

import (
	"fmt"
	"lts-demo/leetcode_test/common"
)

func main() {
	fmt.Println(isValid("()[]{}{"))
}

func isValid(s string) bool {
	vstack := common.NewStack()
	left := 0
	for _, c := range s {
		switch string(c) {
		case "(":
			fallthrough
		case "[":
			fallthrough
		case "{":
			vstack.Push(string(c))
			left++
		case ")":
			j := vstack.Pop()
			if j != "(" {
				return false
			}
			left--
		case "]":
			j := vstack.Pop()
			if j != "[" {
				return false
			}
			left--
		case "}":
			j := vstack.Pop()
			if j != "{" {
				return false
			}
			left--
		}
	}
	return left == 0
}

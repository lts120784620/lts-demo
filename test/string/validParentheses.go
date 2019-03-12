package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()[]{}{"))
}
type (
	stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)

func NewStack() *stack {
	return &stack{nil,0}
}

func (this *stack) pop() interface{} {
	if this == nil || this.length == 0{
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *stack) push(value interface{}) {
	top := &node{value:value,prev:this.top}
	this.top = top
	this.length++
}

func isValid(s string) bool {
	vstack := NewStack()
	left := 0
	for _,c := range s{
		switch string(c) {
		case "(":
			fallthrough
		case "[":
			fallthrough
		case "{":
			vstack.push(string(c))
			left++
		case ")":
			j := vstack.pop()
			if j != "("{
				return false
			}
			left--
		case "]":
			j := vstack.pop()
			if j != "["{
				return false
			}
			left--
		case "}":
			j := vstack.pop()
			if j != "{"{
				return false
			}
			left--
		}
	}
	return left == 0
}
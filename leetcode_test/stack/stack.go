package stack

/**
思路:用链表实现的栈结构
*/

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Peek() interface{} {
	if this == nil || this.length == 0 {
		return nil
	}
	n := this.top
	return n.value
}

func (this *Stack) Pop() interface{} {
	if this == nil || this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value interface{}) {
	top := &node{value: value, prev: this.top}
	this.top = top
	this.length++
}

func (this *Stack) IsEmpty() bool {
	if this == nil {
		return false
	}
	return this.length == 0
}

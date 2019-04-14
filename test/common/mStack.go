package common

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

func (this *stack) Top() interface{} {
	if this == nil || this.length == 0{
		return nil
	}
	n := this.top
	return n.value
}

func (this *stack) Pop() interface{} {
	if this == nil || this.length == 0{
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *stack) Push(value interface{}) {
	top := &node{value:value,prev:this.top}
	this.top = top
	this.length++
}

func (this *stack) IsEmpty() bool{
	if this == nil{
		return false
	}
	return this.length == 0
}
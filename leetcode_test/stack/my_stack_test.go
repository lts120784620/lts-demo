package stack

import (
	"fmt"
	"lts-demo/leetcode_test/queue"
	"testing"
)

/**
No.225 用队列实现栈
思路：
    1、
*/

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type MyStack struct {
	input  *queue.Queue
	output *queue.Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue.New(),
		queue.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.input.Enqueue(x)
	this.output.Enqueue(this.input.Dequeue())
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.output.Dequeue().(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.output.Peek().(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.output.Len() == 0
}

func TestMyStack(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Empty())
	s.Top()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Empty())
}

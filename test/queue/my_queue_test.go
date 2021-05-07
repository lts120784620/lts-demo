package queue

import (
	"fmt"
	"lts-demo/test/stack"
	"testing"
)

/**
No.232 用栈实现队列
思路：
    1、
*/

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyQueue struct {
	input  *stack.Stack
	output *stack.Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack.NewStack(),
		stack.NewStack(),
	}

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input.Push(x)
	t := this.input.Pop()
	this.output.Push(t)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.output.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.output.Peek().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.output.IsEmpty()
}

func TestQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Empty())
	q.Peek()
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}
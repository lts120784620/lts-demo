package deque

/**
思路：用链表实现双端队列
*/

type (
	Deque struct {
		start, end *node
		length     int
	}
	node struct {
		value interface{}
		next  *node
	}
)

// Create a new queue
func New() *Deque {
	return &Deque{nil, nil, 0}
}

// Take the next item off the front of the queue
func (this *Deque) Dequeue() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return n.value
}

// Put an item on the end of a queue
func (this *Deque) Enqueue(value interface{}) {
	n := &node{value, nil}
	if this.length == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.next = n
		this.end = n
	}
	this.length++
}

// Return the number of items in the queue
func (this *Deque) Len() int {
	return this.length
}

// Return the first item in the queue without removing it
func (this *Deque) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.start.value
}

func (this *Deque) IsEmpty() bool {
	return this.length == 0
}

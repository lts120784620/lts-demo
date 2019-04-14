package common

type MyQueue struct {
	length int
	items  []interface{}
}

func NewQueue() *MyQueue {
	return &MyQueue{0, nil}
}

func (this *MyQueue) Add(v interface{}) {
	if(this == nil){
		return
	}
	this.items = append(this.items, v)
	this.length++
}

func (this *MyQueue) Offer() interface{} {
	if(this == nil || this.length == 0){
		return nil
	}
	res := this.items[0]
	this.items = this.items[1:this.length]
	this.length--
	return res
}

func (this *MyQueue) IsEmpty() bool {
	if(this == nil){
		return true
	}
	return this.length == 0
}
func (this *MyQueue) Size() int {
	if(this == nil){
		return 0
	}
	return this.length
}


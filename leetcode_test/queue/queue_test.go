package queue

import (
	"fmt"
	"testing"
)

/**
No.
描述：
	二叉树测试
思路：
    1、
时间：
    1、2021/9/28
*/

func TestBinaryTree(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.DequeueFront())
	fmt.Println(q.DequeueFront())
	fmt.Println(q.DequeueFront())
}

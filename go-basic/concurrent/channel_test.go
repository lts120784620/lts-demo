package concurrent

import (
	"fmt"
	"testing"
)

/**
No.
描述：
	channel测试
思路：
    1、
时间：
    1、2021/9/21
*/

func TestChannel(t *testing.T) {
	c := make(chan string, 3)

	c <- "a"
	c <- "b"
	c <- "c"
	c <- "d"

	go consumerA(c)
	go consumerB(c)
}

func consumerA(a <-chan string) {
	val := <-a
	fmt.Println("goRoutineA received the data", val)
}

func consumerB(b <-chan string) {
	val := <-b
	fmt.Println("goRoutineB received the data", val)
}

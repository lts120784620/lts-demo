package concurrent

import (
	"fmt"
	"testing"
	"time"
)

/**
No. 两个协程交替打印奇偶数
思路：
    1、
*/

var oddChannel = make(chan int, 1)
var evenChannel = make(chan int, 1)
var num int

func TestName(t *testing.T) {

	go odd()
	go even()

	for {
		select {
		case odd := <-oddChannel:
			fmt.Println("odd is ", odd)
		case even := <-evenChannel:
			fmt.Println("even is ", even)
		}
	}
}

func odd() {

	for i := 0; i < 101; i++ {
		if i%2 != 0 {
			oddChannel <- i
			time.Sleep(1 * time.Second)
		}
	}
}

func even() {
	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			evenChannel <- i
			time.Sleep(1 * time.Second)
		}
	}
}

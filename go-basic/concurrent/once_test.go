package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

var c = make(chan int, 3)

func TestOnce(t *testing.T) {
	var one sync.Once

	for {
		go func() {
			one.Do(closeChan)
		}()
	}
}

func closeChan() {
	close(c)
	fmt.Println("close channel")
}

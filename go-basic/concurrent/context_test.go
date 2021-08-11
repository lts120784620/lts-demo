package concurrent

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())

	go callMeFather(parent)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(4 * time.Second)
	fmt.Println("over")
}

func callMeFather(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("hahahaha ,just kidding")
			// 这里的退出很重要，否则会卡在这里一直接收退出信号
			return
		default:
			fmt.Println("im your father")
		}
	}

}

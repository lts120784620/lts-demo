package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"time"
)

func allocate() {
	_ = make([]byte, 1<<20)
}

func main() {
	// 1.通过执行命令 GODEBUG=gctrace=1 ./gc_main
	// 2.通过读取trace.out文件 go tool trace ../../trace.out
	f, _ := os.Create("./trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	for n := 1; n < 200000; n++ {
		allocate()
		//time.Sleep(time.Second)
	}

	//go printGCStats()
	//go printMemStats()
	//select {}
}

// 3.定时输出gc状态
func printGCStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}

// 4.通过runtime对内存状态监控
func printMemStats() {
	t := time.NewTicker(time.Second)
	s := runtime.MemStats{}

	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(&s)
			fmt.Printf("gc %d last@%v, next_heap_size@%vMB\n", s.NumGC, time.Unix(int64(time.Duration(s.LastGC).Seconds()), 0), s.NextGC/(1<<20))
		}
	}

}

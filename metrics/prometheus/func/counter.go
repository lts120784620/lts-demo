package _func

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"time"
)

/**
No.
描述：

思路：
    1、
时间：
    1、2021/9/25
*/

var Net = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace:   "host",
		Subsystem:   "container",
		Name:        "net_if_in",
		Help:        "这是一个测试用的net指标",
		ConstLabels: nil,
	})

func init() {
	go func() {
		for {
			netStat()
			time.Sleep(10 * time.Second)
		}
	}()
}

func netStat() {
	Net.Add(float64(rand.Intn(99)))
}

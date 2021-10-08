package _func

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"time"
)

/**
No.
描述：
	gauge 类型
思路：
    1、
时间：
    1、2021/9/25
*/

var Cpu = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace:   "host",
	Subsystem:   "container",
	Name:        "cpu_busy",
	Help:        "这个一个测试用的cpu指标",
	ConstLabels: nil,
})

func init() {
	go func() {
		for {
			cpuStat()
			time.Sleep(10 * time.Second)
		}
	}()
}

func cpuStat() {
	Cpu.Set(float64(rand.Intn(99)))
}

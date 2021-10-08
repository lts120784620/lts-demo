package _func

import "github.com/prometheus/client_golang/prometheus"

/**
No.
描述：

思路：
    1、
时间：
    1、2021/9/25
*/

func Init() {
	prometheus.MustRegister(Cpu)
	prometheus.MustRegister(Net)
}

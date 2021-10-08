package prometheus

import (
	_func "lts-demo/metrics/prometheus/func"
	"lts-demo/metrics/prometheus/http"
	"testing"
)

/**
No.
描述：

思路：
    1、
时间：
    1、2021/9/24
*/

func TestPromTest(t *testing.T) {
	_func.Init()
	http.Start()
}

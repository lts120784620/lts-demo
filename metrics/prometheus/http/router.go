package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/**
No.
描述：
	http路由
思路：
    1、
时间：
    1、2021/9/24
*/

func route(r *gin.Engine) {

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

}

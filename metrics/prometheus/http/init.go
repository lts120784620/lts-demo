package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

/**
No.
描述：

思路：
    1、
时间：
    1、2021/9/24
*/

func Start() {
	r := gin.Default()

	route(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("http server start error!")
		return
	}
}

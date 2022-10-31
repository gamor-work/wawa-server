package api

import (
	"github.com/gin-gonic/gin"
)

var request *gin.Engine

func Register(app *gin.Engine) {
	request = app
}

/*
 统一封装 Get 方法用于数据捕获和记录日志
*/
func WawaGet(url string, handler func(c *gin.Context)) {
	request.GET(url, func(req *gin.Context) {
		handler(req)
	})
}

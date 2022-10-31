package router

import (
	request "gin/api"
	Controller "gin/controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	// 给全局请求增加 gin.Engine 对象
	request.Register(app)

	request.WawaGet("/language", Controller.UserHandler)
	request.WawaGet("/redis", Controller.RedisHandler)
}

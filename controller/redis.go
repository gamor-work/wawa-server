package controller

import (
	"context"
	"fmt"
	"gin/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedisHandler(c *gin.Context) {
	ctx := context.Background()
	err := global.RedisDB.Set(ctx, "name", "wawa", 0).Err()

	if err != nil {
		//panic(err)
	}

	val, err := global.RedisDB.Get(ctx, "name").Result()
	if err != nil {
		//panic(err)
	}
	fmt.Println("key", val)

	c.JSON(http.StatusOK, val)
}

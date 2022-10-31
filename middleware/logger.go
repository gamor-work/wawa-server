package middleware

import (
	"fmt"
	global "gin/global"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()
		//data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Printf(`「🎉 %s」: 接口 %s 于 %s 发送请求 请求总时间：%s，请求结果: %s
`, global.SYSTEM_NAME, c.FullPath(), t.Format("2006-01-02 15:04:05"), latency, strconv.Itoa(status))
	}
}

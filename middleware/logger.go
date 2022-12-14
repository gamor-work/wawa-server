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
		fmt.Printf(`ãð %sã: æ¥å£ %s äº %s åéè¯·æ± è¯·æ±æ»æ¶é´ï¼%sï¼è¯·æ±ç»æ: %s
`, global.SYSTEM_NAME, c.FullPath(), t.Format("2006-01-02 15:04:05"), latency, strconv.Itoa(status))
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/jeanphorn/log4go"
	"time"
)

// 日志中间件
type LogMiddleware struct {
}
// 拦截所有的请求，输出到日志
func (*LogMiddleware) Logger() gin.HandlerFunc {

	return func (c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		log.Info(" %3d | %3v  | %3s | %s | %s",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

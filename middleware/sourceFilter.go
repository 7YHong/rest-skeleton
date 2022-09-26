package middleware

import (
	"github.com/gin-gonic/gin"
)

func SourceFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		ip := c.ClientIP()
		if len(host) < 11 || len(ip) < 9 {
			c.AbortWithStatus(400)
		} else if host[len(host)-11:] == "cpic.com.cn" || ip[:6] == "10.84." {
			// 通过匹配，什么都不做
		} else {
			c.AbortWithStatus(403)
		}

	}
}

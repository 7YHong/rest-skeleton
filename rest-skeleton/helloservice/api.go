package helloservice

import (
	"github.com/gin-gonic/gin"
	coremiddleware "rest-skeleton/core/middleware"
	"rest-skeleton/helloservice/middleware"
)

func SetRouter(r *gin.Engine) {
	r.GET("/hello",
		coremiddleware.CorsMiddleware(),
		middleware.ReqLogMiddleware,
		helloHanlder(),
	)
}

func helloHanlder() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(200, "hello world!")
	}
}

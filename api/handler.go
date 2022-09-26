package api

import (
	"github.com/gin-gonic/gin"
	"io"
)

func setRouter() {
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(200)
	})

	//router.Any("echo", GinCtx(func(c *ginCtx) (interface{}, error) {
	//	body, _ := io.ReadAll(c.Request.Body)
	//	return gin.H{
	//		"method": c.Request.Method,
	//		"header": c.Request.Header,
	//		"body":   string(body)}, nil
	//}))

	router.Any("echo", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		c.JSON(200, gin.H{
			"method": c.Request.Method,
			"header": c.Request.Header,
			"body":   string(body)})
	})

}

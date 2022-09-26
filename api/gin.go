package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"rest-skeleton/e"
	"rest-skeleton/middleware"
	"rest-skeleton/utils"
	"runtime"
)

var router *gin.Engine

type (
	ginCtx struct{ *gin.Context }
)

func init() {
	router = gin.New()
	router.Use(gin.CustomRecovery(middleware.Recovery))
	router.Use(middleware.GinLogger(utils.Logger))
}

func Startup(cCtx *cli.Context) error {
	//appkey := cCtx.String("appkey")
	port := cCtx.Uint("port")

	setRouter()

	welcome(port)
	return router.Run(fmt.Sprintf(":%d", port))
}

func GinCtx(fn func(c *ginCtx) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if data, err := fn(&ginCtx{ctx}); err != nil {
			_ = ctx.Error(err)
			ee := e.Error{}
			if errors.As(err, &ee) {
				ctx.JSON(200, gin.H{
					"status": ee.Code(),
					"errmsg": err.Error(),
				})
			} else {
				ctx.JSON(200, gin.H{
					"status": 500,
					"errmsg": err.Error(),
				})
			}
		} else {
			ctx.JSON(200, gin.H{
				"status": 0,
				"data":   data,
			})
		}
	}
}

func welcome(port uint) {
	utils.Logger.Sugar().Infof("ginServer   Name:      %s", "flyPenguin")
	utils.Logger.Sugar().Infof("Listen      Port:      %d", port)
	utils.Logger.Sugar().Infof("System      Name:      %s", runtime.GOOS)
	utils.Logger.Sugar().Infof("Go          Version:   %s", runtime.Version()[2:])
}

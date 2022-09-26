package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mix-go/dotenv"
	"net/http"
	"os"
	"os/signal"
	"rest-skeleton/core/di"
	_ "rest-skeleton/core/dotenv"
	"rest-skeleton/core/middleware"
	"syscall"
	"time"
)

var server *ginServer

type ginServer struct {
	server *http.Server
	Router *gin.Engine
}

func Server() *ginServer {
	return server
}

func (s *ginServer) Startup() {
	logger := di.Zap()
	// signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		logger.Info("ginServer shutdown")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err := s.server.Shutdown(ctx); err != nil {
			logger.Errorf("ginServer shutdown error: %s", err)
		}
	}()

	// run
	welcome(s.server.Addr)
	logger.Infof("ginServer start at %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
}

func init() {
	addr := dotenv.Getenv("GIN_ADDR").String(":8080")
	mode := dotenv.Getenv("GIN_MODE").String(gin.ReleaseMode)

	// server
	gin.SetMode(mode)
	router := gin.New()
	router.Use(gin.CustomRecovery(middleware.Recovery))

	s := &http.Server{}
	s.Addr = addr
	s.Handler = router

	server = &ginServer{s, router}
}

package initmodule

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/heshiyingw/gin-ext/extend"
	"github.com/heshiyingw/gin-ext/middleware"
	"gitlab.myshuju.top/heshiying/gin-skeleton/config"
	"gitlab.myshuju.top/heshiying/gin-skeleton/g"
	internalHttp "gitlab.myshuju.top/heshiying/gin-skeleton/internal/http"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

var httpServer *http.Server

// Gin 初始化gin
func Gin() *gin.Engine {
	engine := gin.New()
	extend.RegisterTranslations(engine)
	logger, _ := zap.NewProductionConfig().Build()

	engine.Use(middleware.GinLogger(logger), middleware.GinRecovery(logger, true))
	internalHttp.SetRouter(engine)

	return engine
}

// StartHttpServer httpServer启动
func StartHttpServer(r http.Handler) error {
	httpServer = &http.Server{
		Addr:              config.GetConfig().HttpServerConfig.GetAddr(),
		Handler:           r,
		ReadTimeout:       config.GetConfig().HttpServerConfig.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      config.GetConfig().HttpServerConfig.WriteTimeout,
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			g.Error("httpServer shutdown err", zap.Error(err))
			panic(err)
		}
	}()
	<-ctx.Done()
	stop()

	// 给30秒处理未完成的任务
	return endHttpServer()
}
func endHttpServer() error {
	if httpServer == nil {
		return nil
	}
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()
	err := httpServer.Shutdown(timeoutCtx)
	if err != nil {
		g.Error("httpServer shutdown err", zap.Error(err))
		return err
	}
	g.Info("http Server had shutdown")
	return nil
}

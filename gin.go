package gin_skeleton

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heshiyingw/gin-ext/middleware"
	"gitlab.myshuju.top/heshiying/gin-skeleton/config"
	"gitlab.myshuju.top/heshiying/gin-skeleton/g"
	extend2 "gitlab.myshuju.top/heshiying/gin-skeleton/pkg/ginext/resp"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

type RegisterRouterFunc func(r *gin.Engine) error
type ServerStartFunc func() error

type tempManger struct {
	midMap            map[uintptr]struct{}
	hadCustomerRouter bool
}

var (
	httpServer  *http.Server
	engin       *gin.Engine
	serverFuncs []ServerStartFunc
	tm          = &tempManger{
		midMap:            make(map[uintptr]struct{}, 50),
		hadCustomerRouter: false,
	}
)

func init() {
	engin = createGin()
}

func UseMiddleware(mid ...gin.HandlerFunc) {
	if tm == nil {
		g.Panic("请在server启动前添加中间件")
		panic("")
	}
	for _, handlerFunc := range mid {
		fvp := reflect.ValueOf(handlerFunc).Pointer()
		if _, ok := tm.midMap[fvp]; ok {
			g.Panic("中间件重复添加", zap.String("func", fmt.Sprintf("%p", handlerFunc)))
			panic("中间件重复添加")
			tm.midMap[fvp] = struct{}{}
		}

	}
	engin.Use(mid...)

}
func RegisterRouter(f RegisterRouterFunc) {
	err := f(engin)
	if err != nil {
		g.Panic("RegisterRouterErr", zap.Error(err))
		return
	}
}
func registerDefaultRouter(r *gin.Engine) error {
	r.GET("/ping", func(ctx *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
		ctx.String(http.StatusOK, "pong")
	})
	r.GET("/ok", func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "ok")
	})
	return nil
}

// createGin 初始化gin
func createGin() *gin.Engine {
	engine := gin.New()
	extend2.RegisterTranslations(engine)
	logger, _ := zap.NewProductionConfig().Build()

	engine.Use(middleware.GinLogger(logger), middleware.GinRecovery(logger, true))

	return engine
}

// StartServer Server启动
func StartServer(c *config.HttpConfig) error {
	if !tm.hadCustomerRouter {
		RegisterRouter(registerDefaultRouter)
	}
	if c == nil {
		c = &config.DefaultHttpConfig
	}

	httpServer = &http.Server{
		Addr:              c.GetAddr(),
		Handler:           engin,
		ReadTimeout:       time.Duration(c.ReadTimeoutSec) * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Duration(c.WriteTimeoutSec) * time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	tm.midMap = nil
	go func() {
		g.Info("start server,", zap.String("listen", config.GetConfig().HttpServerConfig.GetAddr()))
		err := httpServer.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				g.Info("Server Closed")
			} else {
				g.Error("httpServer shutdown err", zap.Error(err))
				panic(err)
			}
		}
	}()
	for _, serverFunc := range serverFuncs {
		go func() {
			if err := serverFunc(); err != nil {
				g.Error("serverFunc star err", zap.Error(err))
				panic(err)
			}
		}()
	}
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

func RegisterServer(f ServerStartFunc) {
	serverFuncs = append(serverFuncs, f)
}

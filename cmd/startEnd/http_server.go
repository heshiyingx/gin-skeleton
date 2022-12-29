package startEnd

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"skeleton/config"
	"time"
)

var httpServer http.Server

// StartHttpServer httpServer启动
func StartHttpServer(r *gin.Engine) error {
	httpServer = http.Server{
		Addr:              config.GetConfig().HttpServerConfig.GetAddr(),
		Handler:           r,
		ReadTimeout:       config.GetConfig().HttpServerConfig.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      config.GetConfig().HttpServerConfig.WriteTimeout,
	}
	return httpServer.ListenAndServe()
}
func StartGoroutineHttpServer(r *gin.Engine) {
	go func() {
		err := StartHttpServer(r)
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}
func EndHttpServer() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()
	err := httpServer.Shutdown(ctx)
	if err != nil {
		return
	}
	log.Println("httpserver end")
}

package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetRouter(r *gin.Engine) {
	r.GET("/ping", func(ctx *gin.Context) {
		time.Sleep(time.Second * 20)
		ctx.String(http.StatusOK, "pong")
	})
	r.GET("/ok", func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "ok")
	})
}

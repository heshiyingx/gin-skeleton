package initmodule

import (
	"github.com/gin-gonic/gin"
	"github.com/heshiyingw/gin-ext/extend"
	"github.com/heshiyingw/gin-ext/middleware"
	"go.uber.org/zap"
	internalHttp "skeleton/internal/http"
)

// Gin 初始化gin
func Gin() *gin.Engine {
	engine := gin.New()
	extend.RegisterTranslations(engine)
	logger, _ := zap.NewProductionConfig().Build()

	engine.Use(middleware.GinLogger(logger), middleware.GinRecovery(logger, true))
	internalHttp.SetRouter(engine)

	return engine
}

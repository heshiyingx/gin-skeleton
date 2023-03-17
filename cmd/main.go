package main

import (
	gin_skeleton "gitlab.myshuju.top/heshiying/gin-skeleton"
	"gitlab.myshuju.top/heshiying/gin-skeleton/config"
	"gitlab.myshuju.top/heshiying/gin-skeleton/initmodule"
	"gitlab.myshuju.top/heshiying/gin-skeleton/pkg/ginext/middleware"
)

func main() {
	cfg := config.GetConfig()
	initmodule.ConfigToModel("./config.yaml", cfg)
	gin_skeleton.UseMiddleware(middleware.UserPromGatewayMiddleware("mall-order-svc", "http://192.168.31.16:9091/"))
	gin_skeleton.StartServer(&cfg.HttpServerConfig)

	//gin_skeleton.StartHttpServer(nil)
}

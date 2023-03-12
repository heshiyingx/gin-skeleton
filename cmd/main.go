package main

import (
	gin_skeleton "gitlab.myshuju.top/heshiying/gin-skeleton"
	"gitlab.myshuju.top/heshiying/gin-skeleton/config"
	"gitlab.myshuju.top/heshiying/gin-skeleton/initmodule"
)

func main() {
	cfg := config.GetConfig()
	initmodule.ConfigToModel("./config.yaml", cfg)
	gin_skeleton.StartHttpServer(&cfg.HttpServerConfig)
	//gin_skeleton.StartHttpServer(nil)
}

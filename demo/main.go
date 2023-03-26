package main

import (
	gin_skeleton "gitlab.myshuju.top/heshiying/gin-skeleton"
	"gitlab.myshuju.top/heshiying/gin-skeleton/demo/config"
	"gitlab.myshuju.top/heshiying/gin-skeleton/initmodule"
	"gitlab.myshuju.top/heshiying/gin-skeleton/pkg/db/gormtool"
)

type User struct {
	ID      uint64 `gorm:"column:id"`
	Name    string `gorm:"column:name"`
	Comment string `gorm:"column:comment"`
	Code    string `gorm:"column:code"`
}

func (u User) TableName() string {
	return "city"
}

func main() {
	cfg := config.GetConfig()
	initmodule.ConfigToModel("./config.yaml", cfg)
	err := gormtool.InitClient(gormtool.Config{
		Database: "study",
		Host:     "192.168.31.11",
		Port:     3306,
		User:     "john",
		Password: "asdqwe123",
		Retry:    3,
	})
	if err != nil {
		return
	}
	client, err := gormtool.Client()
	if err != nil {
		return
	}
	u := User{}
	err = client.Where("id = ?", 1).Find(&u).Error
	//gin_skeleton.UseMiddleware(middleware.UserPromGatewayMiddleware("mall-order-svc", "http://192.168.31.16:9091/"))
	gin_skeleton.StartServer(&cfg.HttpServerConfig)

	//gin_skeleton.StartHttpServer(nil)
}

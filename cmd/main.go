package main

import (
	"gitlab.myshuju.top/heshiying/gin-skeleton/initmodule"
)

func main() {
	r := initmodule.Gin()
	initmodule.StartHttpServer(r)

}

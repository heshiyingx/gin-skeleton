package main

import (
	"gitlab.myshuju.top/heshiying/gin-skeleton/cmd/initmodule"
)

func main() {
	r := initmodule.Gin()

	initmodule.StartHttpServer(r)

}

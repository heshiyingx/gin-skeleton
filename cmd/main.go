package main

import (
	"log"
	"os"
	"os/signal"
	"skeleton/cmd/initmodule"
	"skeleton/cmd/startEnd"
	"syscall"
)

func main() {
	gin := initmodule.Gin()

	startEnd.StartGoroutineHttpServer(gin)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	pid := os.Getpid()
	log.Println("pid为", pid)
	<-c
	startEnd.EndHttpServer()
}

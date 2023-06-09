package main

import (
	"os"
	"os/signal"
	"pandora-wechat/application"
	_ "pandora-wechat/common"
	"syscall"
)

func main() {
	application.Run()
	defer exit()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func exit() {
	application.Exit()
}

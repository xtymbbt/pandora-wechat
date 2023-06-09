package application

import (
	"bytes"
	"log"
	"os"
	"pandora-wechat/handlers"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/eatmoreapple/openwechat"
)

var (
	wechatBot *openwechat.Bot
)

func Run() {
	wechatBot = openwechat.DefaultBot(openwechat.Desktop)
	wechatBot.MessageHandler = handlers.Handler
	wechatBot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	err := wechatBot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		if err = wechatBot.Login(); err != nil {
			logrus.Errorf("wechat bot login failed, error is: %s", err)
			return
		}
	}
	go detectWechatBotStatus()
}

func detectWechatBotStatus() {
	buffer := bytes.NewBuffer(nil)
	log.SetOutput(buffer)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if !wechatBot.Alive() {
				logrus.Warn("WechatBot is not alive, try to login...")
				err := wechatBot.Login()
				if err != nil {
					logrus.Fatalf("WechatBot login failed, error is: %s", err)
				}
			}
		}
	}()
	go func() {
		err := wechatBot.Block()
		if err != nil {
			logrus.Errorf("WechatBot block failed, error is: %s", err)
			logrus.Info("WechatBot crash reason is: %s", wechatBot.CrashReason())
			os.Exit(1)
		}
	}()
}

func Exit() {
	logrus.Info("WechatBot has received exit signal, exiting...")
	wechatBot.Exit()
	logrus.Info("WechatBot has successfully exited.")
}

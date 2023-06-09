package handlers

import (
	"github.com/sirupsen/logrus"
	"pandora-wechat/pandoraGpt"
	"strings"

	"github.com/eatmoreapple/openwechat"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)

const AiFailMsg = "很抱歉，AI助手暂时出现问题，请稍后再试。"

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
}

// handle 处理消息
func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	return &UserMessageHandler{}
}

// ReplyText 发送文本消息到用户
func (g *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收私聊消息
	sender, err := msg.Sender()
	logrus.Debugf("Received User %v, ID: %s, Text Msg : %v", sender.NickName, sender.ID(), msg.Content)

	// 向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	reply, err := pandoraGpt.Completions("user_"+sender.ID(), requestText, pandoraGpt.AiIdAssistant)
	if err != nil {
		logrus.Errorf("gtp request error: %v \n", err)
		_, err := msg.ReplyText(AiFailMsg)
		if err != nil {
			logrus.Errorf("response user error: %s", err)
			return err
		}
		return err
	}
	if reply == "" {
		return nil
	}

	// 回复用户
	reply = strings.TrimSpace(reply)
	reply = strings.Trim(reply, "\n")
	_, err = msg.ReplyText(reply)
	if err != nil {
		logrus.Errorf("response user error: %v \n", err)
	}
	return err
}

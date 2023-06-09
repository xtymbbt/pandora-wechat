package handlers

import (
	"github.com/sirupsen/logrus"

	"github.com/eatmoreapple/openwechat"
)

// MessageHandlerInterface 消息处理接口
type MessageHandlerInterface interface {
	handle(*openwechat.Message) error
	ReplyText(*openwechat.Message) error
}

type HandlerType string

const (
	GroupHandler = "group"
	UserHandler  = "user"
)

// handlers 所有消息类型类型的处理器
var handlers map[HandlerType]MessageHandlerInterface

func init() {
	handlers = make(map[HandlerType]MessageHandlerInterface)
	handlers[GroupHandler] = NewGroupMessageHandler()
	handlers[UserHandler] = NewUserMessageHandler()
}

// Handler 全局处理入口
func Handler(msg *openwechat.Message) {
	// 是我本人则返回空
	if msg.IsSendByFriend() {
		logrus.Debugf("Received msg : %s", msg.Content)
		switch msg.MsgType {
		case openwechat.MsgTypeText:
		case openwechat.MsgTypeVoice:
			// 语音消息需要转换为文字
		}
		err := handlers[UserHandler].handle(msg)
		if err != nil {
			logrus.Errorf("userHandler handle failed, error is: %s", err)
		}
	}
	// 群消息
	//if msg.IsSendByGroup() {
	//	handlers[GroupHandler].handle(msg)
	//	return
	//}
}

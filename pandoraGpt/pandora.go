package pandoraGpt

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

var (
	conversations = NewExpireMap()
)

const (
	PandoraAddr   = "http://0.0.0.0:8018"
	MODEL         = "text-davinci-002-render-sha"
	EXPIRE        = time.Minute * 10
	AiIdAssistant = "从现在开始，至整个对话结束，你都将扮演一个王总专属AI助手的角色。你可以随意发挥，但请不要离开角色。"
)

func Completions(senderID string, requestText string, aiID string) (reply string, err error) {
	conversation, exist, err := conversations.LoadOrStore(senderID)
	if err != nil {
		return "", err
	}
	var gptResponse *GptResponse
	if !exist {
		conversation = &Conversation{
			ConversationID:  "",
			MessageID:       uuid.NewString(),
			ParentMessageID: uuid.NewString(),
		}
		//gptResponse, err = pandoraTalk(aiID, conversation)
		//if err != nil {
		//	return
		//}
		//conversation.ConversationID = gptResponse.ConversationID
		//conversation.ParentMessageID = conversation.MessageID
		//conversation.MessageID = uuid.NewString()
		gptResponse, err = pandoraTalk(requestText, conversation)
		if err != nil {
			return
		}
		replies := gptResponse.Message.Content.Parts
		for i := 0; i < len(replies); i++ {
			reply += replies[i]
		}
		if err = conversations.Store(senderID, conversation); err != nil {
			return "", err
		}
	} else {
		conversation.ParentMessageID = conversation.MessageID
		conversation.MessageID = uuid.NewString()
		gptResponse, err = pandoraTalk(requestText, conversation)
		if err != nil {
			return
		}
		replies := gptResponse.Message.Content.Parts
		for i := 0; i < len(replies); i++ {
			reply += replies[i]
		}
	}
	return
}

func pandoraTalk(requestText string, conversation *Conversation) (gptResponsePointer *GptResponse, err error) {
	request := GptRequest{
		Prompt:          requestText,
		Model:           MODEL,
		MessageId:       conversation.MessageID,
		ParentMessageId: conversation.ParentMessageID,
		ConversationId:  conversation.ConversationID,
		Stream:          false,
	}
	marshal, err := json.Marshal(request)
	if err != nil {
		logrus.Errorf("json.Marshal request failed, error is: %s", err)
		return nil, err
	}
	payload := bytes.NewReader(marshal)
	response, err := http.Post(PandoraAddr+`/api/conversation/talk`, "application/json", payload)
	if err != nil {
		logrus.Errorf("http.Post to Pandora failed, error is: %s", err)
		return nil, err
	}
	defer response.Body.Close()
	// 读取 response.Body 的内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("ioutil.ReadAll failed, error is: %s", err)
		return nil, err
	}
	if response.StatusCode != 200 {
		logrus.Errorf("response.StatusCode is: %d, body is: %s", response.StatusCode, string(body))
		return nil, err
	}
	var gptResponse GptResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		logrus.Errorf("json.Unmarshal failed, error is: %s", err)
		return nil, err
	}
	conversation.MessageID = gptResponse.Message.ID
	return &gptResponse, nil
}

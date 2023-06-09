package pandoraGpt

// GptRequest
// * `prompt` 提问的内容。
// * `model` 对话使用的模型，通常整个会话中保持不变。
// * `message_id` 消息ID，首次通常使用`str(uuid.uuid4())`来生成一个。
// * `parent_message_id` 父消息ID，首次同样需要生成。之后获取上一条回复的消息ID即可。
// * `conversation_id` 首次对话可不传。`ChatGPT`回复时可获取。
// * `stream` 是否使用流的方式输出内容，默认为：`True`
// 根据上面的参数，构造 struct
type GptRequest struct {
	Prompt          string `json:"prompt"`
	Model           string `json:"model"`
	MessageId       string `json:"message_id"`        // uuid
	ParentMessageId string `json:"parent_message_id"` // uuid
	ConversationId  string `json:"conversation_id"`
	Stream          bool   `json:"stream"`
}

// GptResponse
// {"conversation_id":"e5063fd8-e191-44d3-93cb-bb65e32ba256","error":null,"message":{"author":{"metadata":{},"name":null,"role":"assistant"},"content":{"content_type":"text","parts":["\u4f60\u597d\uff01\u6709\u4ec0\u4e48\u6211\u53ef\u4ee5\u5e2e\u52a9\u4f60\u7684\u5417\uff1f"]},"create_time":1683454174.062439,"end_turn":true,"id":"692b3619-62ff-4e3d-bebc-796db498abe9","metadata":{"finish_details":{"stop":"<|im_end|>","type":"stop"},"message_type":"next","model_slug":"text-davinci-002-render-sha"},"recipient":"all","update_time":null,"weight":1.0}}
type GptResponse struct {
	ConversationID string      `json:"conversation_id"`
	Error          interface{} `json:"error"`
	Message        struct {
		Author struct {
			Metadata struct {
			} `json:"metadata"`
			Name interface{} `json:"name"`
			Role string      `json:"role"`
		} `json:"author"`
		Content struct {
			ContentType string   `json:"content_type"`
			Parts       []string `json:"parts"`
		} `json:"content"`
		CreateTime float64 `json:"create_time"`
		EndTurn    bool    `json:"end_turn"`
		ID         string  `json:"id"`
		Metadata   struct {
			FinishDetails struct {
				Stop string `json:"stop"`
				Type string `json:"type"`
			} `json:"finish_details"`
			MessageType string `json:"message_type"`
			ModelSlug   string `json:"model_slug"`
		} `json:"metadata"`
		Recipient  string      `json:"recipient"`
		UpdateTime interface{} `json:"update_time"`
		Weight     float64     `json:"weight"`
	} `json:"message"`
}

type Conversation struct {
	ConversationID  string `json:"conversation_id"`
	MessageID       string `json:"message_id"`
	ParentMessageID string `json:"parent_message_id"`
}

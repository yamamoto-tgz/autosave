package line

type RequestBody struct {
	To       string        `json:"to"`
	Messages []TextMessage `json:"messages"`
}

type TextMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewTextMessage(text string) TextMessage {
	return TextMessage{
		Type: "text",
		Text: text,
	}
}

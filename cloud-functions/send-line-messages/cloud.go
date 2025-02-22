package saveexpenses

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	"github.com/yamamoto-tgz/autosave/modules/line"
	"github.com/yamamoto-tgz/autosave/modules/pubsubdata"
)

func init() {
	functions.CloudEvent("send-line-messages", sendLineMessages)
}

func sendLineMessages(ctx context.Context, e event.Event) error {
	var p pubsubdata.PubsubData
	json.Unmarshal(e.DataEncoded, &p)

	d, err := p.DataDecoded()
	if err != nil {
		return err
	}
	fmt.Printf("Pubsub data: %s\n", d)

	var messages []line.TextMessage
	json.Unmarshal(d, &messages)

	body := line.RequestBody{
		To:       os.Getenv("LINE_USER_ID"),
		Messages: messages,
	}

	res, err := post(body)
	if err != nil {
		return err
	}

	fmt.Printf("Response status code: %d\n", res.StatusCode)

	return nil
}

func post(body line.RequestBody) (*http.Response, error) {
	jsn, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Request body: %s\n", jsn)

	req, err := http.NewRequest("POST", "https://api.line.me/v2/bot/message/push", bytes.NewBuffer(jsn))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Line-Retry-Key", uuid.New().String())
	req.Header.Set("Authorization", "Bearer "+os.Getenv("LINE_TOKEN"))

	return http.DefaultClient.Do(req)
}

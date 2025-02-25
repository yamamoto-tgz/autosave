package receivegmailhistory

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/packages/push"
)

func init() {
	functions.CloudEvent("receive-gmail-history", receiveGmailHistory)
}

func receiveGmailHistory(ctx context.Context, e event.Event) error {
	var h History
	err := json.Unmarshal(push.ExtractData(e), &h)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Received history id: %d", h.HistoryId)
	fmt.Println(msg)

	p := push.NewDefaultLinePusher()
	err = p.SendLineMessage(ctx, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

type History struct {
	EmailAddress string `json:"emailAddress"`
	HistoryId    uint32 `json:"historyId"`
}

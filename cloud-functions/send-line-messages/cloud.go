package saveexpenses

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/line"
	"github.com/yamamoto-tgz/autosave/modules/push"
)

func init() {
	functions.CloudEvent("send-line-messages", sendLineMessages)
}

func sendLineMessages(ctx context.Context, e event.Event) error {
	txt := string(push.ExtractData(e))
	fmt.Printf("Received text: %s\n", txt)

	l := &line.User{
		Id:    os.Getenv("LINE_USER_ID"),
		Token: os.Getenv("LINE_TOKEN"),
	}
	l.SendTextMessage(txt)

	return nil
}

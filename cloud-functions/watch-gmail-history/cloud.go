package watchgmailhistory

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/packages/oauth"
	"github.com/yamamoto-tgz/autosave/packages/push"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func init() {
	functions.CloudEvent("watch-gmail-history", watchGmailHistory)
}

func watchGmailHistory(ctx context.Context, e event.Event) error {
	cl, err := oauth.NewDefaultClient(ctx)
	if err != nil {
		return err
	}

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(cl))
	if err != nil {
		return err
	}

	request := &gmail.WatchRequest{
		LabelIds:            []string{os.Getenv("RKTN_PAY_LABEL"), os.Getenv("RKTN_DEBIT_LABEL")},
		LabelFilterBehavior: "include",
		TopicName:           "projects/autosave-tgz/topics/receive-gmail-history",
	}

	res, err := srv.Users.Watch("me", request).Do()
	if err != nil {
		return err
	}

	fmt.Printf("Status code: : %d\n", res.HTTPStatusCode)

	msg := fmt.Sprintf("Start watching from: %d", res.HistoryId)
	fmt.Println(msg)

	p := push.NewDefaultLinePusher()
	err = p.SendLineMessage(ctx, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

package watchgmailhistory

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/oauth"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var BUCKET_NAME string = os.Getenv("BUCKET_NAME")
var CREDENTIALS_JSON string = os.Getenv("CREDENTIALS_JSON")
var TOKEN_JSON string = os.Getenv("TOKEN_JSON")
var PROJECT_NAME string = os.Getenv("PROJECT_NAME")
var TOPIC_NAME string = os.Getenv("TOPIC_NAME")

func init() {
	if BUCKET_NAME == "" {
		BUCKET_NAME = "autosave-tgz"
	}
	if CREDENTIALS_JSON == "" {
		CREDENTIALS_JSON = "credentials.json"
	}
	if TOKEN_JSON == "" {
		TOKEN_JSON = "token.json"
	}
	if PROJECT_NAME == "" {
		PROJECT_NAME = "autosave-tgz"
	}
	if TOPIC_NAME == "" {
		TOPIC_NAME = "gmail"
	}
	functions.CloudEvent("watch-gmail-history", watchGmailHistory)
}

func watchGmailHistory(ctx context.Context, e event.Event) error {
	cl, err := oauth.NewClient(ctx, BUCKET_NAME, CREDENTIALS_JSON, TOKEN_JSON)
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
		TopicName:           fmt.Sprintf("projects/%s/topics/%s", PROJECT_NAME, TOPIC_NAME),
	}

	res, err := srv.Users.Watch("me", request).Do()
	if err != nil {
		return err
	}

	fmt.Printf("HistoryId: %d\n", res.HistoryId)

	return nil
}

package watchgmailhistory

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
	srv, err := auth(ctx)
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

func auth(ctx context.Context) (*gmail.Service, error) {
	config, err := readAuthConfig(ctx)
	if err != nil {
		return nil, err
	}

	tkn, err := readAuthToken(ctx)
	if err != nil {
		return nil, err
	}

	cl := config.Client(ctx, tkn)
	return gmail.NewService(ctx, option.WithHTTPClient(cl))
}

func readFileFromStorage(ctx context.Context, bucketName string, fileName string) (io.Reader, error) {
	cl, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := cl.Bucket(bucketName).Object(fileName)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return r, nil
}

func readAuthConfig(ctx context.Context) (*oauth2.Config, error) {
	r, err := readFileFromStorage(ctx, BUCKET_NAME, CREDENTIALS_JSON)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return google.ConfigFromJSON(bytes, gmail.GmailReadonlyScope)
}

func readAuthToken(ctx context.Context) (*oauth2.Token, error) {
	r, err := readFileFromStorage(ctx, BUCKET_NAME, TOKEN_JSON)
	if err != nil {
		return nil, err
	}

	tkn := &oauth2.Token{}
	err = json.NewDecoder(r).Decode(tkn)
	if err != nil {
		return nil, err
	}

	return tkn, nil
}

package savegmailhistory

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

var BUCKET_NAME = "autosave-tgz"
var HISTORY_FILE = "history.txt"

type PubsubData struct {
	Message struct {
		Data string `json:"data"`
	}
}

type History struct {
	Id           int    `json:"historyId"`
	EmailAddress string `json:"emailAddress"`
}

func init() {
	functions.CloudEvent("save-gmail-history", saveGmailHistory)
}

func saveGmailHistory(ctx context.Context, e event.Event) error {
	var pdata PubsubData
	e.DataAs(&pdata)

	data, err := base64.URLEncoding.DecodeString(pdata.Message.Data)
	if err != nil {
		return err
	}

	var h History
	err = json.Unmarshal(data, &h)
	if err != nil {
		return err
	}

	writeHistoryIdToStorage(ctx, BUCKET_NAME, HISTORY_FILE, h.Id)

	return nil
}

func writeHistoryIdToStorage(ctx context.Context, bucketName string, fileName string, historyId int) (int, error) {
	cl, err := storage.NewClient(ctx)
	if err != nil {
		return 0, err
	}

	w := cl.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	defer w.Close()

	return fmt.Fprintf(w, "%d", historyId)
}

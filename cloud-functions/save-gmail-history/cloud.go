package savegmailhistory

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/pubsubdata"
)

var BUCKET_NAME = "autosave-tgz"
var HISTORY_FILE = "history.txt"

func init() {
	functions.CloudEvent("save-gmail-history", saveGmailHistory)
}

func saveGmailHistory(ctx context.Context, e event.Event) error {
	var p pubsubdata.PubsubData
	err := json.Unmarshal(e.DataEncoded, &p)
	if err != nil {
		return err
	}

	data, err := p.DataDecoded()
	if err != nil {
		return err
	}

	var history struct {
		Id           int    `json:"historyId"`
		EmailAddress string `json:"emailAddress"`
	}

	err = json.Unmarshal(data, &history)
	if err != nil {
		return err
	}

	writeHistoryIdToStorage(ctx, BUCKET_NAME, HISTORY_FILE, history.Id)

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

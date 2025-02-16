package savegmailhistory

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/pubsubdata"
)

var BUCKET_NAME = os.Getenv("BUCEKT_NAME")
var HISTORY_TXT = os.Getenv("HISTORY_TXT")

func init() {
	if BUCKET_NAME == "" {
		BUCKET_NAME = "autosave-tgz"
	}
	if HISTORY_TXT == "" {
		HISTORY_TXT = "history.txt"
	}
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

	var h struct {
		Id           int    `json:"historyId"`
		EmailAddress string `json:"emailAddress"`
	}

	err = json.Unmarshal(data, &h)
	if err != nil {
		return err
	}

	cl, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	w := cl.Bucket(BUCKET_NAME).Object(HISTORY_TXT).NewWriter(ctx)
	defer w.Close()

	_, err = fmt.Fprintf(w, "%d", h.Id)
	if err != nil {
		return err
	}

	fmt.Printf("historyId: %d\n", h.Id)

	return nil
}

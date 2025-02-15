package savegmailhistory

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
)

func TestSaveGmailHistory(t *testing.T) {
	data := `{"historyId":12345,"emailAddress":"autosave@yamamoto.tgz"}`
	b64 := base64.URLEncoding.EncodeToString([]byte(data))
	jsn := fmt.Sprintf(`{"message":{"data":"%s"}}`, b64)

	type PubsubData struct {
		Message struct {
			Data string
		}
	}

	var p PubsubData
	err := json.Unmarshal([]byte(jsn), &p)
	if err != nil {
		t.Fail()
	}

	e := event.New("1.0")
	e.SetData("application/json", p)

	err = saveGmailHistory(context.Background(), e)
	if err != nil {
		t.Error(err)
	}
}

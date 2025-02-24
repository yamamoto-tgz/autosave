package saveexpenses

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/push"
)

func TestSendLineMessages(t *testing.T) {
	p := push.New([]byte("HELLO WORLD"))
	jsn, _ := json.Marshal(p)

	e := event.New("1.0")
	e.SetData("application/json", jsn)

	err := sendLineMessages(context.Background(), e)
	if err != nil {
		t.Error(err)
	}
}

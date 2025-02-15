package watchgmailhistory

import (
	"context"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
)

func TestWatchGmailHistory(t *testing.T) {
	ctx := context.Background()
	var e event.Event
	err := watchGmailHistory(ctx, e)
	if err != nil {
		t.Error(err)
	}
}

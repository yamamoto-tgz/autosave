package push

import (
	"context"
	"testing"
)

func TestSendLineMessage(t *testing.T) {
	l := LinePusher{
		ProjectId: "autosave-tgz",
		TopicId:   "send-line-messages",
	}

	err := l.SendLineMessage(context.Background(), "HELLO WORLD")
	if err != nil {
		t.Error(err)
	}
}

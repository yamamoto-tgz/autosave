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

	txt := []byte("HELLO WORLD")
	err := l.SendLineMessage(context.Background(), txt)
	if err != nil {
		t.Error(err)
	}
}

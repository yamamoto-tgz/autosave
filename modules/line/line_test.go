package line

import (
	"os"
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	l := &User{
		Id:    os.Getenv("LINE_USER_ID"),
		Token: os.Getenv("LINE_TOKEN"),
	}

	err := l.SendTextMessage("HELLO WORLD")
	if err != nil {
		t.Error(err)
	}
}

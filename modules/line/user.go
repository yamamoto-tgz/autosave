package line

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type User struct {
	Id    string
	Token string
}

func (u *User) SendTextMessage(txt string) error {
	body := &requestBody{
		To: u.Id,
		Messages: []message{{
			Type: "text",
			Text: txt,
		}},
	}

	res, err := u.post(*body)
	if err != nil {
		return err
	}

	fmt.Printf("Status code: %d\n", res.StatusCode)
	return nil
}

func (u *User) post(body requestBody) (*http.Response, error) {
	jsn, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Request body: %s\n", jsn)

	req, err := http.NewRequest("POST", "https://api.line.me/v2/bot/message/push", bytes.NewBuffer(jsn))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Line-Retry-Key", uuid.New().String())
	req.Header.Set("Authorization", "Bearer "+os.Getenv("LINE_TOKEN"))

	return http.DefaultClient.Do(req)
}

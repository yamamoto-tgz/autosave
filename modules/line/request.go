package line

type requestBody struct {
	To       string    `json:"to"`
	Messages []message `json:"messages"`
}

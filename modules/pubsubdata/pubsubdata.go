package pubsubdata

import "encoding/base64"

type PubsubData struct {
	Message struct {
		Data string
	}
}

func (p PubsubData) DataDecoded() ([]byte, error) {
	data, err := base64.URLEncoding.DecodeString(p.Message.Data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func New(data string) *PubsubData {
	b64 := base64.URLEncoding.EncodeToString([]byte(data))
	return &PubsubData{
		Message: struct {
			Data string
		}{
			Data: b64,
		},
	}
}

package oauth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func NewClient(ctx context.Context, bucketName string, credentialJson string, tokenJson string) (*http.Client, error) {
	config, err := readAuthConfig(ctx, bucketName, credentialJson)
	if err != nil {
		return nil, err
	}

	tkn, err := readAuthToken(ctx, bucketName, tokenJson)
	if err != nil {
		return nil, err
	}

	return config.Client(ctx, tkn), nil
}

func readFileFromStorage(ctx context.Context, bucketName string, fileName string) (io.Reader, error) {
	cl, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := cl.Bucket(bucketName).Object(fileName)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return r, nil
}

func readAuthConfig(ctx context.Context, bucketName string, credentialsJson string) (*oauth2.Config, error) {
	r, err := readFileFromStorage(ctx, bucketName, credentialsJson)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return google.ConfigFromJSON(bytes, gmail.GmailReadonlyScope)
}

func readAuthToken(ctx context.Context, bucketName string, tokenJson string) (*oauth2.Token, error) {
	r, err := readFileFromStorage(ctx, bucketName, tokenJson)
	if err != nil {
		return nil, err
	}

	tkn := &oauth2.Token{}
	err = json.NewDecoder(r).Decode(tkn)
	if err != nil {
		return nil, err
	}

	return tkn, nil
}

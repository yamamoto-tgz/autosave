package oauth

import (
	"context"
	"testing"
)

func TestAuth(t *testing.T) {
	ctx := context.Background()
	_, err := NewClient(ctx, "autosave-tgz", "credentials.json", "token.json")
	if err != nil {
		t.Error(err)
	}
}

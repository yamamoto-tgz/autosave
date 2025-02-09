package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	hostname := ""
	if localOnly := os.Getenv("LOCAL_ONLY"); localOnly == "true" {
		hostname = "127.0.0.1"
	}

	if err := funcframework.StartHostPort(hostname, port); err != nil {
		log.Fatalf("funcframework.StartHostPort: %v\n", err)
	}
}

func init() {
	functions.HTTP("hello", helloHandler)
	functions.CloudEvent("event", eventHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HELLO WORLD")
}

func eventHandler(ctx context.Context, e event.Event) error {
	var data struct {
		Message struct {
			Data []byte `json:"data"`
		}
	}
	e.DataAs(&data)

	message := string(data.Message.Data)
	fmt.Println(message)

	return nil
}

package saveexpenses

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/expense"
	"github.com/yamamoto-tgz/autosave/modules/oauth"
	"github.com/yamamoto-tgz/autosave/modules/pubsubdata"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var BUCKET_NAME string = os.Getenv("BUCKET_NAME")
var CREDENTIALS_JSON string = os.Getenv("CREDENTIALS_JSON")
var TOKEN_JSON string = os.Getenv("TOKEN_JSON")
var SPREADSHEET_ID = os.Getenv("SPREADSHEET_ID")
var RANGE = os.Getenv("RANGE")

func init() {
	if BUCKET_NAME == "" {
		BUCKET_NAME = "autosave-tgz"
	}
	if CREDENTIALS_JSON == "" {
		CREDENTIALS_JSON = "credentials.json"
	}
	if TOKEN_JSON == "" {
		TOKEN_JSON = "token.json"
	}
	functions.CloudEvent("save-expenses", saveExpenses)
}

func saveExpenses(ctx context.Context, e event.Event) error {
	var p pubsubdata.PubsubData
	json.Unmarshal(e.DataEncoded, &p)

	b, err := p.DataDecoded()
	if err != nil {
		return err
	}

	var exps []expense.Expense
	json.Unmarshal(b, &exps)

	cl, err := oauth.NewClient(ctx, BUCKET_NAME, CREDENTIALS_JSON, TOKEN_JSON)
	if err != nil {
		return err
	}

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(cl))
	if err != nil {
		return err
	}

	var values [][]interface{}
	for _, ex := range exps {
		values = append(values, []interface{}{"=ROW()-1", ex.Date, ex.Category, ex.Amount, ex.Description, ex.Source})
	}
	fmt.Printf("values: %s\n", values)

	vrange := &sheets.ValueRange{
		Values: values,
	}

	res, err := srv.Spreadsheets.Values.Append(SPREADSHEET_ID, RANGE, vrange).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		return err
	}

	fmt.Printf("status: %d\n", res.HTTPStatusCode)

	return nil
}

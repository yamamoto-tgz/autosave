package saveexpenses

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/yamamoto-tgz/autosave/modules/oauth"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func init() {
	functions.CloudEvent("save-expenses", saveExpenses)
}

func saveExpenses(ctx context.Context, e event.Event) error {
	var exps []Expense
	json.Unmarshal(e.DataEncoded, &exps)

	cl, err := oauth.NewClient(ctx, "autosave-tgz", "credentials.json", "token.json")
	if err != nil {
		return err
	}

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(cl))
	if err != nil {
		return err
	}

	var vls [][]interface{}
	for _, ex := range exps {
		vls = append(vls, []interface{}{"=ROW()-1", ex.Date, ex.Category, ex.Amount, ex.Description, ex.Source})
	}

	vrng := &sheets.ValueRange{
		Values: vls,
	}

	id := os.Getenv("SPREADSHEET_ID")
	rng := "expenses!A2:F100"
	res, err := srv.Spreadsheets.Values.Append(id, rng, vrng).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		return err
	}

	fmt.Printf("status: %d\n", res.HTTPStatusCode)

	return nil
}

type Expense struct {
	Id          uint16
	Date        string
	Category    string
	Amount      uint16
	Description string
	Source      string
}

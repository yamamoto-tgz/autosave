package saveexpenses

import (
	"context"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
)

func TestSaveGmailHistory(t *testing.T) {
	ex := []Expense{
		{
			Id:          1,
			Date:        "2025-01-01T00:00:00.000+09:00",
			Category:    "Test",
			Amount:      1000,
			Description: "This is test 01",
			Source:      "Test",
		},
		{
			Id:          2,
			Date:        "2025-01-02T00:00:00.000+09:00",
			Category:    "Test",
			Amount:      2000,
			Description: "This is test 02",
			Source:      "Test",
		},
	}

	e := event.New("1.0")
	e.SetData("application/json", ex)

	err := saveExpenses(context.Background(), e)
	if err != nil {
		t.Error(err)
	}
}

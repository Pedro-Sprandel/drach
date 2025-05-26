package commands

import (
	"drach/db"
	"drach/helpers"
	"drach/models"
	"log"
	"time"
)

func SummaryCmd(args []string) {
	summaries, categoryTotals, monthlyTotals, grandTotal, err := models.GetExpenseSummary(db.DB, time.Now().Year())
	if err != nil {
		log.Fatal(err)
	}

	helpers.PrintExpenseSummary(summaries, categoryTotals, monthlyTotals, grandTotal)
}

package commands

import (
	"drach/db"
	"drach/helpers"
	"drach/models"
	"flag"
	"log"
	"time"
)

func SummaryCmd(args []string) {
	fs := flag.NewFlagSet("summary", flag.ExitOnError)

	month := fs.Int("month", 0, "Month of summary, integer 1-12")
	fs.IntVar(month, "m", 0, "Month of summary, integer 1-12")

	year := fs.Int("year", time.Now().Year(), "Year of summary, integer 1-12")
	fs.IntVar(year, "y", time.Now().Year(), "Year of summary, integer 1-12")

	summaries, categoryTotals, monthlyTotals, grandTotal, err := models.GetExpenseSummary(db.DB, *year, *month)
	if err != nil {
		log.Fatal(err)
	}

	helpers.PrintExpenseSummary(summaries, categoryTotals, monthlyTotals, grandTotal)
}

package commands

import (
	"flag"
	"log"
	"time"

	"drach/db"
	"drach/helpers"
	"drach/services"
)

func SummaryCmd(args []string) {
	fs := flag.NewFlagSet("summary", flag.ExitOnError)

	categoryID := fs.Int("category", 0, "Category Filter, integer 1-n")
	fs.IntVar(categoryID, "c", 0, "Category Filter, integer 1-n")

	month := fs.Int("month", 0, "Month of summary, integer 1-12")
	fs.IntVar(month, "m", 0, "Month of summary, integer 1-12")

	currentYear := time.Now().Year()
	year := fs.Int("year", currentYear, "Year of summary, integer 1-12")
	fs.IntVar(year, "y", currentYear, "Year of summary, integer 1-12")

	summaries, categoryTotals, monthlyTotals, grandTotal, err := services.GetExpenseSummary(db.DB, *categoryID, *year, *month)
	if err != nil {
		log.Fatal(err)
	}

	helpers.PrintExpenseSummary(summaries, categoryTotals, monthlyTotals, grandTotal)
}

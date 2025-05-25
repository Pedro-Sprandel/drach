package commands

import (
	"drach/db"
	"drach/helpers"
	"drach/models"
	"flag"
	"fmt"
	"log"
)

func ListCmd(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)

	category := fs.String("category", "", "Filter by category")
	fs.StringVar(category, "c", "", "Filter by category")

	month := fs.Int("month", 0, "Month of expense, integer")
	fs.IntVar(month, "m", 0, "Month of expense, integer")

	year := fs.Int("year", 0, "Year of expense, integer")
	fs.IntVar(year, "y", 0, "Year of expense, integer")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	expenses, err := models.ListExpenses(db.DB, *category, *month, *year)
	if err != nil {
		log.Fatalf("Error on list expense: %v", err)
	}

	helpers.PrintExpenses(expenses)
}

package commands

import (
	"flag"
	"fmt"
	"log"

	"drach/db"
	"drach/helpers"
	"drach/models"
)

func ListCmd(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)

	categoryID := fs.Int("category", 0, "Filter by category")
	fs.IntVar(categoryID, "c", 0, "Filter by category")

	month := fs.Int("month", 0, "Month of expense, integer")
	fs.IntVar(month, "m", 0, "Month of expense, integer")

	year := fs.Int("year", 0, "Year of expense, integer")
	fs.IntVar(year, "y", 0, "Year of expense, integer")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	expenses, err := models.ListExpenses(db.DB, *categoryID, *month, *year)
	if err != nil {
		log.Fatalf("Error on list expense: %v", err)
	}

	helpers.PrintExpenses(expenses)
}

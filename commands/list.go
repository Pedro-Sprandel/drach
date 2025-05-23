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
	cmd := flag.NewFlagSet("list", flag.ExitOnError)
	category := cmd.String("category", "", "Filter by category")
	cmd.StringVar(category, "c", "", "Filter by category")

	if err := cmd.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	expenses, err := models.ListExpenses(db.DB, *category)
	if err != nil {
		log.Fatalf("Error on list expense: %v", err)
	}

	helpers.PrintExpenses(expenses)
}

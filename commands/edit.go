package commands

import (
	"drach/db"
	"drach/models"
	"flag"
	"fmt"
	"log"
	"os"
)

func EditCmd(args []string) {
	fs := flag.NewFlagSet("edit", flag.ExitOnError)

	id := fs.String("id", "", "ID of the expense to edit")

	description := fs.String("description", "", "Description of expense, string")
	fs.StringVar(description, "d", "", "Alias for --description")

	category := fs.String("category", "", "Category of expense for summary purposes")
	fs.StringVar(category, "c", "", "Alias for --category")

	amount := fs.Float64("amount", 0, "Value of expense, integer")
	fs.Float64Var(amount, "a", 0, "Value of expense, integer")

	if err := fs.Parse(args); err != nil {
		fmt.Print("Error parsing flags")
	}

	if *id == "" {
		fmt.Println("Error: ID is required")
		fs.Usage()
		os.Exit(1)
	}

	if *description == "" && *amount == 0 && *category == "" {
		fmt.Println("Error: at least one of --description, --amount or --category must be provided")
		fs.Usage()
		os.Exit(1)
	}

	err := models.EditExpense(db.DB, *id, *description, *category, *amount)
	if err != nil {
		log.Fatalf("Error on edit expense %v:", err)
	}
}

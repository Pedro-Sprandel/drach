package commands

import (
	"flag"
	"fmt"
	"log"
	"os"

	"drach/db"
	"drach/services"
)

func EditCmd(args []string) {
	fs := flag.NewFlagSet("edit", flag.ExitOnError)

	id := fs.String("id", "", "ID of the expense to edit")

	description := fs.String("description", "", "Description of expense, string")
	fs.StringVar(description, "d", "", "Alias for --description")

	categoryID := fs.Int("category", 0, "Category of expense for summary purposes, integer")
	fs.IntVar(categoryID, "c", 0, "Alias for --category")

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

	if *description == "" && *amount == 0 && *categoryID == 0 {
		fmt.Println("Error: at least one of --description, --amount or --category must be provided")
		fs.Usage()
		os.Exit(1)
	}

	err := services.EditExpense(db.DB, *id, *description, *categoryID, *amount)
	if err != nil {
		log.Fatalf("Error on edit expense %v:", err)
	}
}

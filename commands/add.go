package commands

import (
	"drach/db"
	"drach/models"
	"flag"
	"fmt"
	"log"
)

const defaultValueCategory = "Sem categoria"

func AddCmd(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)

	description := fs.String("description", "", "Description of expense, string")
	fs.StringVar(description, "d", "", "Alias for --description")

	amount := fs.Float64("amount", 0, "Value of expense, integer")
	fs.Float64Var(amount, "a", 0, "Alias for --amount")

	category := fs.String("category", "", "Category of expense for summary purposes")
	fs.StringVar(category, "c", "", "Alias for --category")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	if *category == "" {
		*category = defaultValueCategory
	}

	err := models.AddExpense(db.DB, *description, *amount, *category)
	if err != nil {
		log.Fatalf("Error on add expense: %v", err)
	}

	fmt.Println("Expense added successfully")
}

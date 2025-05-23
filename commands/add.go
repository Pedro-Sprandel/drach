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
	cmd := flag.NewFlagSet("add", flag.ExitOnError)

	description := cmd.String("description", "", "Description of expense, string")
	cmd.StringVar(description, "d", "", "Alias for --description")

	amount := cmd.Float64("amount", 0, "Value of expense, integer")
	cmd.Float64Var(amount, "a", 0, "Alias for --amount")

	category := cmd.String("category", "", "Category of expense for summary purposes")
	cmd.StringVar(category, "c", "", "Alias for --category")

	if err := cmd.Parse(args); err != nil {
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

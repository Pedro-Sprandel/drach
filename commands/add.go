package commands

import (
	"drach/db"
	"drach/helpers"
	"drach/models"
	"flag"
	"fmt"
	"log"
	"os"
)

func AddCmd(args []string) {
	cmd := flag.NewFlagSet("add", flag.ExitOnError)

	amount := cmd.Float64("amount", 0, "Value of expense, integer")
	cmd.Float64Var(amount, "a", 0, "Alias for --amount")

	category := cmd.String("category", "", "Category of expense for summary purposes")
	cmd.StringVar(category, "c", "", "Alias for --category")

	if helpers.FlagProvided("category", cmd) && *category == "" {
		fmt.Print("Error: Category cannot be empty")
		cmd.Usage()
		os.Exit(1)
	}

	if err := cmd.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	err := models.AddExpense(db.DB, *amount, *category)
	if err != nil {
		log.Fatalf("Error on add expense: %v", err)
	}

	fmt.Println("Expense added successfully")
}

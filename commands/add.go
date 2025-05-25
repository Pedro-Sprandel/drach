package commands

import (
	"drach/db"
	"drach/helpers"
	"drach/models"
	"flag"
	"fmt"
	"log"
)

const defaultValueDescription = "Sem categoria"
const defaultValueCategory = "Sem categoria"

func AddCmd(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)

	description := fs.String("description", defaultValueDescription, "Description of expense, string")
	fs.StringVar(description, "d", defaultValueDescription, "Alias for --description")

	amount := fs.Float64("amount", 0, "Value of expense, integer")
	fs.Float64Var(amount, "a", 0, "Alias for --amount")

	category := fs.String("category", defaultValueCategory, "Category of expense for summary purposes")
	fs.StringVar(category, "c", defaultValueCategory, "Alias for --category")

	currentMonth := helpers.CurrentMonth()
	month := fs.Int("month", currentMonth, "Month of expense, integer")
	fs.IntVar(month, "m", currentMonth, "Month of expense, integer")

	currentYear := helpers.CurrentYear()
	year := fs.Int("year", currentYear, "Year of expense, integer")
	fs.IntVar(year, "y", currentYear, "Year of expense, integer")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	err := models.AddExpense(db.DB, *description, *amount, *category, *month, *year)
	if err != nil {
		log.Fatalf("Error on add expense: %v", err)
	}

	fmt.Println("Expense added successfully")
}

package commands

import (
	"flag"
	"fmt"
	"log"
	"os"

	"drach/db"
	"drach/helpers"
	"drach/services"
)

const (
	defaultValueDescription = "Sem descrição"
	defaultValueCategory    = "Sem categoria"
)

func AddCmd(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)

	description := fs.String("description", defaultValueDescription, "Description of expense, string")
	fs.StringVar(description, "d", defaultValueDescription, "Alias for --description")

	amount := fs.Float64("amount", 0, "Value of expense, integer")
	fs.Float64Var(amount, "a", 0, "Alias for --amount")

	categoryID := fs.Int("category", 0, "Category of expense (optional - will prompt if not provided)")
	fs.IntVar(categoryID, "c", 0, "Alias for --category")

	currentMonth := helpers.CurrentMonth()
	month := fs.Int("month", currentMonth, "Month of expense, integer")
	fs.IntVar(month, "m", currentMonth, "Month of expense, integer")

	currentYear := helpers.CurrentYear()
	year := fs.Int("year", currentYear, "Year of expense, integer")
	fs.IntVar(year, "y", currentYear, "Year of expense, integer")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	if *year < 2000 || *year > currentYear+1 {
		fmt.Println("Invalid year")
		fs.Usage()
		os.Exit(1)
	}

	var err error

	if *categoryID == 0 {
		*categoryID, err = helpers.SelectCategory()
		if err != nil {
			log.Fatalf("Erro ao selecionar categoria: %v", err)
		}
	}

	err = services.AddExpense(db.DB, *description, *amount, *categoryID, *month, *year)
	if err != nil {
		log.Fatalf("Error on add expense: %v", err)
	}

	fmt.Println("Expense added successfully")
}

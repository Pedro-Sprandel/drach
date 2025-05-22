package commands

import (
	// "drach/db"
	// "drach/models"
	"drach/helpers"
	"flag"
	"fmt"
	"os"
)

func AddCmd(args []string) {
	cmd := flag.NewFlagSet("add", flag.ExitOnError)

	amount := cmd.Int("amount", 0, "Value of expense, integer")
	cmd.IntVar(amount, "a", 0, "Alias for --amount")

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

	fmt.Println("Called add")
	fmt.Println(*amount)
	fmt.Println(*category)
}

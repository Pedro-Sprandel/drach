package commands

import (
	"flag"
	"fmt"
	"log"
	"os"

	"drach/db"
	"drach/services"
)

func RemoveCmd(args []string) {
	cmd := flag.NewFlagSet("remove", flag.ExitOnError)

	id := cmd.String("id", "", "ID of item to remove")

	if err := cmd.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	if *id == "" {
		fmt.Print("Error: ID cannot be empty")
		cmd.Usage()
		os.Exit(1)
	}

	err := services.RemoveExpense(db.DB, *id)
	if err != nil {
		log.Fatalf("Error on add expense: %v", err)
	}

	fmt.Printf("Removed expense with id = %v", *id)
}

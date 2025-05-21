package commands

import (
	"drach/db"
	"drach/models"
	"flag"
	"fmt"
	"log"
)

func AddCmd(args []string) {
	cmd := flag.NewFlagSet("add", flag.ExitOnError)
	description := cmd.String("d", "", "Task description")

	if err := cmd.Parse(args); err != nil {
		fmt.Printf("Error on args parsing: %v", err)
	}

	if *description == "" {
		fmt.Println("Error: description must not be empty")
		cmd.Usage()
		return
	}

	err := models.AddTask(db.DB, *description)
	if err != nil {
		log.Fatalf("Error on add task: %v", err)
	}

	fmt.Println("Task added successfully")
}

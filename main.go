package main

import (
	"fmt"
	"log"
	"os"

	"drach/commands"
	"drach/db"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error on database initialization: %v", err)
	}

	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Printf("Error on close db connection: %v", err)
		}
	}()

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "list":
		commands.ListCmd(os.Args[2:])
	case "add":
		commands.AddCmd(os.Args[2:])
	case "edit":
		commands.EditCmd(os.Args[2:])
	case "remove":
		commands.RemoveCmd(os.Args[2:])
	case "summary":
		commands.SummaryCmd(os.Args[2:])
	case "category":
		commands.CategoryCmd(os.Args[2:])
	case "env":
		commands.EnvCmd(os.Args[2:])
	case "help":
		printHelp()
	default:
		fmt.Println("Unkwown command")
		fmt.Println()
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage: drach <command> [options]")
	fmt.Println()
	fmt.Println("Commands list:")
	fmt.Println()
	fmt.Println("Add new task: add -d <description> -c <category> -a <amount>* -m <month> -y <year>")
	fmt.Println("List tasks: list -c <category> -m <month> -y <year>")
	fmt.Println("Edit task: edit -id <id>* -d <description> -c <category> -a <amount>* -m <month> -y <year>")
	fmt.Println("Edit task: edit -id <id>* -d <description> -c <category> -a <amount>* -m <month> -y <year>")
	fmt.Println("Remove task: remove -id <id>*")
	fmt.Println("Environment: env switch <env> | status")
}

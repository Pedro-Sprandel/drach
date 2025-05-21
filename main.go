package main

import (
	"drach/commands"
	"drach/db"
	"fmt"
	"log"
	"os"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
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
	case "add":
		commands.AddCmd(os.Args[2:])
	case "list":
		commands.ListCmd()
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage: drach <command> [options]")
	fmt.Println("Commands list:")
	fmt.Println("  add -d \"description\"  Add new task")
	fmt.Println("  list                 List all tasks")
}

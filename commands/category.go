package commands

import (
	"flag"
	"fmt"
	"log"
	"os"

	"drach/db"
	"drach/services"
)

func CategoryCmd(args []string) {
	switch args[0] {
	case "list":
		services.ListCategories(db.DB)
	case "remove":
		fs := flag.NewFlagSet("remove", flag.ExitOnError)
		ID := fs.String("id", "", "ID of item to remove")

		if err := fs.Parse(args[1:]); err != nil {
			fmt.Printf("Error parsing flags")
		}

		if *ID == "" {
			fmt.Print("Error: ID cannot be empty")
			fs.Usage()
			os.Exit(1)
		}

		if err := fs.Parse(args); err != nil {
			fmt.Printf("Error parsing flags")
		}

		err := services.RemoveCategory(db.DB, *ID)
		if err != nil {
			log.Fatal("Erro ao deleter categoria", err)
		} else {
			fmt.Printf("Categoria [%s] deletada com successo", *ID)
		}
	}
}

package commands

import (
	"drach/helpers"
	"flag"
	"fmt"
	"os"
)

func RemoveCmd(args []string) {
	cmd := flag.NewFlagSet("remove", flag.ExitOnError)

	id := cmd.String("id", "", "ID of item to remove")

	if err := cmd.Parse(args); err != nil {
		fmt.Printf("Error parsing flags")
	}

	if helpers.FlagProvided(*id, cmd) && *id == "" {
		fmt.Print("Error: ID cannot be empty")
		cmd.Usage()
		os.Exit(1)
	}
}

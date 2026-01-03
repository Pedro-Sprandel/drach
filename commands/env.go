package commands

import (
	"fmt"
	"os"
)

func EnvCmd(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: env <command> [options]")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("  switch <env>   Switch environment (prod|test)")
		fmt.Println("  status         Show current environment")
		return
	}

	switch args[0] {
	case "switch":
		if len(args) < 2 {
			fmt.Println("Usage: env switch <env>")
			return
		}
		switchEnv(args[1])
	case "status":
		printStatus()
	default:
		fmt.Println("Unknown environment command")
	}
}

func switchEnv(env string) {
	if env != "prod" && env != "test" {
		fmt.Println("Invalid environment. Use 'prod' or 'test'.")
		return
	}

	err := os.WriteFile(".drach_config", []byte(env), 0644)
	if err != nil {
		fmt.Printf("Error switching environment: %v\n", err)
		return
	}

	fmt.Printf("Switched to environment: %s\n", env)
}

func printStatus() {
	config, err := os.ReadFile(".drach_config")
	env := "prod"
	if err == nil {
		env = string(config)
	}
	fmt.Printf("Current environment: %s\n", env)
}

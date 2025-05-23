package helpers

import (
	"drach/models"
	"fmt"
)

func PrintExpenses(expenses []models.Expense) {
	var sum float64 = 0

	fmt.Printf("%-5s | %-24s | %-15s | %-2s | %-4s | %-7s\n", "ID", "Description", "Category", "Month", "Year", "Amount")
	fmt.Println("------|--------------------------|-----------------|-------|------|--------")
	for _, expense := range expenses {
		sum += expense.Amount
		fmt.Printf(
			"%-5d | %-24s | %-15s | %-5s | %-4d | R$%-4.2f\n",
			expense.ID,
			expense.Description,
			expense.Category,
			MonthName(expense.Month),
			expense.Year,
			expense.Amount,
		)
	}
	fmt.Printf("\nTotal: R$%-4.2f", sum)
}

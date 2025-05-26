package helpers

import (
	"drach/models"
	"fmt"
	"strings"
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

func PrintExpenseSummary(summaries []models.ExpenseSummary, categoryTotals map[string]float64, monthlyTotals map[int]float64, grandTotal float64) {
	fmt.Println("\n=== RESUMO DE GASTOS ===")
	fmt.Printf("%-15s | %-10s | %-10s | %-12s\n", "Categoria", "Mês", "Ano", "Total")
	fmt.Println(strings.Repeat("-", 45))

	for _, s := range summaries {
		fmt.Printf("%-15s | %-10s | %-10d | R$ %9.2f\n",
			s.Category,
			MonthName(s.Month),
			s.Year,
			s.TotalAmount)
	}

	fmt.Println("\n=== TOTAIS POR CATEGORIA ===")
	for category, total := range categoryTotals {
		fmt.Printf("%-20s: R$ %9.2f\n", category, total)
	}

	fmt.Println("\n=== TOTAIS MENSAIS ===")
	for month, total := range monthlyTotals {
		fmt.Printf("%-10s: R$ %9.2f\n", MonthName(month), total)
	}

	fmt.Println(strings.Repeat("=", 30))
	fmt.Printf("%-10s: R$ %9.2f\n", "TOTAL ANUAL", grandTotal)
	fmt.Println(strings.Repeat("=", 30))
}

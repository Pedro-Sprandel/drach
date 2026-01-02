package helpers

import (
	"fmt"
	"strings"

	"drach/services"
)

const (
	width  = 34
	indent = "    "
)

func Right(text string) string {
	if len(text) >= width {
		return indent + text
	}
	leftPad := width - len(text)
	return indent + strings.Repeat(" ", leftPad) + text
}

func Center(text string) string {
	if len(text) >= width {
		return indent + text
	}
	padding := width - len(text)
	left := padding / 2
	right := padding - left
	return indent + strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}

func Line(char string) string {
	return indent + strings.Repeat(char, width)
}

func LineWithWidth(char string, w int) string {
	return indent + strings.Repeat(char, w)
}

func PrintExpenses(expenses []services.ExpenseWithCategoryName) {
	var sum float64 = 0

	fmt.Println()
	fmt.Println()
	fmt.Print(indent + fmt.Sprintf("%-5s | %-16s | %-12s | %-8s | %-7s\n", "ID", "Descrição", "Categoria", "Mês/Ano", "Valor"))
	fmt.Println(LineWithWidth("-", 60))
	for _, expense := range expenses {
		sum += expense.Amount
		fmt.Print(indent + fmt.Sprintf(
			"%-5d | %-16s | %-12s | %-8s | R$%-4.2f\n",
			expense.ID,
			expense.Description,
			expense.CategoryName,
			fmt.Sprintf("%s/%d", MonthName(expense.Month), expense.Year),
			expense.Amount,
		))
	}
	fmt.Printf("\n%sTotal: R$%-4.2f", indent, sum)
}

func PrintExpenseSummary(summaries []services.ExpenseSummary, categoryTotals map[string]float64, monthlyTotals map[int]float64, grandTotal float64) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(Center("RESUMO DE GASTOS"))
	fmt.Println(Line("="))
	fmt.Print(indent + fmt.Sprintf("%-10s | %-8s | %-12s\n", "Categoria", "Mês/Ano", "Total"))
	fmt.Println(Line("-"))

	for _, s := range summaries {
		fmt.Print(indent + fmt.Sprintf("%-10s | %-8s | R$ %7.2f\n",
			s.CategoryName,
			fmt.Sprintf("%s/%d", MonthName(s.Month), s.Year),
			s.TotalAmount))
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println(Center("TOTAIS POR CATEGORIA"))
	fmt.Println(Line("="))
	for category, total := range categoryTotals {
		line := fmt.Sprintf("%-10s R$ %7.2f\n", category+":", total)
		fmt.Print(Right(line))
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println(Center("TOTAIS MENSAIS"))
	fmt.Println(Line("="))
	for month, total := range monthlyTotals {
		line := fmt.Sprintf("%-3s: R$ %7.2f\n", MonthName(month), total)
		fmt.Print(Right(line))
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println(Center("TOTAL ANUAL"))
	fmt.Println(Line("="))
	fmt.Print(Right(fmt.Sprintf("R$ %7.2f\n", grandTotal)))
}

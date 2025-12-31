package models

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

type Expense struct {
	ID          int
	Description string
	Amount      float64
	CategoryID  string
	Month       int
	Year        int
	CreatedAt   time.Time
}

func AddExpense(db *sql.DB, description string, amount float64, categoryID int64, month int, year int) error {
	_, err := db.Exec(
		"INSERT INTO expenses(description, amount, category_id, month, year) VALUES (?, ?, ?, ?, ?)",
		description,
		amount,
		categoryID,
		month,
		year,
	)

	return err
}

func ListExpenses(db *sql.DB, categoryID string, month int, year int) ([]Expense, error) {
	query := `
        SELECT id, description, amount, category_id, month, year, created_at 
        FROM expenses
    `

	var args []any
	filters := []string{}

	if month > 0 {
		filters = append(filters, "month = ?")
		args = append(args, month)
	}

	if year > 0 {
		filters = append(filters, "year = ?")
		args = append(args, year)
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	var expenses []Expense
	for rows.Next() {
		var e Expense

		err := rows.Scan(&e.ID, &e.Description, &e.Amount, &e.CategoryID, &e.Month, &e.Year, &e.CreatedAt)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, e)
	}

	return expenses, nil
}

func EditExpense(db *sql.DB, id string, description string, categoryID int64, amount float64) error {
	var query strings.Builder
	query.WriteString("UPDATE expenses SET ")

	var args []any
	var updates []string

	if description != "" {
		updates = append(updates, "description = ?")
		args = append(args, description)
	}

	if amount != 0 {
		updates = append(updates, "amount = ?")
		args = append(args, amount)
	}

	// MUDAR DEPOIS
	if categoryID > 0 {
		updates = append(updates, "category = ?")
		args = append(args, category)
	}

	query.WriteString(strings.Join(updates, ", "))
	query.WriteString(" WHERE id = ?")
	args = append(args, id)

	_, err := db.Exec(query.String(), args...)
	return err
}

func RemoveExpense(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM expenses WHERE id = ?", id)

	return err
}

type ExpenseSummary struct {
	Category    string
	Month       int
	Year        int
	TotalAmount float64
}

func GetExpenseSummary(db *sql.DB, year int, month ...int) ([]ExpenseSummary, map[string]float64, map[int]float64, float64, error) {
	query := `
        SELECT 
            category, 
            month, 
            year, 
            SUM(amount) as total_amount
        FROM 
            expenses
        WHERE 
            year = ?
    `

	args := []any{year}

	if len(month) > 0 {
		query += " AND month = ?"
		args = append(args, month[0])
	}

	query += `
        GROUP BY 
            category, month, year
        ORDER BY 
            year, month, category
    `

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, nil, nil, 0, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	var summaries []ExpenseSummary
	categoryTotals := make(map[string]float64)
	monthlyTotals := make(map[int]float64)
	var grandTotal float64

	for rows.Next() {
		var s ExpenseSummary
		err := rows.Scan(&s.Category, &s.Month, &s.Year, &s.TotalAmount)
		if err != nil {
			return nil, nil, nil, 0, err
		}

		summaries = append(summaries, s)
		categoryTotals[s.Category] += s.TotalAmount
		monthlyTotals[s.Month] += s.TotalAmount
		grandTotal += s.TotalAmount
	}

	if len(month) > 0 {
		monthTotal := monthlyTotals[month[0]]
		return summaries, categoryTotals, nil, monthTotal, nil
	}

	return summaries, categoryTotals, monthlyTotals, grandTotal, nil
}

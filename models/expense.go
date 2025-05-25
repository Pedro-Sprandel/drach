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
	Category    string
	Month       int
	Year        int
	CreatedAt   time.Time
}

func AddExpense(db *sql.DB, description string, amount float64, category string, month int, year int) error {
	_, err := db.Exec(
		"INSERT INTO expenses(description, amount, category, month, year) VALUES (?, ?, ?, ?, ?)",
		description,
		amount,
		category,
		month,
		year,
	)

	return err
}

func ListExpenses(db *sql.DB, category string) ([]Expense, error) {
	rows, err := db.Query("SELECT id, description, amount, category, month, year, created_at FROM expenses")
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

		err := rows.Scan(&e.ID, &e.Description, &e.Amount, &e.Category, &e.Month, &e.Year, &e.CreatedAt)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, e)
	}

	return expenses, nil
}

func EditExpense(db *sql.DB, id string, description string, category string, amount float64) error {
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

	if category != "" {
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

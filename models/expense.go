package models

import (
	"database/sql"
	"log"
	"time"
)

type Expense struct {
	ID          int
	Description string
	Amount      float64
	Category    string
	CreatedAt   time.Time
}

func AddExpense(db *sql.DB, description string, amount float64, category string) error {
	_, err := db.Exec("INSERT INTO expenses(description, amount, category) VALUES (?, ?, ?)", description, amount, category)

	return err
}

func ListExpenses(db *sql.DB, category string) ([]Expense, error) {
	rows, err := db.Query("SELECT id, description, amount, category, created_at FROM expenses")
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

		err := rows.Scan(&e.ID, &e.Description, &e.Amount, &e.Category, &e.CreatedAt)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, e)
	}

	return expenses, nil
}

func RemoveExpense(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM expenses WHERE id = ?", id)

	return err
}

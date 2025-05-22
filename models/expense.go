package models

import (
	"database/sql"
)

type Expense struct {
	ID       int
	Amount   float64
	Category string
}

func AddExpense(db *sql.DB, amount float64, category string) error {
	_, err := db.Exec("INSERT INTO expenses(amount, category) VALUES (?, ?)", amount, category)
	return err
}

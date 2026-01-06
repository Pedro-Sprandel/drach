package services

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
	CategoryID  int
	Month       int
	Year        int
	CreatedAt   time.Time
}

type ExpenseWithCategoryName struct {
	Expense
	CategoryName string
}

func AddExpense(db *sql.DB, description string, amount float64, categoryID int, month int, year int) error {
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

func ListExpenses(db *sql.DB, categoryID int, month int, year int) ([]ExpenseWithCategoryName, error) {
	query := `
  		SELECT 
  			e.id, e.description, e.amount, e.category_id, 
  			c.name, e.month, e.year, e.created_at 
  		FROM expenses e	
  		LEFT JOIN categories c ON e.category_id = c.id
    `

	var args []any
	filters := []string{}

	if month > 0 {
		filters = append(filters, "e.month = ?")
		args = append(args, month)
	}

	if year > 0 {
		filters = append(filters, "e.year = ?")
		args = append(args, year)
	}

	if categoryID > 0 {
		filters = append(filters, "e.category_id = ?")
		args = append(args, categoryID)
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	query += " ORDER BY e.created_at DESC"

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	var expenses []ExpenseWithCategoryName
	for rows.Next() {
		var e ExpenseWithCategoryName

		err := rows.Scan(&e.ID, &e.Description, &e.Amount, &e.CategoryID, &e.CategoryName, &e.Month, &e.Year, &e.CreatedAt)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, e)
	}

	return expenses, nil
}

func EditExpense(db *sql.DB, id string, description string, categoryID int, amount float64, year int, month int) error {
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

	if categoryID > 0 {
		updates = append(updates, "category_id = ?")
		args = append(args, categoryID)
	}

	if year != 0 {
		updates = append(updates, "year = ?")
		args = append(args, year)
	}

	if month != 0 {
		updates = append(updates, "month = ?")
		args = append(args, month)
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
	CategoryName string
	CategoryID   int
	Month        int
	Year         int
	Amount       float64
	TotalAmount  float64
}

func GetExpenseSummary(db *sql.DB, categoryID int, year int, month int) ([]ExpenseSummary, map[string]float64, map[int]float64, float64, error) {
	query := `
        SELECT 
  					c.name,
  					e.category_id,
  					e.month,
  					e.year,
            e.amount, 
            SUM(e.amount) as total_amount
        FROM 
            expenses e
  			LEFT JOIN categories c ON e.category_id = c.id
        WHERE 
            year = ?
  `

	args := []any{year}

	if month != 0 {
		query += " AND month = ?"
		args = append(args, month)
	}

	if categoryID != 0 {
		query += " AND category_id = ?"
		args = append(args, categoryID)
	}

	query += `
        GROUP BY 
            category_id, month, year
        ORDER BY 
            year, month, category_id
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
		err := rows.Scan(&s.CategoryName, &s.CategoryID, &s.Month, &s.Year, &s.Amount, &s.TotalAmount)
		if err != nil {
			return nil, nil, nil, 0, err
		}

		summaries = append(summaries, s)
		categoryTotals[s.CategoryName] += s.TotalAmount
		monthlyTotals[s.Month] += s.TotalAmount
		grandTotal += s.TotalAmount
	}

	if month != 0 {
		monthTotal := monthlyTotals[month]
		return summaries, categoryTotals, nil, monthTotal, nil
	}

	return summaries, categoryTotals, monthlyTotals, grandTotal, nil
}

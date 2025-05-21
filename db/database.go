package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	dbPath := "drach.db"
	var err error

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if _, err := os.Create(dbPath); os.IsNotExist(err) {
			return fmt.Errorf("error to craete os %v", err)
		}
	}

	DB, err = sql.Open("sqlite3", dbPath+"?_loc=auto&_time_format=sqlite")
	if err != nil {
		return fmt.Errorf("error to open database: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		created_at DATATIME DEFAULT CURRENT_TIMESTAMP,
		completed BOOLEAN DEFAULT FALSE
	)`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error to creatte database: %v", err)
	}

	return nil
}

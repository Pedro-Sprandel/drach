package models

import (
	"database/sql"
	"log"
	"time"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	Completed   bool
}

func AddTask(db *sql.DB, description string) error {
	_, err := db.Exec("INSERT INTO tasks (description) VALUES (?)", description)
	return err
}

func ListTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, description, created_at, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()
	var tasks []Task
	for rows.Next() {
		var t Task
		var createdAtStr string

		err := rows.Scan(&t.ID, &t.Description, &createdAtStr, &t.Completed)
		if err != nil {
			return nil, err
		}

		t.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

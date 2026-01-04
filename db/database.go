package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type TableDefinition struct {
	Name    string
	Schema  string
	Indexes []string
}

var tables = []TableDefinition{
	{
		Name: "categories",
		Schema: `(
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						name TEXT NOT NULL,
						description TEXT,
						created_at DATETIME DEFAULT CURRENT_TIMESTAMP
				)`,
		Indexes: []string{
			"CREATE INDEX IF NOT EXISTS idx_categories_name ON categories(name)",
		},
	},
	{
		Name: "expenses",
		Schema: `(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            description VARCHAR(50) NOT NULL,
            amount DECIMAL(10, 2) NOT NULL,
        		category_id INTEGER NOT NULL,  
            month INTEGER NOT NULL,
            year INTEGER NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT
        )`,
		Indexes: []string{
			"CREATE INDEX IF NOT EXISTS idx_expenses_category_id ON expenses(category_id)",
			"CREATE INDEX IF NOT EXISTS idx_expenses_created_at ON expenses(created_at)",
		},
	},
}

var DB *sql.DB

func GetCurrentEnv() string {
	config, err := os.ReadFile(".drach_config")
	if err != nil {
		return "prod"
	}
	return string(config)
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(home, path[2:])
	}
	return path
}

func InitDB() error {
	env := GetCurrentEnv()
	dbPath := expandPath("~/Documentos/projects/drach")

	if env == "test" {
		dbPath += "/drach_test.db"
	} else {
		dbPath += "/drach.db"
	}
	fmt.Println(dbPath)

	var err error

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if _, err := os.Create(dbPath); os.IsNotExist(err) {
			return fmt.Errorf("error to craete os %v", err)
		}
	}

	DB, err = sql.Open("sqlite3", dbPath+"?_foreign_keys=on&_loc=auto&_time_format=sqlite")
	if err != nil {
		return fmt.Errorf("error to open database: %v", err)
	}

	_, err = DB.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return fmt.Errorf("failed to enable foreign keys: %v", err)
	}

	for _, table := range tables {
		_, err := DB.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s %s", table.Name, table.Schema))
		if err != nil {
			return fmt.Errorf("failed to create table %s: %v", table.Name, err)
		}

		for _, indexSQL := range table.Indexes {
			_, err := DB.Exec(indexSQL)
			if err != nil {
				return fmt.Errorf("failed to create index for table %s: %v", table.Name, err)
			}
		}
	}

	return nil
}

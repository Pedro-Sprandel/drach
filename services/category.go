package services

import (
	"database/sql"
	"log"
)

func AddCategory(db *sql.DB, name string, description string) (sql.Result, error) {
	result, err := db.Exec("INSERT INTO categories(name, description) VALUES (?, ?)", name, description)

	return result, err
}

func GetAllCategories(db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT id, name, description FROM categories ORDER BY name")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error on closing rows at category %v", err)
		}
	}()

	var categories []map[string]interface{}
	for rows.Next() {
		var id int64
		var name string
		var description sql.NullString

		err := rows.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}

		category := map[string]interface{}{
			"id":   id,
			"name": name,
		}
		if description.Valid {
			category["description"] = description.String
		}

		categories = append(categories, category)
	}

	return categories, nil
}

package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Category struct {
	ID   int64
	Name string
}

func AddCategory(db *sql.DB, name string) (sql.Result, error) {
	result, err := db.Exec("INSERT INTO categories(name) VALUES (?)", name)

	return result, err
}

func EditCategory(db *sql.DB, id string, name string) (sql.Result, error) {
	result, err := db.Exec("UPDATE categories SET name = ? WHERE id = ?", name, id)

	return result, err
}

func ListCategories(db *sql.DB) {
	categories, err := GetAllCategories(db)
	if err != nil {
		log.Fatal("Erro ao listar categorias")
	}

	for _, category := range categories {
		fmt.Printf("[%d] %s\n", category.ID, category.Name)
	}
}

func GetAllCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories ORDER BY id")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error on closing rows at category %v", err)
		}
	}()

	var categories []Category
	for rows.Next() {
		var category Category

		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func RemoveCategory(db *sql.DB, ID string) error {
	_, err := db.Exec("DELETE FROM categories WHERE id = ?", ID)

	return err
}

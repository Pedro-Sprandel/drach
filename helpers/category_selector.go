package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"drach/db"
	"drach/models"
)

func SelectCategory() (int, error) {
	categories, err := models.GetAllCategories(db.DB)
	if err != nil {
		return 0, fmt.Errorf("erro ao buscar categorias: %v", err)
	}

	fmt.Println("\n=== Selecione uma categoria ===")
	fmt.Println()

	for i, cat := range categories {
		fmt.Printf("  [%d] %s", i+1, cat["name"])
		if desc, ok := cat["description"].(string); ok && desc != "" {
			fmt.Printf(" - %s", desc)
		}
		fmt.Println()
	}

	fmt.Printf("  [%d] Adicionar nova categoria\n", len(categories)+1)
	fmt.Println()
	fmt.Print("Escolha uma opção: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf("opção inválida")
	}

	if choice < 1 || choice > len(categories)+1 {
		return 0, fmt.Errorf("opção fora do intervalo")
	}

	if choice == len(categories)+1 {
		return createNewCategory()
	}

	selectedCat := categories[choice-1]
	return selectedCat["id"].(int), nil
}

func createNewCategory() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== Nova Categoria ===")
	fmt.Print("Nome da categoria: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	name = strings.TrimSpace(name)

	if name == "" {
		return 0, fmt.Errorf("nome da categoria não pode ser vazio")
	}

	fmt.Print("Descrição (opcional, pressione Enter para pular): ")
	description, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	description = strings.TrimSpace(description)

	result, err := models.AddCategory(db.DB, name, description)
	if err != nil {
		return 0, fmt.Errorf("erro ao criar categoria: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Printf("\n✅ Categoria '%s' criada com sucesso!\n", name)
	return int(id), nil
}

func GetCategoryName(categoryID int) (string, error) {
	var name string
	err := db.DB.QueryRow("SELECT name FROM categories WHERE id = ?", categoryID).Scan(&name)
	return name, err
}

package crud

import (
	"context"
	"fmt"
	"os"

	"github.com/ikiruhawk/food-store/database/connection"
	"github.com/ikiruhawk/food-store/internal/models"
)

func GetProducts() ([]models.Product, error) {
	conn := connection.GetConnection()
	defer conn.Close(context.Background())
	products := make([]models.Product, 0)

	rows, err := conn.Query(context.Background(), `SELECT products.product_id, categories.category_id, categories.name, manufacturers.manufacturer_id, manufacturers.name,
													products.name, products.price, products.amount, products.description FROM products, categories, manufacturers
													WHERE categories.category_id = products.category_id AND manufacturers.manufacturer_id = products.manufacturer_id;`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "internal/crud/crud.go GetProducts; Query error: %v\n", err.Error())
		return []models.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			p models.Product
			c models.Category
			m models.Manufacturer
		)

		err := rows.Scan(&p.Id, &c.Id, &c.Name, &m.Id, &m.Name, &p.Name, &p.Price, &p.Amount, &p.Description)
		if err != nil {
			fmt.Fprintf(os.Stderr, "internal/crud/crud.go GetProducts; Querty scan error: %v\n", err.Error())
			return []models.Product{}, err
		}
		p.Category = c
		p.Manufacturer = m
		p.PriceFloat = float64(p.Price) / 100
		products = append(products, p)
	}

	return products, nil
}

func GetCategories() ([]models.Category, error) {
	conn := connection.GetConnection()
	defer conn.Close(context.Background())
	categories := make([]models.Category, 0)

	rows, err := conn.Query(context.Background(), "SELECT category_id, name FROM categories;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "internal/crud/crud.go getCategories; Query error: %v\n", err.Error())
		return []models.Category{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.Id, &c.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "internal/crud/crud.go getCategories; Query scan error: %v\n", err.Error())
			return []models.Category{}, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func GetManufacturers() ([]models.Manufacturer, error) {
	conn := connection.GetConnection()
	defer conn.Close(context.Background())
	manufacturers := make([]models.Manufacturer, 0)

	rows, err := conn.Query(context.Background(), "SELECT manufacturer_id, name FROM manufacturers;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "internal/crud/crud.go getManufacturers; Query error: %v\n", err.Error())
		return []models.Manufacturer{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Manufacturer
		err := rows.Scan(&m.Id, &m.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "internal/crud/crud.go getManufacturers; Query scan error: %v\n", err.Error())
			return []models.Manufacturer{}, err
		}
		manufacturers = append(manufacturers, m)
	}

	return manufacturers, nil
}

func GetProductById(id int) *models.Product {
	p := models.Product{}
	prs, err := GetProducts()
	if err == nil {
		for _, v := range prs {
			if v.Id == id {
				p = v
				break
			}
		}
	}
	return &p
}

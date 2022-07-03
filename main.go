package main

import (
	"database/sql"

	db "github.com/ayrtonbsouza/hexagonal-architecture-poc/adapters/database"
	"github.com/ayrtonbsouza/hexagonal-architecture-poc/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "db.sqlite")
	productDatabaseAdapter := db.NewProductDb(database)
	productService := application.NewProductService(productDatabaseAdapter)
	product, _ := productService.Create("Product 1", 20)
	productService.Enable(product)
}

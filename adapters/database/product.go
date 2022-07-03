package database

import (
	"database/sql"

	"github.com/ayrtonbsouza/hexagonal-architecture-poc/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDatabase struct {
	database *sql.DB
}

func NewProductDb(database *sql.DB) *ProductDatabase {
	return &ProductDatabase{database: database}
}

func (p *ProductDatabase) Get(id string) (application.IProduct, error) {
	var product application.Product
	stmt, err := p.database.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

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

func (p *ProductDatabase) Save(product application.IProduct) (application.IProduct, error) {
	var rows int
	p.database.QueryRow("select id from products where id=?", product.GetId()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDatabase) create(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.database.Prepare(`insert into products (id, name, price, status) values (?,?,?,?);`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDatabase) update(product application.IProduct) (application.IProduct, error) {
	_, err := p.database.Exec(
		"update products set name=?, price=?, status=? where id=?;",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId(),
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}

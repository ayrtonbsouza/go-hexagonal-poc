package database

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func setUp() {
	Database, _ = sql.Open("sqlite3", ":memory:")
	createTable(Database)
	createProduct(Database)
}

func createTable(database *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string,
	);`

	stmt, err := database.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(database *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES ("abc", "Test Product", 0, "disabled");`
	stmt, err := database.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

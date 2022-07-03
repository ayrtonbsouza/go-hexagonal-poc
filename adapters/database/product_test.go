package database_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ayrtonbsouza/hexagonal-architecture-poc/adapters/database"
	"github.com/stretchr/testify/require"
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
			"status" string
			);`

	stmt, err := database.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(database *sql.DB) {
	insert := `insert into products values("abc", "Test Product", 0, "disabled");`
	stmt, err := database.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDatabase_Get(t *testing.T) {
	setUp()
	defer Database.Close()
	productDb := database.NewProductDb(Database)

	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Test Product", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/Codeflix-FullCycle/hexagonal-architecture/adapters/db"
	"github.com/Codeflix-FullCycle/hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTale(Db)
	createProduct(Db)
}

func createTale(db *sql.DB) {
	table := `CREATE TABLE products( 
		"id" varchar,
		"name" varchar,
		"status" varchar,
		"price" float
		);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec()

	if err != nil {
		log.Fatalln(err)
	}
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("123", "first product", "disabled", 0)`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestProduct_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	product := db.NewProduct(Db)

	result, err := product.Get("123")

	require.Nil(t, err)

	require.Equal(t, result.GetID(), "123")
	require.Equal(t, result.GetStatus(), "disabled")
}

func TestProduct_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDB := db.NewProduct(Db)

	product := application.NewProduct()

	product.Name = "product 1"
	product.Status = application.ENABLED
	product.Price = 10

	productDB.Save(product)

	result, err := productDB.Get(product.ID)

	require.Nil(t, err)
	require.Equal(t, result, product)

	product.Name = "product 2"

	productDB.Save(product)

	result, err = productDB.Get(product.ID)

	fmt.Println(result)

	require.Nil(t, err)
	require.Equal(t, result.GetName(), "product 2")

}

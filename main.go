package main

import (
	"database/sql"

	"github.com/Codeflix-FullCycle/hexagonal-architecture/adapters/db"
	"github.com/Codeflix-FullCycle/hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlDb, _ := sql.Open("sqlite3", "db.sqlite")
	productsAdapter := db.NewProduct(sqlDb)

	productsService := application.NewProductService(productsAdapter)

	product, _ := productsService.Create("Product test", 10)

	productsService.Enable(product)
}

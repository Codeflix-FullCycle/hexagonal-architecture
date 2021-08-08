package db

import (
	"database/sql"

	"github.com/Codeflix-FullCycle/hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	DB *sql.DB
}

func NewProduct(db *sql.DB) *ProductDb {
	return &ProductDb{
		DB: db,
	}
}
func (p *ProductDb) Get(id string) (application.ProductsInterface, error) {
	stmt, err := p.DB.Prepare("SELECT id, name, status, price FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	var product application.Products
	if err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.Price); err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductsInterface) (application.ProductsInterface, error) {
	var rows int
	p.DB.QueryRow("SELECT count(id) FROM products WHERE id = ?", product.GetID()).Scan(&rows)

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

func (p *ProductDb) create(product application.ProductsInterface) (application.ProductsInterface, error) {
	stmt, err := p.DB.Prepare(`insert into products (id, name, status, price) values($1, $2, $3, $4)`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetStatus(),
		product.GetPrice(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductsInterface) (application.ProductsInterface, error) {
	_, err := p.DB.Exec("update products set name = ?, status = ?, price = ? where id = ? ",
		product.GetName(),
		product.GetStatus(),
		product.GetPrice(),
		product.GetID())

	if err != nil {
		return nil, err
	}

	return product, nil
}

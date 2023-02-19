package models

import (
	"fmt"
	"loja/db"
)
type Product struct {
	Id 			int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

const SQL_INSERT_ONE = "INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)"
const SQL_FIND_ALL = "SELECT * FROM produtos"
const SQL_DELETE_BY_ID = "DELETE FROM produtos WHERE id = $1"

func FindAll() []Product {
	products := []Product{}
	db := db.GetConnection()

	rows, err := db.Query(SQL_FIND_ALL)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, Product{id, name, description, price, quantity})
	}

	defer db.Close()
	return products

}

func Create(product Product) {

	fmt.Println(product)

	db := db.GetConnection()

	stmt, err := db.Prepare(SQL_INSERT_ONE)

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(product.Name, product.Description, product.Price, product.Quantity)
	defer db.Close()


}

func Delete(id int) {
	fmt.Println(id)

	db := db.GetConnection()

	stmt, err := db.Prepare(SQL_DELETE_BY_ID)

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(id)
	defer db.Close()
}

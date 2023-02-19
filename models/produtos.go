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
const SQL_FIND_BY_ID = "SELECT * FROM produtos WHERE id = $1"
const SQL_UPDATE_ONE = "UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5"

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

func Save(product Product) {

	fmt.Println(product)
	db := db.GetConnection()

	if product.Id == 0 {
		stmt, err := db.Prepare(SQL_INSERT_ONE)

		if err != nil {
			panic(err.Error())
		}
	
		stmt.Exec(product.Name, product.Description, product.Price, product.Quantity)
		
	}

	stmt, err := db.Prepare(SQL_UPDATE_ONE)

		if err != nil {
			panic(err.Error())
		}
	
		stmt.Exec(product.Name, product.Description, product.Price, product.Quantity, product.Id)
	

	
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

func FindById(id int) Product{
	fmt.Println(id)

	db := db.GetConnection()

	rows, err := db.Query(SQL_FIND_BY_ID, id)

	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		product = Product{id, name, description, price, quantity}
	}


	defer db.Close()
	return product
}
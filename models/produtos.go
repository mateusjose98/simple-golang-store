package models

import (
	"loja/db"
)
type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

const SQL_FIND_ALL = "SELECT * FROM produtos"

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
		products = append(products, Product{name, description, price, quantity})
	}

	defer db.Close()
	return products

}

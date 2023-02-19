package controllers

import (
	"loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.FindAll()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price, errPrice := strconv.ParseFloat(r.FormValue("preco") ,64)
		quantity, errQuantity := strconv.Atoi(r.FormValue("quantidade")) 

		if errPrice != nil {
			print("Erro")
		}

		if errQuantity != nil {
			print("Erro")
		}

		models.Create(models.Product{Name: name, Description: description, Price: price, Quantity: quantity})


	}
	http.Redirect(w, r, "/", http.StatusFound)
}
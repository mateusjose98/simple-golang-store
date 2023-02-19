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

		models.Save(models.Product{Name: name, Description: description, Price: price, Quantity: quantity})


	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id")) 
	models.Delete(id)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id")) 

	produto := models.FindById(id)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id, _:= strconv.Atoi(r.FormValue("id"))
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

		models.Save(models.Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity})


	}


	http.Redirect(w, r, "/", http.StatusFound)
}


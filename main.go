package main

import (
	"html/template"
	"net/http"
)

type Products struct {
	Nome        string
	Description string
	Price       float64
	Qnt         int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Products{
		{
			Nome:        "Camiseta",
			Description: "Descrição do Produto 1",
			Price:       9.99,
			Qnt:         5,
		},
		{
			Nome:        "Shorts",
			Description: "Descrição do Produto 2",
			Price:       19.99,
			Qnt:         10,
		},
		{
			Nome:        "Meia",
			Description: "Descrição do Produto 3",
			Price:       14.99,
			Qnt:         3,
		},
		{
			Nome:        "Chuteira",
			Description: "Descrição do Produto 4",
			Price:       24.99,
			Qnt:         8,
		},
	}

	temp.ExecuteTemplate(w, "Index", products)
}

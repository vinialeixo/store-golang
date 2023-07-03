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
			Description: "Descrição da Camiseta: Esta camiseta de algodão de alta qualidade é perfeita para uso diário. Disponível em diferentes tamanhos e cores.",
			Price:       9.99,
			Qnt:         5,
		},
		{
			Nome:        "Shorts",
			Description: "Descrição dos Shorts: Estes shorts leves são ideais para atividades esportivas. Feitos de um material resistente e confortável.",
			Price:       19.99,
			Qnt:         10,
		},
		{
			Nome:        "Meia",
			Description: "Descrição da Meia: Esta meia macia e respirável mantém os pés confortáveis durante o dia todo. Disponível em diferentes cores e tamanhos.",
			Price:       14.99,
			Qnt:         3,
		},
		{
			Nome:        "Chuteira",
			Description: "Descrição da Chuteira: Esta chuteira de alta performance oferece excelente aderência e estabilidade no campo. Ideal para jogadores profissionais.",
			Price:       24.99,
			Qnt:         8,
		},
	}

	temp.ExecuteTemplate(w, "Index", products)
}

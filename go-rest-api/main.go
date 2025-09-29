package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	// pgadmin 		- http://localhost:54321/browser/
	// localhost 	- http://localhost:8000/

	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{ID: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}

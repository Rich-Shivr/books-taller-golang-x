package main

import (
	"log"
	"net/http"
	"serveractualizado/server/cmd/handlers"
	"serveractualizado/server/cmd/repositories"
	"serveractualizado/server/cmd/routes"
)

func main() {
	sqliteRepo, err := repositories.NewSqliteBookRepository("books.db")
	if err != nil {
		log.Fatal(err)
	}
	err = sqliteRepo.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	bookHandler := handlers.NewBookHandler(sqliteRepo)
	routes.SetupPingRoute()
	routes.SetupBookRoutes(bookHandler)
	log.Println("Server is running on port 3030")
	http.ListenAndServe(":3030", nil)
}

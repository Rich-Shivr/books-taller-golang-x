package main

import (
	"log"
	"net/http"
	"server/cmd/handlers"
	"server/cmd/repositories"
	"server/cmd/routes"
)

func main() {
	bookRepo := repositories.NewInMemoryBookRepository()
	bookHandler := handlers.NewBookHandler(bookRepo)
	routes.SetupPingRoute()
	routes.SetupBookRoutes(bookHandler)
	log.Println("Server is running on port 3030")
	http.ListenAndServe(":3030", nil)
}

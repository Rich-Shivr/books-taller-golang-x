package routes

import (
	"net/http"
	"serveractualizado/server/cmd/handlers"
)

func SetupPingRoute() {
	http.HandleFunc("/ping", handlers.PingHandler)
}

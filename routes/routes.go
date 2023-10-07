package routes

import (
	"net/http"
	"russianwords/controllers" // Ensure this matches your actual module and package name
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/send-glossary", controllers.SendGlossaryHandler)
}

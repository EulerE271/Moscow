package routes

import (
	"net/http"
	"russianwords/controllers" // Ensure this matches your actual module and package name
	"russianwords/middleware"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/send-glossary", middleware.RateLimitMiddleware(controllers.SendGlossaryHandler))
}

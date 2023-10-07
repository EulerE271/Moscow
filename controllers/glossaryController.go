package controllers

import (
	"encoding/json"
	"net/http"
	"strconv" // Import strconv package for string conversion

	"russianwords/emailer"
)

func SendGlossaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Email         string `json:"email"`
		GlossaryCount string `json:"glossaryCount"` // Keep this as string to initially decode from JSON
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Convert GlossaryCount to integer
	glossaryCount, err := strconv.Atoi(request.GlossaryCount)
	if err != nil {
		http.Error(w, "Invalid glossary count", http.StatusBadRequest)
		return
	}

	// Trigger the function to send the email with the glossary to the user's email.
	err = emailer.SendGlossary(request.Email, glossaryCount) // Pass integer value here
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully!"})
}

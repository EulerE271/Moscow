package controllers

import (
	"encoding/json"
	"net/http"

	"russianwords/emailer"
)

func SendGlossaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Trigger the function to send the email with the glossary to the user's email.
	err := emailer.SendGlossary(request.Email) // Make sure SendMail has the correct signature and parameters.
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully!"})
}

package controllers

import (
	"e-com/src/dto"
	"e-com/src/helper"
	"e-com/src/services"
	"encoding/json"
	"net/http"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	helper.HandleHeader(w)
	if r.Method != "POST" {
		w.WriteHeader(404)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input dto.UserLoginSchema
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	token, err := services.LoginService(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful!",
		"token":   token,
	})

}

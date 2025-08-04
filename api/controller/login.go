package controllers

import (
	"e-com/internal"
	usecase "e-com/usecase"
	"encoding/json"
	"net/http"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)
	if r.Method != "POST" {
		w.WriteHeader(404)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input internal.UserLoginSchema
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	token, err := usecase.LoginService(input)

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

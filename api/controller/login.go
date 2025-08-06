package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
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
		reponse.Error(w, 400, "Invalid input: "+err.Error(), err)
		return
	}

	token, err := usecase.LoginService(input)

	if err != nil {
		reponse.Error(w, 400, "Login failed!", err)
		return
	}

	reponse.Success(w, 201, "Login successful!", map[string]interface{}{
		"message": "Login successful!",
		"token":   token,
	})

}

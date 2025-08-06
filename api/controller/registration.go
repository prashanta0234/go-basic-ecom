package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	usecase "e-com/usecase"
	"encoding/json"
	"errors"
	"net/http"
)

func RegisterUserController(w http.ResponseWriter, r *http.Request) {

	internal.HandleHeader(w)
	if r.Method != "POST" {
		reponse.Error(w, 404, "Method not allowed", errors.New("method not allowed"))
		return
	}

	var input internal.UserRegistrationSchema
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		reponse.Error(w, 400, "Invalid input: "+err.Error(), err)
		return
	}

	token, err := usecase.RegisterUserService(input)

	if err != nil {
		reponse.Error(w, 400, "Registration failed!", err)
		return
	}

	reponse.Success(w, 201, "Registration successful!", map[string]interface{}{
		"message": "Registration successful!",
		"token":   token,
	})

}

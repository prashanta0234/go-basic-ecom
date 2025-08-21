package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	usecase "e-com/usecase"
	"encoding/json"
	"net/http"
)

// RegisterUserController godoc
// @Summary Register a new user
// @Description Register a new user account with name, email, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body internal.UserRegistrationSchema true "User registration information"
// @Success 201 {object} reponse.SuccessResponse "Registration successful"
// @Failure 400 {object} reponse.ErrorResponse "Registration failed"
// @Router /registration [post]
func RegisterUserController(w http.ResponseWriter, r *http.Request) {

	internal.HandleHeader(w)

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

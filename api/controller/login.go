package controllers

import (
	"e-com/internal"
	"e-com/internal/reponse"
	usecase "e-com/usecase"
	"encoding/json"
	"net/http"
)

// LoginController godoc
// @Summary User login
// @Description Authenticate user with email and password to get JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body internal.UserLoginSchema true "User login credentials"
// @Success 201 {object} reponse.SuccessResponse "Login successful"
// @Failure 400 {object} reponse.ErrorResponse "Login failed"
// @Router /login [post]
func LoginController(w http.ResponseWriter, r *http.Request) {
	internal.HandleHeader(w)

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

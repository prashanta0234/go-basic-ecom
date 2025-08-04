package internal

type UserRegistration struct {
	Name     string
	Email    string
	Password string
}

type UserRegistrationSchema struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"min=6,email"`
	Password string `json:"password" binding:"required,password,min=6"`
}

type UserLoginSchema struct {
	Email    string `json:"email" binding:"min=6,email"`
	Password string `json:"password" binding:"required,password,min=6"`
}

func UserRegistrationDto(data UserRegistrationSchema) UserRegistration {
	return UserRegistration(data)
}

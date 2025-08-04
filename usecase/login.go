package usecase

import (
	"context"
	"e-com/bootstrap"
	domain "e-com/domain"
	"e-com/internal"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(data internal.UserLoginSchema) (string, error) {
	collection := bootstrap.DB.Collection("users")

	var isExisting domain.User

	err := collection.FindOne(context.TODO(), bson.M{"email": data.Email}).Decode(&isExisting)

	if err != nil {
		return "", errors.New("this email is not registerd")

	}

	err = bcrypt.CompareHashAndPassword([]byte(isExisting.Password), []byte(data.Password))

	if err != nil {
		return "", errors.New("sorry wrong password")
	}

	token, err := internal.GenerateJWT(isExisting.ID.Hex())

	if err != nil {
		return "", errors.New("something went wrong")
	}

	return token, nil

}

package services

import (
	"context"
	"e-com/src/dto"
	"e-com/src/helper"
	"e-com/src/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(data dto.UserLoginSchema) (string, error) {
	collection := DB.Collection("users")

	var isExisting models.User

	err := collection.FindOne(context.TODO(), bson.M{"email": data.Email}).Decode(&isExisting)

	if err != nil {
		return "", errors.New("this email is not registerd")

	}

	err = bcrypt.CompareHashAndPassword([]byte(isExisting.Password), []byte(data.Password))

	if err != nil {
		return "", errors.New("sorry wrong password")
	}

	token, err := helper.GenerateJWT(isExisting.ID.Hex())

	if err != nil {
		return "", errors.New("something went wrong")
	}

	return token, nil

}

package services

import (
	"context"
	"e-com/src/dto"
	"e-com/src/helper"
	"e-com/src/models"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserService(data dto.UserRegistrationSchema) (string, error) {

	collection := DB.Collection("users")

	var existing models.User

	err := collection.FindOne(context.TODO(), bson.M{"email": data.Email}).Decode(&existing)

	if err == nil {
		return "", errors.New("email already exists")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", errors.New("failed to process password")
	}

	user := models.User{
		ID:        primitive.NewObjectID(),
		Name:      data.Name,
		Email:     data.Email,
		Password:  string(hashedPass),
		CreatedAt: time.Now(),
	}

	_, err = collection.InsertOne(context.TODO(), &user)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return "", errors.New("failed to create user")
	}

	token, err := helper.GenerateJWT(user.ID.Hex())

	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return "", errors.New("failed to generate token")
	}

	return token, nil

}

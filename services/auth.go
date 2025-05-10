package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/obanlatomiwa/go-inventory-api/database"
	"github.com/obanlatomiwa/go-inventory-api/models"
	"github.com/obanlatomiwa/go-inventory-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(userInput models.UserRequest) (string, error) {
	// create a password
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// create a new user and add the user to the database
	var user models.User = models.User{
		ID:       uuid.New().String(),
		Email:    userInput.Email,
		Password: string(password),
	}

	database.DB.Create(&user)

	// generate the jwt token
	token, err := utils.GenerateToken()
	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(userInput models.UserRequest) (string, error) {
	var user models.User

	// find the user based on email because it's the only unique entry we have stored
	result := database.DB.First(&user, "email = ?", userInput.Email)

	if result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	// compare the password input with the password from the database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		return "", errors.New("invalid password")
	}

	// generate jwt token
	token, err := utils.GenerateToken()
	
	if err != nil {
		return "", err
	}

	return token, nil
}

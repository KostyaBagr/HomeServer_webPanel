// a service layer for user
package services

import (
	"fmt"
	"os"
	"time"
	
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"

	"github.com/KostyaBagr/HomeServer_webPanel/initializers"
	"github.com/KostyaBagr/HomeServer_webPanel/models"
	
)


// func CreateUserService() (models.User, error){
// 	// Create user with password and username
// 	var authInput models.AuthInput

// 	var userFound models.User
// 	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

// 	if userFound.ID != 0 {
// 		return models.User{}, fmt.Errorf("Username already used")
// 	}

// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return models.User{}, fmt.Errorf("Error generating password hash: %v", 
// 		err)
// 	}

// 	user := models.User{
// 		Username: authInput.Username,
// 		Password: string(passwordHash),
// 	}

// 	initializers.DB.Create(&user)
// 	return user, nil
// }

func LoginService() (string, error){
	// Create and return token.

	var authInput models.AuthInput


	var userFound models.User
	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID == 0 {
		return "", fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}
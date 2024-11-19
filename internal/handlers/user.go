// Handlers for user registration
package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/KostyaBagr/HomeServer_webPanel/initializers"
	"github.com/KostyaBagr/HomeServer_webPanel/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser() gin.HandlerFunc {
	// Handler for user regisration
	return func(c *gin.Context) {
		var authInput models.AuthInput

		if err := c.ShouldBind(&authInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userFound models.User
		initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

		if userFound.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(authInput.Username)
		user := models.User{
			Username: authInput.Username,
			Password: string(passwordHash),
		}

		initializers.DB.Create(&user)

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func Login() gin.HandlerFunc{
	// Handler for user login
	return func(c *gin.Context) {
		var authInput models.AuthInput

		if err := c.ShouldBind(&authInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userFound models.User
		initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

		if userFound.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}

		generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  userFound.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
		}

		c.JSON(200, gin.H{
			"token": token,
		})
	}
	
}


func GetUserProfile() gin.HandlerFunc{
	// Get current user
	return func(c *gin.Context) {
		user, _ := c.Get("currentUser")
		c.JSON(200, gin.H{
			"user": user,
		})
	}
	
}
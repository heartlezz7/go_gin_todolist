package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/heartlezz7/go_gin_todolist/model"
	"github.com/heartlezz7/go_gin_todolist/services"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user model.CreateUser

	var err error

	err = c.Bind(&user)

	fmt.Printf("user: %v\n", user)

	if err != nil {
		c.String(http.StatusInternalServerError, "validate fail")
		return
	}
	err = services.RegisterService(user)

	if err != nil {
		c.String(http.StatusBadRequest, "create fail")
		return
	}

	c.String(http.StatusCreated, "create success")
}

func Login(c *gin.Context) {
	var user model.UserLogin

	var UserData model.UserData

	var err error

	err = c.Bind(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "validate fail")
		return
	}
	fmt.Printf("user: %v\n", user)
	err = services.LoginService(user.Email, &UserData)

	if err != nil {
		c.String(http.StatusBadRequest, "create fail")
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(UserData.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	// sign JWT
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  UserData.Id,
		"exp": time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}

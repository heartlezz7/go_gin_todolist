package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/heartlezz7/go_gin_todolist/model"
	reporepository "github.com/heartlezz7/go_gin_todolist/repositories"
)

func AuthenticateMW(c *gin.Context) {
	fmt.Println("Authenticate Middleware")

	//  c.GetHeader is  is a higher-level abstraction provided by the Context object and it might additional logic or processing before returning the header value
	token := c.GetHeader("Authorization")

	//  c.Request.Header.Get is a lower-level abstraction that directly accesses the header map of the request object
	// a := c.Request.Header.Get("Authorization")

	if token == "" || !strings.Contains(token, "Bearer ") {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
	}

	// get payload
	var tokenString = strings.Split(token, " ")[1]

	payload, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte("secretKey"), nil
	})

	// find user with id from valid token
	userId := reporepository.NewUserId(int(payload.Claims.(jwt.MapClaims)["id"].(float64)))

	var user model.UserData
	err := userId.GetUserById(&user)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
	}

	c.Set("user", user)

	c.Next()

}

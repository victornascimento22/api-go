// helloworld
package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secret-key")
var tokens []string

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {

	// r := gin.Default()
	// r.POST("/login", gin.BasicAuth(gin.Accounts{
	// 	"admin": "secret",
	// }), func(c *gin.Context) {

	// 	token, _ := generateJWT()
	// 	tokens = append(tokens, token)

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"token": token,
	// 	})

	// })

}

func createToken(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {

		return "", err

	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {

		return fmt.Errorf("invalid token")

	}
	return nil
}

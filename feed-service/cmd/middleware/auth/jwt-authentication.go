package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(scope, c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func TokenValid(api_scope string, c *gin.Context) error {
	r := c.Request
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return fmt.Errorf("Not Enough Permission")
	}

	// Setting Username to the context
	c.Set("username", claims["username"])
	
	for _, role := range scopes {
		if role == api_scope {
			return nil
		}
	}

	return fmt.Errorf("Not Enough Permission")
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	token_string := ExtractToken(r)

	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bear_token := r.Header.Get("Authorization")
	str_rrr := strings.Split(bear_token, " ")
	if len(str_rrr) == 2 {
		return str_rrr[1]
	}
	return ""
}

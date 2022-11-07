package middlewares

import (
	"crypto/rsa"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-interview/src/models"
	"net/http"
	"strings"
)

var (
	verifyKey *rsa.PublicKey
)

func AuthRequired(h gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])
		tk, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
				return nil, fmt.Errorf("invalid token %v", token.Header["alg"])

			}
			return []byte(models.TokenString), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "token error")
		}

		if tk.Valid {
			claim := tk.Claims.(jwt.MapClaims)
			userId := claim["iss"]
			c.Set("userId", userId)
			h(c)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

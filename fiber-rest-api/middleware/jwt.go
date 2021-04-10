package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	authorization = "Authorization"
	tokenExpired  = "Token expired"
	issuerInvalid = "Issuer invalid"
	audienceInvalid = "Audience invalid"
)

func ValidateJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get(authorization)
		tokenStr = tokenStr[7:len(tokenStr)]

		token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			log.Println(err)
			return c.Status(http.StatusUnauthorized).SendString(http.StatusText(http.StatusUnauthorized))
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			return c.Status(http.StatusUnauthorized).SendString(http.StatusText(http.StatusUnauthorized))
		}

		if claims.ExpiresAt < time.Now().Local().Unix() {
			return c.Status(http.StatusUnauthorized).SendString(tokenExpired)
		}

		if claims.Issuer != os.Getenv("JWT_ISSUER") {
			return c.Status(http.StatusUnauthorized).SendString(issuerInvalid)
		}

		if claims.Audience != os.Getenv("JWT_AUDIENCE") {
			return c.Status(http.StatusUnauthorized).SendString(audienceInvalid)
		}

		return c.Next()
	}
}

package middleware

import (
	"time"

	"github.com/bangadam/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() fiber.Handler {
	conf := config.NewConfig()

	if conf.Middleware.Jwt.Secret == "" {
		panic("JWT secret is not set")
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(conf.Middleware.Jwt.Secret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

type JWTClaims struct {
	Token string `json:"token"`
	Type  string `json:"type"`
	jwt.RegisteredClaims
}

func GenerateTokenAccess(userID uint64) (*JWTClaims, error) {
	conf := config.NewConfig()

	mySigningKey := []byte(conf.Middleware.Jwt.Secret)
	// Create the Claims
	exp := time.Now().Add(conf.Middleware.Jwt.Expiration)
	claims := &JWTClaims{
		Type: "Bearer",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    "go-fiber-starter",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return nil, err
	}

	claims.Token = ss
	return claims, nil
}

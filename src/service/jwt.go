package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	Name  string
	Email string
	jwt.RegisteredClaims
}

func NewJWTService(claims jwt.RegisteredClaims) *JWTService {
	return &JWTService{
		RegisteredClaims: claims,
	}
}

func (j *JWTService) CreateJWTToAccess(id string, name string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":   j.Issuer,
			"aud":   j.Audience,
			"sub":   j.Subject,
			"jti":   id,
			"Name":  name,
			"Email": email,
			"exp":   jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			"iat":   jwt.NewNumericDate(time.Now()),
			"nbf":   jwt.NewNumericDate(time.Now()),
		},
	)
	jwtString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}

func (j *JWTService) ValidateJWTToAccess(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	if auth == "" {
		code := fiber.StatusBadRequest
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg":  "you must provide the authentication code",
			"code": code,
		})
	}
	listString := strings.Split(auth, " ")
	token, err := jwt.Parse(listString[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", t.Header)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return c.Next()
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusBadRequest,
		})
	}
}

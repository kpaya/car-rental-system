package service

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWTToAccess(id string, name string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":   os.Getenv("JWT_ISSUER"),
			"aud":   os.Getenv("JWT_AUDIENCE"),
			"sub":   os.Getenv("JWT_SUBJECT"),
			"jti":   id,
			"name":  name,
			"email": email,
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

func ValidateJWTToAccess(auth_header string) error {
	listString := strings.Split(auth_header, " ")
	token, err := jwt.Parse(listString[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", t.Header)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

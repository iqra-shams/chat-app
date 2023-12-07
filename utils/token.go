package utils

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/go-chi/chi/v5"
)

func GenerateToken(username string, role string) (string, string, error) {

	access_token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	secretKey := []byte(os.Getenv("API_SECRET"))

	if err != nil {
		return "", "", err
	}

	accessExpTime := time.Now().Add(time.Minute * time.Duration(access_token_lifespan)).Unix()
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      accessExpTime,
	})
	accessTokenString, err := access_token.SignedString(secretKey)
	if err != nil {
		return "", "", err

	}

	refresh_token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", "", err
	}

	refreshExpTime := time.Now().Add(time.Hour * time.Duration(refresh_token_lifespan)).Unix()
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      refreshExpTime,
	})
	refreshTokenString, err := refresh_token.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil

}

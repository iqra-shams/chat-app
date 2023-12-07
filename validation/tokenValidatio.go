package validation

import (
	"fmt"
	"os"
	"strings"
     "github.com/dgrijalva/jwt-go"
	"github.com/iqra-shams/chat-app/utils"
)

func ValidateToken(tokenString string) (*utils.MyClaims, error) {
	secretKey := []byte(os.Getenv("API_SECRET"))
	tokenString = strings.TrimSpace(tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &utils.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Provide the key used to sign the token
		return secretKey, nil
	})
	// Check for errors
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("invalid signature")
		} else {
			fmt.Println("error parsing token:", err)
		}
		return nil, err
	}
	// Check if the token is valid
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, err
	}
	// Extract claims
	claims, ok := token.Claims.(*utils.MyClaims)
	if !ok {
		fmt.Println("error extracting claims")
		return nil, err
	}
	return claims, nil
}

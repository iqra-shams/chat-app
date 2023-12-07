package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/iqra-shams/chat-app/models"
	"github.com/iqra-shams/chat-app/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Tokens struct {
	AccessToken string `json:"accesstoken"`
	Refreshoken string `json:"refreshtoken"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var input LoginInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Unable to login", http.StatusBadRequest)
	}
	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	accessToken, refreshToken, err := LoginCheck(u.Username, u.Password)
	tokens := Tokens{
		Refreshoken: refreshToken,
		AccessToken: accessToken,
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("user or password is incorrect")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)

}
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, string, error) {

	var err error

	u := models.User{}
	db := utils.ConnectDB()
	defer db.Close()

	// Validate JWT token
	err = db.Model(models.User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", "", err

	}

	accessToken, refreshToken, err := utils.GenerateToken(u.Username, u.Role)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}

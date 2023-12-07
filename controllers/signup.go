package controllers

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/iqra-shams/chat-app/models"
	"github.com/iqra-shams/chat-app/utils"
	"golang.org/x/crypto/bcrypt"
)

type SignUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	var input SignUpInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Unable to get data", http.StatusBadRequest)
	}

	db := utils.ConnectDB()

	defer db.Close()

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	user := models.User{
		Username: input.Username,
		Password: hashedPassword,
		Role:     "user",
	}

	err = db.Create(&user).Error
	if err != nil {
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Optionally, you can return a success message or the created user data
	response := map[string]interface{}{
		"message": "User created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

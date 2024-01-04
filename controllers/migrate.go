package controllers

import (
	"net/http"

	"github.com/iqra-shams/chat-app/models"
	"github.com/iqra-shams/chat-app/utils"

	_ "github.com/lib/pq"
)

func Migrate(w http.ResponseWriter, r *http.Request) {
	db := utils.ConnectDB()
	defer db.Close()
	db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Message{})
}

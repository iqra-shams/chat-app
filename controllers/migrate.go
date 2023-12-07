package controllers

import (
	"fmt"
	"net/http"

	"github.com/iqra-shams/chat-app/models"
	"github.com/iqra-shams/chat-app/utils"

	_ "github.com/lib/pq"
)

func Migrate(w http.ResponseWriter, r *http.Request) {
	db := utils.ConnectDB()
	fmt.Println("connected2")
	defer db.Close()
	fmt.Println("migrating")
	db.AutoMigrate(&models.User{})

}

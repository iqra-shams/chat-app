package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iqra-shams/chat-app/controllers"
)

func main() {

	r := chi.NewRouter()
	// r.Use(middleware.Allow(middleware.AllowAll()))

	r.Post("/signup", controllers.SignUp)
	r.Get("/migrate", controllers.Migrate)
	r.Post("/login", controllers.Login)
	r.Get("/ws", controllers.HandleSocketConnection)
	go controllers.HandleMessages()

	fmt.Println("Starting server")

	http.ListenAndServe(":3000", r)

}

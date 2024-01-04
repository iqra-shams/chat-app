package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iqra-shams/chat-app/controllers"
	"github.com/iqra-shams/chat-app/controllers/socket"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	r := chi.NewRouter()
	
	// r.Use(middleware.Allow(middleware.AllowAll()))
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/readiness", readinessHandler)

	r.Post("/signup", controllers.SignUp)
	r.Get("/migrate", controllers.Migrate)
	r.Post("/login", controllers.Login)

	socketServer := socket.Server{}
	socketServer.InitServer()
	r.Get("/ws", socketServer.HandleSocketConnection)
	

	// go socketServer.HandleMessages()

	fmt.Println("Starting server")

	http.ListenAndServe(":3333", r)

}

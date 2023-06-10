package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luxarasis4/twittor/middleware"
	"github.com/rs/cors"
)

var (
	userController = newUserController()
)

func Managers(chanel chan error) {
	router := mux.NewRouter()

	router.HandleFunc("/auth/registry", middleware.CheckDB(userController.Registry)).Methods("POST")
	router.HandleFunc("/auth/login", middleware.CheckDB(userController.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	chanel <- http.ListenAndServe(":"+PORT, handler)
}

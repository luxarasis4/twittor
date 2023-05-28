package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luxarasis4/twittor/middleware"
	"github.com/rs/cors"
)

var (
	userController = newUserController()
)

func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/auth/registry", middleware.CheckDB(userController.Registry)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	err := http.ListenAndServe(":"+PORT, handler)
	if err != nil {
		log.Fatal(err.Error())
	}
}

package controllers

import "net/http"

type IAuthController interface {
	Registry(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/luxarasis4/twittor/dtos"
	"github.com/luxarasis4/twittor/services"
)

var (
	authControllerOnce      sync.Once
	authControllerInterface *AuthController
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	authControllerOnce.Do(func() {
		authControllerInterface = &AuthController{
			authService: authService,
		}
	})

	return authControllerInterface
}

func (controller *AuthController) Registry(w http.ResponseWriter, r *http.Request) {
	var request dtos.AuthRegistryRequestDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponseBadRequest(w, fmt.Sprintf("Bad request, error[%s]", err.Error()), err)
		return
	}

	err = dtos.ValidateStruct(request)
	if err != nil {
		writeResponseBadRequest(w, fmt.Sprintf("Bad request, error[%s]", err.Error()), err)
		return
	}

	response, status, err := controller.authService.Registry(request)
	if err != nil {
		writeResponseInternalServerError(w, fmt.Sprintf("Internal server error, error[%s]", err.Error()), err)
		return
	} else if !status {
		writeResponseBadRequest(w, "User already exist", nil)
		return
	}

	writeResponse(w, http.StatusCreated, response)
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var request dtos.AuthLoginRequestDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponseBadRequest(w, fmt.Sprintf("Bad request, error[%s]", err.Error()), err)
		return
	}

	err = dtos.ValidateStruct(request)
	if err != nil {
		writeResponseBadRequest(w, fmt.Sprintf("Bad request, error[%s]", err.Error()), err)
		return
	}

	response, status := controller.authService.Login(request)
	if !status {
		writeResponseUnauthorized(w, "Invalid password or email", nil)
		return
	}

	writeResponse(w, http.StatusOK, response)
}

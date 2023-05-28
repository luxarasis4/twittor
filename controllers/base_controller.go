package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/luxarasis4/twittor/dtos"
)

func writeResponse(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(body)
}

func writeResponseBadRequest(w http.ResponseWriter, msg string, err error) {
	statusCode := http.StatusBadRequest
	body := &dtos.GeneralResponseDTO{
		Status: "bad_request",
		Code:   statusCode,
		Data:   msg,
		Error:  err,
	}

	writeResponse(w, statusCode, body)
}

func writeResponseInternalServerError(w http.ResponseWriter, msg string, err error) {
	statusCode := http.StatusInternalServerError
	body := &dtos.GeneralResponseDTO{
		Status: "internal_server_error",
		Code:   statusCode,
		Data:   msg,
		Error:  err,
	}

	writeResponse(w, statusCode, body)
}

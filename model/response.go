package model

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, code int, status string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func Success(w http.ResponseWriter, code int, message string, data interface{}) {
	JSON(w, code, "ok", message, data)
}

func Error(w http.ResponseWriter, code int, message string) {
	JSON(w, code, "error", message, nil)
}

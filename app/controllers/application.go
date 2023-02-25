package controllers

import (
	"net/http"
)

func writeHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}

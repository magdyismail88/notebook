package main

import (
	"net/http"
)

type Tester struct {}

func (t *Tester) Ping(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Pong!\n"))
}


func (t *Tester) GetImage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
  	w.Header().Set("Content-Type", "application/json")
  	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Pong!\n"))
}


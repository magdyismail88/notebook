package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/magdyismail88/notebook/bootstrap"
	routes "github.com/magdyismail88/notebook/config"
	"github.com/rs/cors"
)

func main() {

	fmt.Println(`
		Server listening on 8888
		http://localhost:8888
	`)

	app := bootstrap.App()
	env := app.Env

	handler := cors.Default().Handler(routes.Setup(env))

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

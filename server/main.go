package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/magdyismail88/notebook/bootstrap"
	"github.com/magdyismail88/notebook/config/routes"
	"github.com/rs/cors"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	msg := fmt.Sprintf(`
		Server listening on %s
		http://localhost:%s
	`, env.Port, env.Port)

	fmt.Println(msg)

	handler := cors.Default().Handler(routes.Setup(env))

	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%s", env.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

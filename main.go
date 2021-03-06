package main

import (
    // "fmt"
    // "database/sql"
    // "encoding/json"
    "net/http"
    "time"
    "log"
    // "io"
    "github.com/rs/cors"
    // "strconv"
    // "github.com/gorilla/mux"
    // "github.com/gorilla/handlers"
    // _ "github.com/mattn/go-sqlite3"
)

// omitempty

// type Message struct {
//     Success bool `json:"success"`
// }

var router Router

func main() {


    // r.HandleFunc("/api/notes/{tab_id:[0-9]+}", GetNotesForTab).Methods("GET")
    // handlerMethods := handlers.CORS(handlers.AllowedOrigins([]string{"*"})
    handler := cors.Default().Handler(router.All())

    srv := &http.Server {
        Handler: handler,
        Addr: "localhost:8888",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
//
    // log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))

    // http.ListenAndServe(":8080", r)

    // log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))

    // cors := handlers.CORS( handlers.AllowedOrigins([]string{"*"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}) )(r)

    log.Fatal(srv.ListenAndServe())

    // go func() {
    //     if err := srv.ListenAndServe(); err != nil {
    //         log.Println(err)
    //     }
    // }()

    // log.Fatal(srv.ListenAndServe(),  cors)
}

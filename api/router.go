package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct{}

var containerCtrl ContainerController

func (router *Router) All() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	fs := http.FileServer(http.Dir("storage"))

	r.Handle("/storage", http.StripPrefix("/storage", fs))

	r.HandleFunc("/storage/{file_name}", func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r);
		http.ServeFile(w, r, r.URL.Path[1:])

	})

	// r.Headers("X-Requested-With", "XMLHttpRequest")
	// Container routes
	r.HandleFunc("/api/containers", containerCtrl.FindAll).Methods("GET")
	r.HandleFunc("/api/containers/{container_id}", containerCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/containers/create", containerCtrl.Create).Methods("POST")
	r.HandleFunc("/api/containers/update", containerCtrl.Update).Methods("POST")
	r.HandleFunc("/api/containers/delete", containerCtrl.Destroy).Methods("POST")

	r.HandleFunc("/api/tabs/{container_id}", GetTabs).Methods("GET") // Get all tabs
	r.HandleFunc("/api/tabs", CreateTab).Methods("POST")             // Create new tab
	r.HandleFunc("/api/tabs/{tab_id}/delete", DestroyTab).Methods("GET")
	r.HandleFunc("/api/tabs/{tab_id}/update", UpdateTab).Methods("POST")
	r.HandleFunc("/api/tabs/{tab_id}/tab", GetTab).Methods("GET")

	r.HandleFunc("/api/notes/{tab_id}", GetNotesForTab).Methods("GET") // Get notes for each tab
	r.HandleFunc("/api/notes", CreateNote).Methods("POST")             // Create note
	r.HandleFunc("/api/notes/{note_id}/delete", DestroyNote).Methods("GET")
	r.HandleFunc("/api/notes/{note_id}/update", UpdateNote).Methods("POST")

	r.HandleFunc("/api/notes/{note_id}/show", GetNote).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

	uploader := Uploader{}

	r.HandleFunc("/api/upload", uploader.Upload)

	// storage := Storage{}

	// r.HandleFunc("/api/storage/{file_name}", storage.Get).Methods("GET")

	return r

}

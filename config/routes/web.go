package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/bootstrap"
)

func (k *kernel) SetupWebRoutes(r *mux.Router, env *bootstrap.Env) {
	fs := http.FileServer(http.Dir("public/storage"))
	r.Handle("/public/storage", http.StripPrefix("/public/storage", fs))

	r.HandleFunc("/tab/{id}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	r.HandleFunc("/note/new", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	r.HandleFunc("/note/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	r.HandleFunc("/public/storage/{file_name}", func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r);
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	// filePath := fmt.Sprintf("/api/%s", env.StoragePath)
}

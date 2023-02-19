package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/controllers"
	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/bootstrap"
)

func Setup(env *bootstrap.Env) *mux.Router {
	containerCtrl := &controllers.ContainerController{
		Container: models.NewContainer(env),
		Env:       env,
	}
	tabCtrl := &controllers.TabController{
		Tab: models.NewTab(env),
		Env: env,
	}
	noteCtrl := &controllers.NoteController{
		Note: models.NewNote(env),
		Env:  env,
	}

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
	// r.HandleFunc("/api/containers/update", containerCtrl.Update).Methods("POST")
	r.HandleFunc("/api/containers/delete", containerCtrl.Destroy).Methods("POST")

	r.HandleFunc("/api/tabs/{container_id}", tabCtrl.FindAll).Methods("GET") // Get all tabs
	r.HandleFunc("/api/tabs/{tab_id}/tab", tabCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/tabs", tabCtrl.Create).Methods("POST") // Create new tab
	r.HandleFunc("/api/tabs/{tab_id}/delete", tabCtrl.Destroy).Methods("GET")
	r.HandleFunc("/api/tabs/{tab_id}/update", tabCtrl.Update).Methods("POST")

	r.HandleFunc("/api/notes/{tab_id}", noteCtrl.FindAll).Methods("GET") // Get notes for each tab
	r.HandleFunc("/api/notes/{note_id}/show", noteCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/notes", noteCtrl.Create).Methods("POST") // Create note
	r.HandleFunc("/api/notes/{note_id}/delete", noteCtrl.Destroy).Methods("GET")
	r.HandleFunc("/api/notes/{note_id}/update", noteCtrl.Update).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))
	// r.PathPrefix("/tab/[0-9]+").Handler(http.FileServer(http.Dir("./frontend/dist")))

	uploader := controllers.Uploader{}

	r.HandleFunc("/api/upload", uploader.Upload)

	// storage := Storage{}

	// r.HandleFunc("/api/storage/{file_name}", storage.Get).Methods("GET")

	return r

}

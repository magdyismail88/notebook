package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/controllers"
	"github.com/magdyismail88/notebook/app/services"
	"github.com/magdyismail88/notebook/bootstrap"
)

func Setup(env *bootstrap.Env) *mux.Router {
	containerCtrl := &controllers.ContainerController{
		ContainerService: services.NewContainerService(env),
		Env:              env,
	}

	tabCtrl := &controllers.TabController{
		TabService: services.NewTabService(env),
		Env:        env,
	}

	noteCtrl := &controllers.NoteController{
		NoteService: services.NewNoteService(env),
		Env:         env,
	}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	fs := http.FileServer(http.Dir("storage"))
	r.Handle("/storage", http.StripPrefix("/storage", fs))

	r.HandleFunc("/tab/{id}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	r.HandleFunc("/note/new", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	r.HandleFunc("/note/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	r.HandleFunc("/storage/{file_name}", func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r);
		http.ServeFile(w, r, r.URL.Path[1:])

	})

	// r.Headers("X-Requested-With", "XMLHttpRequest")
	// Container routes
	r.HandleFunc("/api/v1/containers", containerCtrl.FindAll).Methods("GET")
	r.HandleFunc("/api/v1/containers/{container_id}", containerCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/v1/containers/create", containerCtrl.Create).Methods("POST")
	// r.HandleFunc("/api/v1/containers/update", containerCtrl.Update).Methods("POST")
	r.HandleFunc("/api/v1/containers/delete", containerCtrl.Destroy).Methods("POST")

	r.HandleFunc("/api/v1/tabs/{container_id}", tabCtrl.FindAll).Methods("GET") // Get all tabs
	r.HandleFunc("/api/v1/tabs/{tab_id}/tab", tabCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/v1/tabs", tabCtrl.Create).Methods("POST") // Create new tab
	r.HandleFunc("/api/v1/tabs/{tab_id}/delete", tabCtrl.Destroy).Methods("GET")
	r.HandleFunc("/api/v1/tabs/{tab_id}/update", tabCtrl.Update).Methods("POST")

	r.HandleFunc("/api/v1/notes/{tab_id}", noteCtrl.FindAll).Methods("GET") // Get notes for each tab
	r.HandleFunc("/api/v1/notes/{note_id}/show", noteCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/v1/notes", noteCtrl.Create).Methods("POST") // Create note
	r.HandleFunc("/api/v1/notes/{note_id}/delete", noteCtrl.Destroy).Methods("GET")
	r.HandleFunc("/api/v1/notes/{note_id}/update", noteCtrl.Update).Methods("POST")

	// r.PathPrefix("/tab/[0-9]+").Handler(http.FileServer(http.Dir("./frontend/dist")))

	uploader := &controllers.Uploader{Env: env}

	r.HandleFunc("/api/v1/upload", uploader.Upload)
	// filePath := fmt.Sprintf("/api/%s", env.StoragePath)
	// r.HandleFunc("/api/{filename}", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Content-Type", "application/image")
	// 	w.WriteHeader(http.StatusOK)

	// 	vars := mux.Vars(req)

	// 	fileName := vars["filename"]
	// 	fileURL := fmt.Sprintf("http://localhost:%s/api/%s/%s", env.Port, env.StoragePath, fileName)

	// 	io.WriteString(w, `{"success": true, "image_url": `+fileURL+`}`)

	// }).Methods("GET")

	// storage := Storage{}
	// r.HandleFunc("/tab/15", nil).Handler(http.FileServer(http.Dir("./frontend/dist")))

	// r.HandleFunc("/api/storage/{file_name}", storage.Get).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

	return r

}

package routes

import (
	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/bootstrap"
)

func (k *kernel) SetupApiRoutes(r *mux.Router, env *bootstrap.Env) {
	// r.Headers("X-Requested-With", "XMLHttpRequest")
	// Container routes
	r.HandleFunc("/api/containers", k.ContainerCtrl.FindAll).Methods("GET")
	r.HandleFunc("/api/containers/{container_id}", k.ContainerCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/containers/create", k.ContainerCtrl.Create).Methods("POST")
	// r.HandleFunc("/api/containers/update", containerCtrl.Update).Methods("POST")
	r.HandleFunc("/api/containers/delete", k.ContainerCtrl.Destroy).Methods("POST")

	r.HandleFunc("/api/tabs/{container_id}", k.TabCtrl.FindAll).Methods("GET") // Get all tabs
	r.HandleFunc("/api/tabs/{tab_id}/tab", k.TabCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/tabs", k.TabCtrl.Create).Methods("POST") // Create new tab
	r.HandleFunc("/api/tabs/{tab_id}/delete", k.TabCtrl.Destroy).Methods("GET")
	r.HandleFunc("/api/tabs/{tab_id}/update", k.TabCtrl.Update).Methods("POST")

	r.HandleFunc("/api/notes/{tab_id}", k.NoteCtrl.FindAll).Methods("GET") // Get notes for each tab
	r.HandleFunc("/api/notes/{note_id}/show", k.NoteCtrl.FindOne).Methods("GET")
	r.HandleFunc("/api/notes", k.NoteCtrl.Create).Methods("POST") // Create note
	r.HandleFunc("/api/notes/{note_id}/delete", k.NoteCtrl.Destroy).Methods("GET")
	r.HandleFunc("/api/notes/{note_id}/update", k.NoteCtrl.Update).Methods("POST")

	r.HandleFunc("/api/upload", k.UploaderCtrl.Upload)
	// r.PathPrefix("/tab/[0-9]+").Handler(http.FileServer(http.Dir("./frontend/dist")))

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
}

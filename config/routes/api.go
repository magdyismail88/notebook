package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/bootstrap"
)

func (k *kernel) SetupApiRoutes(r *mux.Router, env *bootstrap.Env) {
	// r.Headers("X-Requested-With", "XMLHttpRequest")
	// Container routes
	r.HandleFunc("/api/containers", k.ContainerCtrl.FindAll).Methods("GET")
	r.HandleFunc("/api/containers/{id}", k.ContainerCtrl.FindOne).Methods(http.MethodGet)
	r.HandleFunc("/api/containers", k.ContainerCtrl.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/containers/{id}", k.ContainerCtrl.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/containers/{id}", k.ContainerCtrl.Destroy).Methods(http.MethodDelete)

	r.HandleFunc("/api/tabs/{containerId}/tabs", k.TabCtrl.FindAll).Methods(http.MethodGet) // Get all tabs
	r.HandleFunc("/api/tabs/{id}", k.TabCtrl.FindOne).Methods(http.MethodGet)
	r.HandleFunc("/api/tabs", k.TabCtrl.Create).Methods(http.MethodPost) // Create new tab
	r.HandleFunc("/api/tabs/{id}", k.TabCtrl.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/tabs/{id}", k.TabCtrl.Destroy).Methods(http.MethodDelete)

	r.HandleFunc("/api/notes/{tabId}/notes", k.NoteCtrl.FindAll).Methods(http.MethodGet) // Get notes for each tab
	r.HandleFunc("/api/notes/{id}", k.NoteCtrl.FindOne).Methods(http.MethodGet)
	r.HandleFunc("/api/notes", k.NoteCtrl.Create).Methods(http.MethodPost) // Create note
	r.HandleFunc("/api/notes/{id}", k.NoteCtrl.Update).Methods(http.MethodPut)
	r.HandleFunc("/api/notes/{id}", k.NoteCtrl.Destroy).Methods(http.MethodDelete)
	r.HandleFunc("/api/actions/move-note-to", k.NoteCtrl.ChangeTab).Methods(http.MethodPut)

	r.HandleFunc("/api/containers-tabs", k.NoteCtrl.FetchAllContainersWithTabs).Methods(http.MethodGet)

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

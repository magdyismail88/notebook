package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/controllers"
	"github.com/magdyismail88/notebook/app/services"
	"github.com/magdyismail88/notebook/bootstrap"
)

type kernel struct {
	ContainerCtrl *controllers.ContainerController
	TabCtrl       *controllers.TabController
	NoteCtrl      *controllers.NoteController
	UploaderCtrl  *controllers.Uploader
}

func handleCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(w, r)
	})
}

func Setup(db *sql.DB, env *bootstrap.Env) *mux.Router {
	containerCtrl := &controllers.ContainerController{
		ContainerService: services.NewContainerService(db),
	}

	tabCtrl := &controllers.TabController{
		TabService: services.NewTabService(db),
	}

	noteCtrl := &controllers.NoteController{
		NoteService: services.NewNoteService(db),
	}

	k := &kernel{
		ContainerCtrl: containerCtrl,
		TabCtrl:       tabCtrl,
		NoteCtrl:      noteCtrl,
		UploaderCtrl:  controllers.NewUploadController(env),
	}

	r := mux.NewRouter()
	r.Use(handleCors)
	// r.Use(mux.CORSMethodMiddleware(r))

	k.SetupApiRoutes(r, env)
	k.SetupWebRoutes(r, env)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))
	return r
}

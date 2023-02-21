package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/controllers/v1"
	"github.com/magdyismail88/notebook/app/services"
	"github.com/magdyismail88/notebook/bootstrap"
)

type kernel struct {
	ContainerCtrl *controllers.ContainerController
	TabCtrl       *controllers.TabController
	NoteCtrl      *controllers.NoteController
	UploaderCtrl  *controllers.Uploader
}

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

	k := &kernel{
		ContainerCtrl: containerCtrl,
		TabCtrl:       tabCtrl,
		NoteCtrl:      noteCtrl,
		UploaderCtrl:  controllers.NewUploadController(env),
	}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	k.SetupApiRoutes(r, env)
	k.SetupWebRoutes(r, env)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))
	return r

}

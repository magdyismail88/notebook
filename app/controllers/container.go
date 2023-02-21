package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/app/services"
	"github.com/magdyismail88/notebook/bootstrap"
)

type ContainerController struct {
	ContainerService services.ContainerService
	Env              *bootstrap.Env
}

type containerForm struct {
	models.Container
}

func (cc *ContainerController) FindAll(w http.ResponseWriter, _ *http.Request) {
	writeHeader(w, http.StatusOK)

	containers, err := cc.ContainerService.FindAll()

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&containers)
	return
}

func (cc *ContainerController) FindOne(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusOK)

	vars := mux.Vars(r)
	containerID, _ := strconv.ParseInt(vars["container_id"], 10, 64)
	container, err := cc.ContainerService.FindOne(int(containerID))

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&container)
	return
}

func (cc *ContainerController) Create(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusCreated)

	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	err := cc.ContainerService.Create(form.Title)

	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

func (cc *ContainerController) Destroy(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusNoContent)

	var err error
	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	_, err = cc.ContainerService.FindOne(form.ID)
	if err != nil {
		panic(err)
	}

	err = cc.ContainerService.Delete(form.ID)
	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/app/services"
)

type ContainerController struct {
	ContainerService services.ContainerService
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
	container, err := cc.ContainerService.FindOne(vars["id"])

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

func (cc *ContainerController) Update(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusNoContent)
	// tab_id
	vars := mux.Vars(r)
	container, err := cc.ContainerService.FindOne(vars["id"])

	if err != nil {
		panic(err)
	}

	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)

	container.Title = form.Title
	err = cc.ContainerService.Update(container)

	if err != nil {
		panic(err)
	}

}

func (cc *ContainerController) Destroy(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusNoContent)

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := cc.ContainerService.FindOne(id)

	if err != nil {
		panic(err)
	}

	err = cc.ContainerService.Delete(id)

	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

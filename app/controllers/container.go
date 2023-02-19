package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/bootstrap"
)

type ContainerController struct {
	Container models.Container
	Env       *bootstrap.Env
}

type containerForm struct {
	models.Container
}

func (cc *ContainerController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	containers, err := cc.Container.FindAll()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(&containers)
	return
}

func (cc *ContainerController) FindOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	containerID, _ := strconv.ParseInt(vars["container_id"], 10, 64)
	container, err := cc.Container.FindOne(int(containerID))
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(&container)
	return
}

func (cc *ContainerController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	err := cc.Container.Create(form.Title)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
	return
}

func (cc *ContainerController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	var err error
	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	container, err := cc.Container.FindOne(form.ID)
	if err != nil {
		panic(err)
	}
	container.Title = form.Title
	err = container.Save()
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
	return
}

func (cc *ContainerController) Destroy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	var err error
	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	container, err := cc.Container.FindOne(form.ID)
	if err != nil {
		panic(err)
	}
	err = container.Delete()
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
	return
}

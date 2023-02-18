package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/models"
)

type ContainerController struct {
	models.Container
}

type containerForm struct {
	models.Container
}

func (cc *ContainerController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	containers, err := models.FindAllContainers()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(&containers)
}

func (cc *ContainerController) FindOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	productId, _ := strconv.ParseInt(vars["productId"], 10, 64)
	container, err := models.FindContainer(int(productId))
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(&container)
}

func (cc *ContainerController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	err := models.NewContainer(form.Title)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
}

func (cc *ContainerController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	var err error
	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	container, err := models.FindContainer(form.ID)
	if err != nil {
		panic(err)
	}
	err = cc.Container.Update(container)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
}

func (cc *ContainerController) Destroy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	var err error
	var form containerForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	container, err := models.FindContainer(form.ID)
	if err != nil {
		panic(err)
	}
	err = container.Delete()
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
}

func DestroyContainer() {

	return
}

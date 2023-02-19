package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/bootstrap"
	_ "github.com/mattn/go-sqlite3"
)

type TabController struct {
	Tab models.Tab
	Env *bootstrap.Env
}

type tabForm struct {
	models.Tab
}

func (tc *TabController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	containerID, _ := strconv.ParseInt(vars["container_id"], 10, 64)
	tabs, err := tc.Tab.FindAll(int(containerID))

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&tabs)
	return
}

func (tc *TabController) FindOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Get tab id
	vars := mux.Vars(r)
	tabID, _ := strconv.ParseInt(vars["tab_id"], 10, 64)
	tab, err := tc.Tab.FindOne(int(tabID))
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(&tab)
	return
}

func (tc *TabController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var form tabForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	err := tc.Tab.Create(form.Title, form.ContainerID)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
	return
}

func (tc *TabController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	// tab_id
	vars := mux.Vars(r)
	tabID, _ := strconv.ParseInt(vars["tab_id"], 10, 64)
	tab, err := tc.Tab.FindOne(int(tabID))
	if err != nil {
		panic(err)
	}
	var form tabForm
	tab.Title = form.Title
	_ = json.NewDecoder(r.Body).Decode(&form)
	err = tc.Tab.Update(tab)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
	return
}

func (tc *TabController) Destroy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	vars := mux.Vars(r)
	tabID, _ := strconv.ParseInt(vars["tab_id"], 10, 64)
	tab, err := tc.Tab.FindOne(int(tabID))
	if err != nil {
		panic(err)
	}
	err = tab.Delete(int(tabID))
	if err != nil {
		panic(err)
	}
	io.WriteString(w, `{"success": true}`)
	return
}

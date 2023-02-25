package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/app/services"
	_ "github.com/mattn/go-sqlite3"
)

type TabController struct {
	TabService services.TabService
}

type tabForm struct {
	models.Tab
}

func (tc *TabController) FindAll(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusOK)

	vars := mux.Vars(r)

	tabs, err := tc.TabService.FindAll(vars["containerId"])

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&tabs)
	return
}

// Find tab and fetch all notes into this tab
func (tc *TabController) FindOne(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusOK)
	// Get tab id
	vars := mux.Vars(r)

	tab, err := tc.TabService.FindOne(vars["id"])

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&tab)
	return
}

func (tc *TabController) Create(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusCreated)

	var form tabForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	err := tc.TabService.Create(form.Title, form.ContainerID)

	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

func (tc *TabController) Update(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusNoContent)
	// tab_id
	vars := mux.Vars(r)

	tab, err := tc.TabService.FindOne(vars["id"])

	if err != nil {
		panic(err)
	}

	var form tabForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	tab.Title = form.Title
	err = tc.TabService.Update(tab)

	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

func (tc *TabController) Destroy(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusNoContent)

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := tc.TabService.FindOne(id)

	if err != nil {
		panic(err)
	}

	err = tc.TabService.Delete(id)

	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

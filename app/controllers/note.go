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

type NoteController struct {
	NoteService services.NoteService
	Env         *bootstrap.Env
}
type noteForm struct {
	models.Note
}

func (nc *NoteController) FindAll(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusOK)

	vars := mux.Vars(r)
	tabID, _ := strconv.ParseInt(vars["tab_id"], 10, 64)
	notes, err := nc.NoteService.FindAll(int(tabID))

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&notes)
	return
}

func (nc *NoteController) FindOne(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusOK)

	vars := mux.Vars(r)
	noteID, _ := strconv.ParseInt(vars["note_id"], 10, 64)
	note, err := nc.NoteService.FindOne(int(noteID))

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&note)
	return
}

func (nc *NoteController) Create(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusCreated)

	var form noteForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	err := nc.NoteService.Create(form.Title, form.Content, form.TabID)

	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

func (nc *NoteController) Update(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusNoContent)

	vars := mux.Vars(r)
	noteID, _ := strconv.ParseInt(vars["note_id"], 10, 64)
	note, err := nc.NoteService.FindOne(int(noteID))

	if err != nil {
		panic(err)
	}

	var form noteForm
	_ = json.NewDecoder(r.Body).Decode(&form)
	note.Title = form.Title
	note.Content = form.Content
	err = nc.NoteService.Update(note)

	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

func (nc *NoteController) Destroy(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusNoContent)

	vars := mux.Vars(r)
	noteID, _ := strconv.ParseInt(vars["note_id"], 10, 64)
	_, err := nc.NoteService.FindOne(int(noteID))

	if err != nil {
		panic(err)
	}

	err = nc.NoteService.Delete(int(noteID))

	if err != nil {
		panic(err)
	}

	io.WriteString(w, `{"success": true}`)
	return
}

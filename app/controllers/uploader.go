package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/magdyismail88/notebook/bootstrap"
)

type Uploader struct {
	Env *bootstrap.Env
}

func (u *Uploader) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	defer w.WriteHeader(http.StatusOK)

	file, handler, err := r.FormFile("file")

	// fileName := r.FormValue("file_name")

	fileURL := fmt.Sprint("http://localhost:8000/storage/", handler.Filename)

	// json.NewEncoder(w).Encode(map[string]string{"link": fileURL})

	// return

	if err != nil {
		panic(err)
	}

	defer file.Close()

	f, err := os.OpenFile("storage/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// _, _ = io.WriteString(w, "File " + fileName + " uploaded successfully")
	_, _ = io.Copy(f, file)

	json.NewEncoder(w).Encode(map[string]string{"link": fileURL})
}

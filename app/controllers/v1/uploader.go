package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/magdyismail88/notebook/bootstrap"
)

type Uploader struct {
	Env *bootstrap.Env
}

func NewUploadController(env *bootstrap.Env) *Uploader {
	return &Uploader{Env: env}
}

func (u *Uploader) Upload(w http.ResponseWriter, r *http.Request) {
	writeHeader(w, http.StatusOK)

	file, handler, err := r.FormFile("file")

	// fileName := r.FormValue("file_name")

	fileURL := fmt.Sprintf(
		"http://localhost:%s/%s/%s", u.Env.Port, u.Env.StoragePath, handler.Filename,
	)

	// json.NewEncoder(w).Encode(map[string]string{"link": fileURL})

	// return

	if err != nil {
		panic(err)
	}

	defer file.Close()

	f, err := os.OpenFile(path.Join(u.Env.StoragePath, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// _, _ = io.WriteString(w, "File " + fileName + " uploaded successfully")
	_, _ = io.Copy(f, file)

	json.NewEncoder(w).Encode(map[string]string{"link": fileURL})
}

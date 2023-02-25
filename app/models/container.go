package models

import (
	_ "github.com/mattn/go-sqlite3"
)

const ContainerTable = "containers"

type Container struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

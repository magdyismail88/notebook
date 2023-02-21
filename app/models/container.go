package models

import (
	_ "github.com/mattn/go-sqlite3"
)

const ContainerTable = "containers"

type Container struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

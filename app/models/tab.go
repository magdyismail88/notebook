package models

import (
	_ "github.com/mattn/go-sqlite3"
)

const TabTable = "tabs"

type Tab struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ContainerID string `json:"containerId"`
}

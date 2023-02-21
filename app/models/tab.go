package models

import (
	_ "github.com/mattn/go-sqlite3"
)

const TabTable = "tabs"

type Tab struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	ContainerID int    `json:"container_id"`
}

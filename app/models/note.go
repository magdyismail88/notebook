package models

const NoteTable = "notes"

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	TabID   int    `json:"tab_id"`
}

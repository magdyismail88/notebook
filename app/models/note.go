package models

const NoteTable = "notes"

type Note struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentText string `json:"contentText"`
	TabID       string `json:"tabId"`
}

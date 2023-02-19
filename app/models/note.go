package models

import (
	"database/sql"
	"fmt"

	"github.com/magdyismail88/notebook/bootstrap"
)

const noteTable = "notes"

type NoteEntity struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	TabID   int    `json:"tab_id"`
}

type Note struct {
	NoteEntity
	Env *bootstrap.Env
}

func NewNote(env *bootstrap.Env) Note {
	return Note{Env: env}
}

func (n *Note) FindAll(tabID int) ([]*Note, error) {
	var notes []*Note
	database, err := sql.Open(n.Env.DatabaseAdapter, n.Env.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `tab_id` FROM %s WHERE `tab_id` = ?", noteTable)
	rows, err := database.Query(stmt, tabID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var note Note
		rows.Scan(&note.ID, &note.Title, &note.Content, &note.TabID)
		notes = append(notes, &note)
	}
	return notes, nil
}

func (n *Note) FindOne(noteID int) (*Note, error) {
	var note Note
	database, err := sql.Open(n.Env.DatabaseAdapter, n.Env.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `tab_id` FROM `%s` WHERE `id` = ?", noteTable)
	res := database.QueryRow(stmt, noteID)
	res.Scan(&note.ID, &note.Title, &note.Content, &note.TabID)
	return &note, nil
}

func (n *Note) Create(title, content string, tabID int) error {
	database, err := sql.Open(n.Env.DatabaseAdapter, n.Env.DatabaseName)
	if err != nil {
		return err
	}
	defer database.Close()
	stmt := fmt.Sprintf("INSERT INTO %s (id, title, content, tab_id) VALUES (NULL, ?, ?, ?)", noteTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		panic(err)
	}
	_, err = res.Exec(title, content, tabID)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Update(note *Note) error {
	database, err := sql.Open(n.Env.DatabaseAdapter, n.Env.DatabaseName)
	if err != nil {
		return err
	}
	defer database.Close()
	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ?, `content` = ? WHERE id = ?", noteTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(note.Title, note.Content, note.ID)
	if err != nil {
		return err
	}
	return nil
}

func (n *Note) Delete(id int) error {
	database, err := sql.Open(n.Env.DatabaseAdapter, n.Env.DatabaseName)
	if err != nil {
		return err
	}
	defer database.Close()
	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", noteTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

package services

import (
	"database/sql"
	"fmt"

	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/bootstrap"
)

type NoteService struct {
	Env *bootstrap.Env
}

func NewNoteService(env *bootstrap.Env) NoteService {
	return NoteService{Env: env}
}

func (ns *NoteService) FindAll(tabID int) ([]*models.Note, error) {
	var notes []*models.Note

	database, err := sql.Open(ns.Env.DatabaseAdapter, ns.Env.DatabaseName)
	if err != nil {
		return nil, err
	}

	defer database.Close()

	stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `tab_id` FROM %s WHERE `tab_id` = ?", models.NoteTable)
	rows, err := database.Query(stmt, tabID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note models.Note
		rows.Scan(&note.ID, &note.Title, &note.Content, &note.TabID)
		notes = append(notes, &note)
	}

	return notes, nil
}

func (ns *NoteService) FindOne(noteID int) (*models.Note, error) {
	var note models.Note

	database, err := sql.Open(ns.Env.DatabaseAdapter, ns.Env.DatabaseName)
	if err != nil {
		return nil, err
	}

	defer database.Close()

	stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `tab_id` FROM `%s` WHERE `id` = ?", models.NoteTable)
	res := database.QueryRow(stmt, noteID)

	res.Scan(&note.ID, &note.Title, &note.Content, &note.TabID)

	return &note, nil
}

func (ns *NoteService) Create(title, content string, tabID int) error {
	database, err := sql.Open(ns.Env.DatabaseAdapter, ns.Env.DatabaseName)

	if err != nil {
		return err
	}

	defer database.Close()

	stmt := fmt.Sprintf("INSERT INTO %s (id, title, content, tab_id) VALUES (NULL, ?, ?, ?)", models.NoteTable)
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

func (ns *NoteService) Update(note *models.Note) error {
	database, err := sql.Open(ns.Env.DatabaseAdapter, ns.Env.DatabaseName)

	if err != nil {
		return err
	}

	defer database.Close()
	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ?, `content` = ? WHERE id = ?", models.NoteTable)

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

func (ns *NoteService) Delete(id int) error {
	database, err := sql.Open(ns.Env.DatabaseAdapter, ns.Env.DatabaseName)
	if err != nil {
		return err
	}

	defer database.Close()
	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", models.NoteTable)
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

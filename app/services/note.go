package services

import (
	"database/sql"
	"fmt"

	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/bootstrap"
)

type NoteService struct {
	db  *sql.DB
	Env *bootstrap.Env
}

func NewNoteService(db *sql.DB, env *bootstrap.Env) NoteService {
	return NoteService{db: db, Env: env}
}

func (ns *NoteService) FindAll(tabID int) ([]*models.Note, error) {
	var notes []*models.Note

	stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `tab_id` FROM %s WHERE `tab_id` = ?", models.NoteTable)
	rows, err := ns.db.Query(stmt, tabID)

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

	stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `tab_id` FROM `%s` WHERE `id` = ?", models.NoteTable)
	res := ns.db.QueryRow(stmt, noteID)

	res.Scan(&note.ID, &note.Title, &note.Content, &note.TabID)

	return &note, nil
}

func (ns *NoteService) Create(title, content string, tabID int) error {
	stmt := fmt.Sprintf("INSERT INTO %s (id, title, content, tab_id) VALUES (NULL, ?, ?, ?)", models.NoteTable)
	res, err := ns.db.Prepare(stmt)
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
	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ?, `content` = ? WHERE id = ?", models.NoteTable)

	res, err := ns.db.Prepare(stmt)
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
	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", models.NoteTable)
	res, err := ns.db.Prepare(stmt)

	if err != nil {
		return err
	}
	_, err = res.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

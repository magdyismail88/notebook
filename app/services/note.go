package services

import (
	"database/sql"
	"fmt"

	"github.com/kennygrant/sanitize"
	"github.com/magdyismail88/notebook/app/models"
)

type NoteService struct {
	db *sql.DB
}

func NewNoteService(db *sql.DB) NoteService {
	return NoteService{db: db}
}

func (ns *NoteService) FindAll(tabID string) ([]*models.Note, error) {
	var notes []*models.Note

	stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `content_text`, `tab_id` FROM %s WHERE `tab_id` = ?", models.NoteTable)
	rows, err := ns.db.Query(stmt, tabID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note models.Note
		rows.Scan(&note.ID, &note.Title, &note.Content, &note.ContentText, &note.TabID)
		notes = append(notes, &note)
	}

	return notes, nil
}

func (ns *NoteService) FindOne(noteID string) (*models.Note, error) {
	var note models.Note

	stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `content_text`, `tab_id` FROM `%s` WHERE `id` = ?", models.NoteTable)
	res := ns.db.QueryRow(stmt, noteID)

	res.Scan(&note.ID, &note.Title, &note.Content, &note.ContentText, &note.TabID)

	return &note, nil
}

func (ns *NoteService) Create(title, content, tabID string) error {
	stmt := fmt.Sprintf("INSERT INTO %s (`id`, `title`, `content`, `content_text`, `tab_id`) VALUES (?, ?, ?, ?, ?)", models.NoteTable)

	res, err := ns.db.Prepare(stmt)

	if err != nil {
		panic(err)
	}

	id := generateUUIDV4()

	var contentText string

	sanitizedText := sanitizeHTML(content)
	if len(sanitizedText) > 800 {
		contentText = sanitizedText[:800]
	} else {
		contentText = sanitizedText
	}

	_, err = res.Exec(id, title, content, contentText, tabID)
	if err != nil {
		return err
	}

	return nil
}

func (ns *NoteService) Update(note *models.Note) error {
	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ?, `content` = ?, `content_text` = ? WHERE id = ?", models.NoteTable)

	res, err := ns.db.Prepare(stmt)
	if err != nil {
		return err
	}

	contentText := sanitizeHTML(note.Content)

	_, err = res.Exec(note.Title, note.Content, contentText, note.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ns *NoteService) GetContainersAndTabs() string {
	query := `
		SELECT %s.title, %s.title
		FROM %
		INNER JOIN %
		ON %.id = %.container_id
	`
	stmt := fmt.Sprintf(
		query,
		models.ContainerTable,
		models.TabTable,
		models.ContainerTable,
		models.TabTable,
		models.ContainerTable,
		models.TabTable,
	)

	return stmt
}

func (ns *NoteService) Delete(id string) error {
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

// func (ns *NoteService) MoveTo(id, query string) error {
// 	// delete current note by id
// 	// split query by > symbol
// 	// first element is container, second element is tab
// 	// if split has two element so move to container > tab
// 	// if split has one element so move to another tab
// 	elements := strings.Split(query, ">")

// 	return nil
// }

func sanitizeHTML(content string) (contentText string) {
	// allowingChars := []string{"<p>", "</p>"}
	// contentText, _ = sanitize.HTMLAllowing(content, allowingChars)
	contentText = sanitize.HTML(content)
	return
}

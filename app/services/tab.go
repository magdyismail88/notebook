package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/bootstrap"
	_ "github.com/mattn/go-sqlite3"
)

type TabService struct {
	db  *sql.DB
	Env *bootstrap.Env
}

func NewTabService(db *sql.DB, env *bootstrap.Env) TabService {
	return TabService{db: db, Env: env}
}

func (ts *TabService) FindAll(containerID int) ([]*models.Tab, error) {
	var tabs []*models.Tab

	stmt := fmt.Sprintf("SELECT `id`, `title`, `slug`, `container_id` FROM `%s` WHERE `container_id` = ?", models.TabTable)

	rows, err := ts.db.Query(stmt, containerID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tab models.Tab
		rows.Scan(&tab.ID, &tab.Title, &tab.Slug, &tab.ContainerID)
		tabs = append(tabs, &tab)
	}

	return tabs, nil
}

func (ts *TabService) FindOne(tabID int) (*models.Tab, error) {
	var tab models.Tab

	stmt := fmt.Sprintf("SELECT `id`, `title`, `container_id` FROM `%s` WHERE `id` = ?", models.TabTable)

	res := ts.db.QueryRow(stmt, tabID)

	res.Scan(&tab.ID, &tab.Title, &tab.ContainerID)

	return &tab, nil
}

func (ts *TabService) Create(title string, containerID int) error {

	stmt := fmt.Sprintf("INSERT INTO `%s` (id, title, slug, container_id) VALUES (NULL, ?, ?, ?)", models.TabTable)

	res, err := ts.db.Prepare(stmt)

	if err != nil {
		return err
	}

	slug := generateSlug(title)

	_, err = res.Exec(title, slug, containerID)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TabService) Update(tab *models.Tab) error {

	slug := generateSlug(tab.Title)

	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ?, `slug` = ? WHERE `id` = ?", models.TabTable)

	res, err := ts.db.Prepare(stmt)

	if err != nil {
		return err
	}

	_, err = res.Exec(tab.Title, slug, tab.ID)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TabService) Delete(id int) error {
	// Delete all notes related this tab
	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE `tab_id` = ?", models.NoteTable)

	res, err := ts.db.Prepare(stmt)
	if err != nil {
		return err
	}

	_, err = res.Exec(id)

	if err != nil {
		panic(err)
	}
	// Delete tab
	stmt = fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", models.TabTable)
	res, err = ts.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func generateSlug(title string) string {
	return strings.ReplaceAll(strings.ToLower(title), " ", "-")
}

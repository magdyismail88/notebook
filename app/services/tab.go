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
	Env *bootstrap.Env
}

func NewTabService(env *bootstrap.Env) TabService {
	return TabService{Env: env}
}

func (ts *TabService) FindAll(containerID int) ([]*models.Tab, error) {
	var tabs []*models.Tab

	database, err := sql.Open(ts.Env.DatabaseAdapter, ts.Env.DatabaseName)

	if err != nil {
		return nil, err
	}

	defer database.Close()

	stmt := fmt.Sprintf("SELECT `id`, `title`, `slug`, `container_id` FROM `%s` WHERE `container_id` = ?", models.TabTable)

	rows, err := database.Query(stmt, containerID)

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

func (t *TabService) FindOne(tabID int) (*models.Tab, error) {
	var tab models.Tab

	database, err := sql.Open(t.Env.DatabaseAdapter, t.Env.DatabaseName)

	if err != nil {
		return nil, err
	}

	defer database.Close()

	stmt := fmt.Sprintf("SELECT `id`, `title`, `container_id` FROM `%s` WHERE `id` = ?", models.TabTable)

	res := database.QueryRow(stmt, tabID)

	if err != nil {
		return nil, err
	}

	res.Scan(&tab.ID, &tab.Title, &tab.ContainerID)

	return &tab, nil
}

func (t *TabService) Create(title string, containerID int) error {
	database, err := sql.Open(t.Env.DatabaseAdapter, t.Env.DatabaseName)

	if err != nil {
		return err
	}

	defer database.Close()

	stmt := fmt.Sprintf("INSERT INTO `%s` (id, title, slug, container_id) VALUES (NULL, ?, ?, ?)", models.TabTable)

	res, err := database.Prepare(stmt)
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

func (t *TabService) Update(tab *models.Tab) error {
	database, err := sql.Open(t.Env.DatabaseAdapter, t.Env.DatabaseName)

	if err != nil {
		return err
	}

	defer database.Close()

	slug := generateSlug(tab.Title)

	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ?, `slug` = ? WHERE `id` = ?", models.TabTable)

	res, err := database.Prepare(stmt)

	if err != nil {
		return err
	}

	_, err = res.Exec(tab.Title, slug, tab.ID)

	if err != nil {
		return err
	}

	return nil
}

func (t *TabService) Delete(id int) error {
	database, err := sql.Open(t.Env.DatabaseAdapter, t.Env.DatabaseName)

	if err != nil {
		return err
	}
	defer database.Close()
	// Delete all notes related this tab
	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE `tab_id` = ?", models.NoteTable)

	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}

	_, err = res.Exec(id)

	if err != nil {
		panic(err)
	}
	// Delete tab
	stmt = fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", models.TabTable)
	res, err = database.Prepare(stmt)
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

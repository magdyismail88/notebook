package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/magdyismail88/notebook/bootstrap"
	_ "github.com/mattn/go-sqlite3"
)

const tabTable = "tabs"

type Tab struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	ContainerID int    `json:"container_id"`
	Env         *bootstrap.Env
}

func NewTab(env *bootstrap.Env) *Tab {
	return &Tab{Env: env}
}

func (t *Tab) FindAll(containerID int) ([]*Tab, error) {
	var tabs []*Tab
	database, err := sql.Open(t.Env.DatabaseAdapter, t.Env.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	stmt := fmt.Sprintf("SELECT `id`, `title`, `slug`, `container_id` FROM `%s` WHERE `container_id` = ?", tabTable)
	rows, err := database.Query(stmt, containerID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tab Tab
		rows.Scan(&tab.ID, &tab.Title, &tab.Slug, &tab.ContainerID)
		tabs = append(tabs, &tab)
	}
	return tabs, nil
}

func (t *Tab) FindOne(tabID int) (*Tab, error) {
	var tab Tab
	database, err := sql.Open(DB_ADAPTER, DB_PATH)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	stmt := fmt.Sprintf("SELECT `id`, `title`, `container_id` FROM `%s` WHERE `id` = ?", tabTable)
	res := database.QueryRow(stmt, tabID)
	if err != nil {
		return nil, err
	}
	res.Scan(&tab.ID, &tab.Title, &tab.ContainerID)
	return &tab, nil
}

func (t *Tab) Create(title string, containerID int) error {
	database, err := sql.Open(t.Env.DatabaseAdapter, t.Env.DatabaseName)
	if err != nil {
		return err
	}
	defer database.Close()
	stmt := fmt.Sprintf("INSERT INTO `%s` (id, title, slug, container_id) VALUES (NULL, ?, ?, ?)", tabTable)
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

func (t *Tab) Update(tab *Tab) error {
	database, err := sql.Open(t.Env.DatabaseAdapter, t.Env.DatabaseName)
	if err != nil {
		return err
	}
	defer database.Close()
	slug := generateSlug(tab.Title)
	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ?, `slug` = ? WHERE `id` = ?", tabTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec("HERE", slug, tab.ContainerID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tab) Delete(tabID int) error {
	database, err := sql.Open(t.Env.DatabaseAdapter, t.Env.DatabaseName)
	if err != nil {
		return err
	}
	defer database.Close()
	// Delete all notes related this tab
	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE `tab_id` = ?", noteTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(t.ID)
	if err != nil {
		panic(err)
	}
	// Delete tab
	stmt = fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", tabTable)
	res, err = database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(t.ID)
	if err != nil {
		return err
	}
	return nil
}

func generateSlug(title string) string {
	return strings.ReplaceAll(strings.ToLower(title), " ", "-")
}

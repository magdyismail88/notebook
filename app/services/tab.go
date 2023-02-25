package services

import (
	"database/sql"
	"fmt"

	"github.com/magdyismail88/notebook/app/models"
	_ "github.com/mattn/go-sqlite3"
)

type TabService struct {
	db *sql.DB
}

func NewTabService(db *sql.DB) TabService {
	return TabService{db: db}
}

func (ts *TabService) FindAll(containerID string) ([]*models.Tab, error) {
	var tabs []*models.Tab

	stmt := fmt.Sprintf("SELECT `id`, `title`, `container_id` FROM `%s` WHERE `container_id` = ?", models.TabTable)

	rows, err := ts.db.Query(stmt, containerID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tab models.Tab
		rows.Scan(&tab.ID, &tab.Title, &tab.ContainerID)
		tabs = append(tabs, &tab)
	}

	return tabs, nil
}

func (ts *TabService) FindOne(tabID string) (*models.Tab, error) {
	var tab models.Tab

	stmt := fmt.Sprintf("SELECT `id`, `title`, `container_id` FROM `%s` WHERE `id` = ?", models.TabTable)

	res := ts.db.QueryRow(stmt, tabID)

	res.Scan(&tab.ID, &tab.Title, &tab.ContainerID)

	return &tab, nil
}

func (ts *TabService) Create(title, containerID string) error {

	stmt := fmt.Sprintf("INSERT INTO `%s` (id, title, container_id) VALUES (?, ?, ?)", models.TabTable)

	res, err := ts.db.Prepare(stmt)

	if err != nil {
		return err
	}

	id := generateUUIDV4()
	_, err = res.Exec(id, title, containerID)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TabService) Update(tab *models.Tab) error {

	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ? WHERE `id` = ?", models.TabTable)

	res, err := ts.db.Prepare(stmt)

	if err != nil {
		return err
	}

	_, err = res.Exec(tab.Title, tab.ID)

	if err != nil {
		return err
	}

	return nil
}

func (ts *TabService) Delete(id string) error {
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

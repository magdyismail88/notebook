package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_ADAPTER = "sqlite3"
	DB_PATH    = "./data/notebook.db"
)

const containerTable = "containers"

type Container struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func FindAllContainers() ([]*Container, error) {
	var containers []*Container
	database, err := sql.Open(DB_ADAPTER, DB_PATH)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s`", containerTable)
	res, err := database.Query(stmt)
	var container Container
	for res.Next() {
		res.Scan(&container.ID, &container.Title)
		containers = append(containers, &container)
	}
	return containers, nil
}

func FindContainer(id int) (*Container, error) {
	var container Container
	database, err := sql.Open(DB_ADAPTER, DB_PATH)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s` WHERE id = ?", containerTable)
	res := database.QueryRow(stmt, id)
	if err != nil {
		return nil, err
	}
	res.Scan(&container.ID, &container.Title)
	return &container, nil
}

func NewContainer(title string) error {
	database, err := sql.Open(DB_ADAPTER, DB_PATH)
	if err != nil {
		return err
	}
	defer database.Close()
	stmt := fmt.Sprintf("INSERT INTO `%s` (id, title) VALUES (NULL, ?)", containerTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(title)
	if err != nil {
		return err
	}
	return nil
}

func (c *Container) Update(container *Container) error {
	database, err := sql.Open(DB_ADAPTER, DB_PATH)
	if err != nil {
		return err
	}
	defer database.Close()
	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ? WHERE `id` = ?", containerTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(&container.Title, &container.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Container) Delete() error {
	database, err := sql.Open(DB_ADAPTER, DB_PATH)
	if err != nil {
		return err
	}
	defer database.Close()
	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", containerTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(c.ID)
	if err != nil {
		return err
	}
	return nil
}

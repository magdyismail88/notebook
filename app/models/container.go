package models

import (
	"database/sql"
	"fmt"

	"github.com/magdyismail88/notebook/bootstrap"
	_ "github.com/mattn/go-sqlite3"
)

const containerTable = "containers"

type Container struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Env   *bootstrap.Env
}

func NewContainer(env *bootstrap.Env) Container {
	return Container{Env: env}
}

func (c *Container) FindAll() ([]*Container, error) {
	var containers []*Container
	database, err := sql.Open(c.Env.DatabaseAdapter, c.Env.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer database.Close()
	stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s`", containerTable)
	res, err := database.Query(stmt)
	for res.Next() {
		var container Container
		res.Scan(&container.ID, &container.Title)
		containers = append(containers, &container)
	}
	return containers, nil
}

func (c *Container) FindOne(id int) (*Container, error) {
	var container Container
	database, err := sql.Open(c.Env.DatabaseAdapter, c.Env.DatabaseName)
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

func (c *Container) Create(title string) error {
	database, err := sql.Open(c.Env.DatabaseAdapter, c.Env.DatabaseName)
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

func (c *Container) Save() error {
	database, err := sql.Open(c.Env.DatabaseAdapter, c.Env.DatabaseName)
	if err != nil {
		return err
	}
	defer database.Close()
	stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ? WHERE `id` = ?", containerTable)
	res, err := database.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = res.Exec(c.Title, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Container) Delete() error {
	database, err := sql.Open(c.Env.DatabaseAdapter, c.Env.DatabaseName)
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

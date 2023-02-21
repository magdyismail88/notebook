package services

import (
	"database/sql"
	"fmt"

	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/bootstrap"
	_ "github.com/mattn/go-sqlite3"
)

type ContainerService struct {
	Env *bootstrap.Env
}

func NewContainerService(env *bootstrap.Env) ContainerService {
	return ContainerService{Env: env}
}

func (cs *ContainerService) FindAll() ([]*models.Container, error) {
	var containers []*models.Container

	database, err := sql.Open(cs.Env.DatabaseAdapter, cs.Env.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s`", models.ContainerTable)
	res, err := database.Query(stmt)

	for res.Next() {
		var container models.Container
		res.Scan(&container.ID, &container.Title)
		containers = append(containers, &container)
	}

	return containers, nil
}

func (cs *ContainerService) FindOne(id int) (*models.Container, error) {
	var container models.Container

	database, err := sql.Open(cs.Env.DatabaseAdapter, cs.Env.DatabaseName)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s` WHERE id = ?", models.ContainerTable)
	res := database.QueryRow(stmt, id)

	if err != nil {
		return nil, err
	}

	res.Scan(&container.ID, &container.Title)

	return &container, nil
}

func (cs *ContainerService) Create(title string) error {
	database, err := sql.Open(cs.Env.DatabaseAdapter, cs.Env.DatabaseName)
	if err != nil {
		return err
	}
	defer database.Close()

	stmt := fmt.Sprintf("INSERT INTO `%s` (id, title) VALUES (NULL, ?)", models.ContainerTable)
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

func (cs *ContainerService) Delete(id int) error {
	database, err := sql.Open(cs.Env.DatabaseAdapter, cs.Env.DatabaseName)

	if err != nil {
		return err
	}

	defer database.Close()

	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", models.ContainerTable)
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

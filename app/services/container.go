package services

import (
	"database/sql"
	"fmt"

	"github.com/magdyismail88/notebook/app/models"
	"github.com/magdyismail88/notebook/bootstrap"
	_ "github.com/mattn/go-sqlite3"
)

type ContainerService struct {
	db  *sql.DB
	Env *bootstrap.Env
}

func NewContainerService(db *sql.DB, env *bootstrap.Env) ContainerService {
	return ContainerService{db: db, Env: env}
}

func (cs *ContainerService) FindAll() ([]*models.Container, error) {
	var containers []*models.Container

	stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s`", models.ContainerTable)
	res, err := cs.db.Query(stmt)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		var container models.Container
		res.Scan(&container.ID, &container.Title)
		containers = append(containers, &container)
	}

	return containers, nil
}

func (cs *ContainerService) FindOne(id int) (*models.Container, error) {
	var container models.Container

	stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s` WHERE id = ?", models.ContainerTable)
	res := cs.db.QueryRow(stmt, id)

	res.Scan(&container.ID, &container.Title)

	return &container, nil
}

func (cs *ContainerService) Create(title string) error {
	stmt := fmt.Sprintf("INSERT INTO `%s` (id, title) VALUES (NULL, ?)", models.ContainerTable)

	res, err := cs.db.Prepare(stmt)

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
	stmt := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", models.ContainerTable)
	res, err := cs.db.Prepare(stmt)

	if err != nil {
		return err
	}
	_, err = res.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

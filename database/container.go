package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Container struct{}

func (c *Container) Up() error {
	database, err := sql.Open(DatabaseAdapter, DatabasePath)
	defer database.Close()

	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `containers` (`id` CHAR NOT NULL PRIMARY KEY UNIQUE, `title` CHAR)")

	_, err = database.Exec(stmt)

	if err != nil {
		return err
	}

	return nil
}

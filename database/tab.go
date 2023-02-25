package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Tab struct{}

func (t *Tab) Up() error {
	database, err := sql.Open(DatabaseAdapter, DatabasePath)
	defer database.Close()

	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `tabs` (`id` CHAR NOT NULL UNIQUE PRIMARY KEY, `title` TEXT NOT NULL, `container_id` CHAR)")

	_, err = database.Exec(stmt)

	if err != nil {
		return err
	}

	return nil
}

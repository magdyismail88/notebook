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

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `containers` (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, title CHAR)")

	_, err = database.Exec(stmt)

	if err != nil {
		return err
	}

	return nil
}

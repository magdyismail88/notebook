package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct{}

func (n *Note) Up() error {
	database, err := sql.Open(DatabaseAdapter, DatabasePath)
	defer database.Close()

	if err != nil {
		return err
	}

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `notes` (`id` CHAR NOT NULL PRIMARY KEY UNIQUE, `title` TEXT NOT NULL, `tab_id` CHAR NOT NULL,`content` TEXT, `content_text` TEXT)")

	_, err = database.Exec(stmt)

	if err != nil {
		return err
	}

	return nil
}

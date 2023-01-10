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

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `notes` (`id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, `title` TEXT NOT NULL, `tab_id`  INTEGER NOT NULL,`content` TEXT, `is_textual` NUMERIC DEFAULT 0)")

	_, err = database.Exec(stmt)

	if err != nil {
		return err
	}

	return nil
}

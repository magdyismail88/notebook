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

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `tabs` (`id` INTEGER NOT NULL UNIQUE, `title` TEXT NOT NULL, `parent`  INTEGER, `slug`  TEXT NOT NULL, `container_id` INTEGER, PRIMARY KEY(`id` AUTOINCREMENT), FOREIGN KEY(`parent`) REFERENCES `tabs`(`id`))")

	_, err = database.Exec(stmt)

	if err != nil {
		return err
	}

	return nil
}

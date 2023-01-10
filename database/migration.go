package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var container Container
var tab Tab
var note Note

func checkConnection() error {
	database, err := sql.Open(DatabaseAdapter, DatabasePath)
	defer database.Close()

	if err != nil {
		return fmt.Errorf("Database not connected %s", err)
	}

	return nil
}

func Run() error {

	if err := checkConnection(); err != nil {
		return err
	}

	if err := container.Up(); err != nil {
		return err
	}

	if err := tab.Up(); err != nil {
		return err
	}

	if err := note.Up(); err != nil {
		return err
	}

	fmt.Println("Database migrated.")

	return nil
}

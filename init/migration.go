package main

import (
  "fmt"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
)

const (
  DatabaseAdapter = "sqlite3"
  DatabasePath = "./data/notebook.db"
)

func createContainersTable() {
  database, err := sql.Open(DatabaseAdapter, DatabasePath)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `containers` (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, title CHAR)")

  _, err = database.Exec(stmt)

  if err != nil {
    panic(err)
  }

  return
}

func createTabsTable() {
  database, err := sql.Open(DatabaseAdapter, DatabasePath)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `tabs` (`id` INTEGER NOT NULL UNIQUE, `title` TEXT NOT NULL, `parent`  INTEGER, `slug`  TEXT NOT NULL, `container_id` INTEGER, PRIMARY KEY(`id` AUTOINCREMENT), FOREIGN KEY(`parent`) REFERENCES `tabs`(`id`))")

  _, err = database.Exec(stmt)

  if err != nil {
    panic(err)
  }

  return
}

func createNotesTable() {
  database, err := sql.Open(DatabaseAdapter, DatabasePath)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `notes` (`id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, `title` TEXT NOT NULL, `tab_id`  INTEGER NOT NULL,`content` TEXT, `is_textual` NUMERIC DEFAULT 0)")

  _, err = database.Exec(stmt)

  if err != nil {
    panic(err)
  }

  return
}

func main() {
  createContainersTable()
  createTabsTable()
  createNotesTable()
  log.Println("Migration done")
}
package main

import (
  "fmt"
  "database/sql"
  "encoding/json"
  "net/http"
  "io"
  "github.com/gorilla/mux"
  _ "github.com/mattn/go-sqlite3"
)

type Note struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Content string `json:"content"`
    TabID int `json:"tab_id"`
    IsTextual bool `json:"is_textual"`
}

func GetNotesForTab(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  vars := mux.Vars(r)

  var notes []Note

  database, err := sql.Open(DB_ADAPTER, DB_PATH)
  defer database.Close()

  if err != nil {
      panic(err)
  }

  stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `tab_id`, `is_textual` FROM `%s` WHERE `tab_id` = %s", DB_NOTES, vars["tab_id"])

  rows, err := database.Query(stmt)

  if err != nil {
      panic(err)
  }

  var note Note

  for rows.Next() {
      rows.Scan(&note.ID, &note.Title, &note.Content, &note.TabID ,&note.IsTextual)
      notes = append(notes, note)
  }

  json.NewEncoder(w).Encode(&notes)
  return
}

func GetNote(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  vars := mux.Vars(r)

  var note Note

  database, err := sql.Open(DB_ADAPTER, DB_PATH)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("SELECT `id`, `title`, `content`, `tab_id`, `is_textual` FROM `%s` WHERE `id` = ?", DB_NOTES)

  res := database.QueryRow(stmt, vars["note_id"])

  res.Scan(&note.ID, &note.Title, &note.Content, &note.TabID, &note.IsTextual)

  json.NewEncoder(w).Encode(&note)
  return
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)

  var note Note

  _ = json.NewDecoder(r.Body).Decode(&note)

  database, err := sql.Open(DB_ADAPTER, DB_PATH)
  defer database.Close()

  if err != nil {
      panic(err)
  }

  // stmt := sql.Sprintf("INSERT INTO notes (id, title, content) VALUES (NULL, %s, %s)", vars["title"], vars["content"])

  stmt := fmt.Sprintf("INSERT INTO %s (id, title, content, tab_id, is_textual) VALUES (NULL, ?, ?, ?, ?)", DB_NOTES)

  res, err := database.Prepare(stmt)

  if err != nil {
      panic(err)
  }

  _, err = res.Exec(&note.Title, &note.Content, &note.TabID, &note.IsTextual)


  if err != nil {
      panic(err)
  }

  // id, err := res.LastInsertId()

  // message := &Message{Success: true}

  // json.NewEncoder(w).Encode(&message)

  io.WriteString(w, `{"success": true}`)
  // map[string]bool{"ok": true}
  return
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusNoContent)

  vars := mux.Vars(r)

  var note Note

  _ = json.NewDecoder(r.Body).Decode(&note)

  database, err := sql.Open(DB_ADAPTER, DB_PATH)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ?, `content` = ?, `is_textual` = ? WHERE id = ?", DB_NOTES)

  res, err := database.Prepare(stmt)

  if err != nil {
    panic(err)
  }

  _, err = res.Exec(&note.Title, &note.Content, &note.IsTextual, vars["note_id"])

  if err != nil {
    panic(err)
  }

  io.WriteString(w, `{"success": true}`)
  return
}

func DestroyNote(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusNoContent)

  vars := mux.Vars(r)

  database, err := sql.Open(DB_ADAPTER, DB_PATH)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", DB_NOTES)

  res, err := database.Prepare(stmt)

  if err != nil {
    panic(err)
  }

  _, err = res.Exec(vars["note_id"])

  if err != nil {
    panic(err)
  }

  io.WriteString(w, `{"success": true}`)
  return
}
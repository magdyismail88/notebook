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

type Tab struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Slug string `json:"slug"`
    ContainerID int `json:"container_id"`
}

func GetTabs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    vars := mux.Vars(r)

    var tabs []Tab

    database, err := sql.Open(DB_ADAPTER, DB_PATH)
    defer database.Close()

    if err != nil {
        panic(err)
    }

    stmt := fmt.Sprintf("SELECT `id`, `title`, `slug`, `container_id` FROM `%s` WHERE `container_id` = ?", DB_TABS)

    rows, err := database.Query(stmt, vars["container_id"])

    if err != nil {
        panic(err)
    }

    var tab Tab;

    for rows.Next() {
        rows.Scan(&tab.ID, &tab.Title, &tab.Slug, &tab.ContainerID)
        tabs = append(tabs, tab)
        // fmt.Println(strconv.Itoa(id) + ": " + title)
    }

    json.NewEncoder(w).Encode(&tabs)
    return
}

func GetTab(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  vars := mux.Vars(r)

  var tab Tab

  database, err := sql.Open(DB_ADAPTER, DB_PATH)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("SELECT `id`, `title`, `container_id` FROM `%s` WHERE `id` = ?", DB_TABS)

  res := database.QueryRow(stmt, vars["tab_id"])

  if err != nil {
    panic(err)
  }

  res.Scan(&tab.ID, &tab.Title, &tab.ContainerID)

  json.NewEncoder(w).Encode(&tab)
  return
}

func CreateTab(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    var tab Tab

    _ = json.NewDecoder(r.Body).Decode(&tab)


    database, err := sql.Open(DB_ADAPTER, DB_PATH)
    defer database.Close()

    if err != nil {
        panic(err)
    }

    stmt := fmt.Sprintf("INSERT INTO `%s` (id, title, slug, container_id) VALUES (NULL, ?, ?, ?)", DB_TABS)

    res, err := database.Prepare(stmt)

    if err != nil {
        panic(err)
    }

    _, err = res.Exec(&tab.Title, &tab.Title, &tab.ContainerID)


    if err != nil {
        panic(err)
    }

    database.Close()

    // message := &Message{ Success: true }

    // json.NewEncoder(w).Encode(&message)

    io.WriteString(w, `{"success": true}`)
    return
}

func UpdateTab(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNoContent)

    vars := mux.Vars(r)

    var tab Tab

    _ = json.NewDecoder(r.Body).Decode(&tab)

    database, err := sql.Open("sqlite3", DB_PATH)
    defer database.Close()

    if err != nil {
      panic(err)
    }

    stmt := fmt.Sprintf("UPDATE `%s` SET title = ? WHERE id = ?", DB_TABS)

    res, err := database.Prepare(stmt)

    if err != nil {
      panic(err)
    }

    _, err = res.Exec(&tab.Title, vars["tab_id"])

    if err != nil {
      panic(err)
    }

    io.WriteString(w, `{"success": true}`)
    return
}

func DestroyTab(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusNoContent)

  vars := mux.Vars(r)

  database, err := sql.Open("sqlite3", DB_PATH)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("DELETE FROM `%s` WHERE `tab_id` = ?", DB_NOTES)

  res, err := database.Prepare(stmt)

  if err != nil {
    panic(err)
  }

  _, err = res.Exec(vars["tab_id"])

  if err != nil {
    panic(err)
  }

  stmt = fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", DB_TABS)

  res, err = database.Prepare(stmt)

  if err != nil {
    panic(err)
  }

  _, err = res.Exec(vars["tab_id"])

  if err != nil {
    panic(err)
  }

  io.WriteString(w, `{"success": true}`)
  return
}
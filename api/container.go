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

type Container struct {
    ID int `json:"id"`
    Title string `json:"title"`
}

func GetContainers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    var containers []Container

    database, err := sql.Open(DB_ADAPTER, DB_PATH)
    defer database.Close()

    if err != nil {
        panic(err)
    }

    stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s`", DB_CONTAINERS)

    res, err := database.Query(stmt)

    var container Container

    for res.Next() {
        res.Scan(&container.ID, &container.Title)
        containers = append(containers, container)
    }

    json.NewEncoder(w).Encode(&containers)
    return
}

func GetContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    vars := mux.Vars(r)

    var container Container

    database, err := sql.Open(DB_ADAPTER, DB_PATH)
    defer database.Close()

    if err != nil {
        panic(err)
    }

    stmt := fmt.Sprintf("SELECT `id`, `title` FROM `%s` WHERE id = ?", DB_CONTAINERS)

    res := database.QueryRow(stmt, vars["container_id"])

    if err != nil {
        panic(err)
    }

    res.Scan(&container.ID, &container.Title)

    json.NewEncoder(w).Encode(&container)
    return
}

func CreateContainer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    var container Container

    _ = json.NewDecoder(r.Body).Decode(&container)

    database, err := sql.Open(DB_ADAPTER, DB_PATH)
    defer database.Close()

    if err != nil {
        panic(err)
    }

    stmt := fmt.Sprintf("INSERT INTO `%s` (id, title) VALUES (NULL, ?)", DB_CONTAINERS)

    res, err := database.Prepare(stmt)

    if err != nil {
        panic(err)
    }

    _, err = res.Exec(&container.Title)

    if err != nil {
        panic(err)
    }

    io.WriteString(w, `{"success": true}`)
    return
}

func UpdateContainer(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusNoContent)

  var container Container

  _ = json.NewDecoder(r.Body).Decode(&container)

  database, err := sql.Open(DB_ADAPTER, DB_PATH)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("UPDATE `%s` SET `title` = ? WHERE `id` = ?", DB_CONTAINERS)

  res, err := database.Prepare(stmt)

  if err != nil {
    panic(err)
  }

  _, err = res.Exec(&container.Title, &container.ID)

  if err != nil {
    panic(err)
  }

  io.WriteString(w, `{"success": true}`)
  return
}

func DestroyContainer(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusNoContent)

  var container Container

  _ = json.NewDecoder(r.Body).Decode(&container)
  
  database, err := sql.Open(DB_ADAPTER, DB_PATH)
  defer database.Close()

  if err != nil {
    panic(err)
  }

  stmt := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", DB_CONTAINERS)

  res, err := database.Prepare(stmt)

  if err != nil {
    panic(err)
  }

  _, err = res.Exec(&container.ID)

  if err != nil {
    panic(err)
  }

  io.WriteString(w, `{"success": true}`)
  return
}